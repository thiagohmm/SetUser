package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/godror/godror"
	"github.com/joho/godotenv"
)

// UsuarioService encapsula as operações relacionadas a usuários
type UsuarioService struct {
	db *sql.DB
}

// NovoUsuarioService cria uma nova instância de UsuarioService
func NovoUsuarioService(db *sql.DB) *UsuarioService {
	return &UsuarioService{db: db}
}

// ObterIDRevendedor obtém o ID do revendedor a partir do código IBM
func (u *UsuarioService) ObterIDRevendedor(codigoIBM string) (int, error) {
	var idRevendedor int
	query := "SELECT IDREVENDEDOR FROM revendedor WHERE codigoibm = :1"
	err := u.db.QueryRow(query, codigoIBM).Scan(&idRevendedor)
	return idRevendedor, err
}

// InserirUsuario insere um novo usuário na tabela usuario
func (u *UsuarioService) InserirUsuario(email string, idRevendedor int) (int, error) {
	var maxID int
	insertSQL := `
		INSERT INTO usr_hermes.usuario (
			ID, EMAIL, IDREVENDEDOR, ATIVO, DATACRIACAO, DATAALTERACAO
		) VALUES (
			(SELECT NVL(MAX(ID), 0) + 1 FROM usr_hermes.usuario),
			:1,
			:2,
			'1',
			SYSDATE,
			SYSDATE
		)
	`
	_, err := u.db.Exec(insertSQL, email, idRevendedor)
	if err != nil {
		return 0, err
	}

	// Obtém o ID do usuário recém inserido
	err = u.db.QueryRow("SELECT NVL(MAX(ID), 0) FROM usr_hermes.usuario").Scan(&maxID)
	return maxID, err
}

// InserirUsuarioAcesso insere um novo acesso para o usuário na tabela usuario_acessos
func (u *UsuarioService) InserirUsuarioAcesso(idUsuario int) error {
	insertSQL := `
		INSERT INTO usr_hermes.usuario_acessos (
			ID, IDTIPOUSUARIO, IDUSUARIO, ATIVO, DATACRIACAO, DATAALTERACAO
		) VALUES (
			:1, 2, :2, '1', SYSDATE, SYSDATE
		)
	`
	_, err := u.db.Exec(insertSQL, idUsuario, idUsuario)
	return err
}

func main() {
	// Verifica se os parâmetros foram passados
	if len(os.Args) != 3 {
		log.Fatalf("Uso: %s <codigo_ibm> <email>", os.Args[0])
	}

	codigoIBM := os.Args[1]
	email := os.Args[2]

	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	// Obtém as variáveis de ambiente
	user := os.Getenv("DB_Hermes_USER")
	password := os.Getenv("DB_Hermes_PASSWD")
	connectString := os.Getenv("DB_Hermes_CONNECTSTRING")

	// Monta a string de conexão
	dsn := fmt.Sprintf("%s/%s@%s", user, password, connectString)

	// Abre a conexão com o banco de dados Oracle
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	// Inicia uma transação
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Erro ao iniciar transação: %v", err)
	}

	// Cria o serviço de usuário
	usuarioService := NovoUsuarioService(tx)

	// Obtém o ID do revendedor
	idRevendedor, err := usuarioService.ObterIDRevendedor(codigoIBM)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Erro ao obter ID do revendedor: %v", err)
	}

	// Insere o usuário
	idUsuario, err := usuarioService.InserirUsuario(email, idRevendedor)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Erro ao inserir usuário: %v", err)
	}

	// Insere o acesso do usuário
	err = usuarioService.InserirUsuarioAcesso(idUsuario)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Erro ao inserir acesso do usuário: %v", err)
	}

	// Confirma a transação
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Erro ao confirmar a transação: %v", err)
	}

	fmt.Println("Transação concluída com sucesso!")
}

