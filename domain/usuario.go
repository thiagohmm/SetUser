package domain

type Usuario struct {
	ID           int
	Email        string
	IDRevendedor int
}

type UsuarioRepository interface {
	ObterIDRevendedor(codigoIBM string) (int, error)
	InserirUsuario(usuario Usuario) (int, error)
	InserirUsuarioAcesso(idUsuario int) error
	ChecarUsuarioExistente(email string) (bool, error)
}
