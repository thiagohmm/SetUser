// infrastructure/repository/uow.go
package repository

import (
	"database/sql"

	"github.com/setUserDb/domain"
)

type UnitOfWorkOracle struct {
	tx          *sql.Tx
	usuarioRepo domain.UsuarioRepository
}

func NovoUnitOfWork(db *sql.DB) (*UnitOfWorkOracle, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &UnitOfWorkOracle{
		tx:          tx,
		usuarioRepo: NovoUsuarioRepositoryTx(tx),
	}, nil
}

func (uow *UnitOfWorkOracle) Begin() error {
	// A transação já foi iniciada em NovoUnitOfWork
	return nil
}

func (uow *UnitOfWorkOracle) Commit() error {
	return uow.tx.Commit()
}

func (uow *UnitOfWorkOracle) Rollback() error {
	return uow.tx.Rollback()
}

func (uow *UnitOfWorkOracle) UsuarioRepository() domain.UsuarioRepository {
	return uow.usuarioRepo
}
