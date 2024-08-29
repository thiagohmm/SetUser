// cmd/main/main.go
package main

import (
	"log"
	"os"

	"github.com/setUserDb/config"
	"github.com/setUserDb/infraestructure/database"
	"github.com/setUserDb/infraestructure/repository"
	"github.com/setUserDb/interface/cli"
	"github.com/setUserDb/usecase"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Uso: %s <codigo_ibm> <email>", os.Args[0])
	}

	codigoIBM := os.Args[1]
	email := os.Args[2]

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
	handler.CadastrarUsuarioCLI(codigoIBM, email)
}
