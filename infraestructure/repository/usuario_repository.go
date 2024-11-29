package repository

import (
	"database/sql"

	"github.com/setUserDb/domain"
)

type UsuarioRepositoryOracleTx struct {
	tx *sql.Tx
}

// ChecarUsuarioExistente verifica se um usuário com o email especificado já existe.
func (r *UsuarioRepositoryOracleTx) ChecarUsuarioExistente(email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM usr_hermes.usuario WHERE email = :1"
	err := r.tx.QueryRow(query, email).Scan(&count)
	return count > 0, err
}

// NovoUsuarioRepositoryTx cria uma nova instância de UsuarioRepositoryOracleTx.
func NovoUsuarioRepositoryTx(tx *sql.Tx) domain.UsuarioRepository {
	return &UsuarioRepositoryOracleTx{tx: tx}
}

// ObterIDRevendedor retorna o ID do revendedor baseado no código IBM.
func (r *UsuarioRepositoryOracleTx) ObterIDRevendedor(codigoIBM string) (int, error) {
	var idRevendedor int
	query := "SELECT IDREVENDEDOR FROM revendedor WHERE codigoibm = :1"
	err := r.tx.QueryRow(query, codigoIBM).Scan(&idRevendedor)
	return idRevendedor, err
}

// InserirUsuario insere um novo usuário na tabela usr_hermes.usuario e retorna o ID gerado.
func (r *UsuarioRepositoryOracleTx) InserirUsuario(usuario domain.Usuario) (int, error) {
	insertSQL := `
        INSERT INTO  usuario (
            EMAIL, IDREVENDEDOR, ATIVO, DATACRIACAO, DATAALTERACAO
        ) VALUES (
            :1, :2, '1', SYSDATE, SYSDATE
        )
    `
	_, err := r.tx.Exec(insertSQL, usuario.Email, usuario.IDRevendedor)
	if err != nil {
		return 0, err
	}

	// Obter o ID do usuário recém-inserido
	var idUsuario int
	err = r.tx.QueryRow("SELECT MAX(ID) FROM usuario WHERE email = :1", usuario.Email).Scan(&idUsuario)
	if err != nil {
		return 0, err
	}

	return idUsuario, nil
}

// InserirUsuarioAcesso insere um novo acesso de usuário na tabela usuario_acessos.
func (r *UsuarioRepositoryOracleTx) InserirUsuarioAcesso(idUsuario int, permissao int) error {
	var insertSQL string

	if permissao == 0 {

		insertSQL = `
        INSERT INTO usuario_acessos (
            IDTIPOUSUARIO, IDUSUARIO, ATIVO, DATACRIACAO, DATAALTERACAO
        ) VALUES (
            2, :1, '1', SYSDATE, SYSDATE
        )
    `
	} else {
		insertSQL = `
        INSERT INTO usuario_acessos (
            IDTIPOUSUARIO, IDUSUARIO, ATIVO, DATACRIACAO, DATAALTERACAO
        ) VALUES (
            1, :1, '1', SYSDATE, SYSDATE
        )
    `
	}

	stmt, err := r.tx.Prepare(insertSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(idUsuario)
	return err
}
