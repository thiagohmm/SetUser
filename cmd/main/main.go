// cmd/main/main.go
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/setUserDb/config"
	"github.com/setUserDb/infraestructure/database"
	"github.com/setUserDb/infraestructure/repository"
	"github.com/setUserDb/interface/cli"
	"github.com/setUserDb/usecase"
)

func main() {

	if len(os.Args) != 4 {
		log.Fatalf("Uso: %s <codigo_ibm> %s <email> %s <permissao> \n permissao 0 para revendedor e 1 para Admin", os.Args[1], os.Args[2], os.Args[3])
	}

	codigoIBM := os.Args[1]
	email := os.Args[2]
	permissaoStr := os.Args[3]
	permissao, err := strconv.Atoi(permissaoStr)
	if err != nil {
		log.Fatalf("Erro ao converter permissão: %v", err)
	}

	// Carrega a configuração
	cfg, err := config.LoadConfig("../../.env")
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Conecta ao banco de dados

	db, err := database.ConectarBanco(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Cria o Unit of Work
	uow, err := repository.NovoUnitOfWork(db)
	if err != nil {
		log.Fatalf("Erro ao criar Unit of Work: %v", err)
	}

	// Cria o caso de uso com o UoW
	usuarioUseCase := &usecase.UsuarioUseCase{UoW: uow}

	// Cria o handler e executa a ação
	handler := cli.NovoUsuarioHandler(usuarioUseCase)
	handler.CadastrarUsuarioCLI(codigoIBM, email, permissao)
}
