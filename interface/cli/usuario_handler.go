// interface/cli/usuario_handler.go
package cli

import (
	"fmt"
	"log"

	"github.com/setUserDb/usecase"
)

type UsuarioHandler struct {
	UsuarioUseCase *usecase.UsuarioUseCase
}

func NovoUsuarioHandler(u *usecase.UsuarioUseCase) *UsuarioHandler {
	return &UsuarioHandler{UsuarioUseCase: u}
}

func (h *UsuarioHandler) CadastrarUsuarioCLI(codigoIBM, email string) {

	err := h.UsuarioUseCase.CadastrarUsuario(codigoIBM, email)
	if err != nil {
		log.Fatalf("Erro ao cadastrar usuário: %v", err)
	}

	fmt.Println("Usuário cadastrado com sucesso!")
}
