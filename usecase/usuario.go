// usecase/usuario.go
package usecase

import "github.com/setUserDb/domain"

type UsuarioUseCase struct {
	UoW domain.UnitOfWork
}

func (u *UsuarioUseCase) CadastrarUsuario(codigoIBM, email string) error {
	err := u.UoW.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			u.UoW.Rollback()
		}
	}()

	idRevendedor, err := u.UoW.UsuarioRepository().ObterIDRevendedor(codigoIBM)
	if err != nil {
		u.UoW.Rollback()
		return err
	}

	usuario := domain.Usuario{
		Email:        email,
		IDRevendedor: idRevendedor,
	}

	idUsuario, err := u.UoW.UsuarioRepository().InserirUsuario(usuario)
	if err != nil {
		u.UoW.Rollback()
		return err
	}

	err = u.UoW.UsuarioRepository().InserirUsuarioAcesso(idUsuario)
	if err != nil {
		u.UoW.Rollback()
		return err
	}

	err = u.UoW.Commit()
	if err != nil {
		return err
	}

	return nil
}
