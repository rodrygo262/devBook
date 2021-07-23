package repositorios

import (
	"api/src/model"
	"database/sql"
)

// Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario model.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios(nome, nick, email, senha)  values ( ?, ? , ?, ? )",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuario.Senha,
	)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}
