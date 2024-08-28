// infrastructure/repository/usuario_repository.go
package repository

import (
	"database/sql"

	"github.com/setUserDb/domain"
)

type UsuarioRepositoryOracleTx struct {
	tx *sql.Tx
}

// ChecarUsuarioExistente implements domain.UsuarioRepository.
func (r *UsuarioRepositoryOracleTx) ChecarUsuarioExistente(email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM usr_hermes.usuario WHERE email = :1"
	err := r.tx.QueryRow(query, email).Scan(&count)
	return count > 0, err
}

func NovoUsuarioRepositoryTx(tx *sql.Tx) domain.UsuarioRepository {
	return &UsuarioRepositoryOracleTx{tx: tx}
}

func (r *UsuarioRepositoryOracleTx) ObterIDRevendedor(codigoIBM string) (int, error) {
	var idRevendedor int
	query := "SELECT IDREVENDEDOR FROM revendedor WHERE codigoibm = :1"
	err := r.tx.QueryRow(query, codigoIBM).Scan(&idRevendedor)
	return idRevendedor, err
}

func (r *UsuarioRepositoryOracleTx) InserirUsuario(usuario domain.Usuario) (int, error) {
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
	_, err := r.tx.Exec(insertSQL, usuario.Email, usuario.IDRevendedor)
	if err != nil {
		return 0, err
	}

	err = r.tx.QueryRow("SELECT NVL(MAX(ID), 0) FROM usr_hermes.usuario").Scan(&maxID)
	return maxID, err
}

func (r *UsuarioRepositoryOracleTx) InserirUsuarioAcesso(idUsuario int) error {
	insertSQL := `
		INSERT INTO usr_hermes.usuario_acessos (
			ID, IDTIPOUSUARIO, IDUSUARIO, ATIVO, DATACRIACAO, DATAALTERACAO
		) VALUES (
			:1, 2, :2, '1', SYSDATE, SYSDATE
		)
	`
	_, err := r.tx.Exec(insertSQL, idUsuario, idUsuario)
	return err
}
