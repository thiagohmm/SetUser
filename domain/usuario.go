package domain

type Usuario struct {
	Email        string
	IDRevendedor int
	Permissao    int
}

type UsuarioRepository interface {
	ObterIDRevendedor(codigoIBM string) (int, error)
	InserirUsuario(usuario Usuario) (int, error)
	InserirUsuarioAcesso(idUsuario int, permissao int) error
	ChecarUsuarioExistente(email string) (bool, error)
}
