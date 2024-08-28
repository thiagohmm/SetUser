// domain/uow.go
package domain

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	UsuarioRepository() UsuarioRepository
}
