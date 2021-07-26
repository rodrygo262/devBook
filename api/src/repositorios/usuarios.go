package repositorios

import (
	"api/src/model"
	"database/sql"
	"fmt"
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

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]model.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []model.Usuario

	for linhas.Next() {
		var usuario model.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorId traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorId(id uint64) (model.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		id,
	)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario model.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return model.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Atualizar atualiza um usuário do banco de dados
func (repositorio Usuarios) Atualizar(usuarioID uint64, usuario model.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuarioID,
	); erro != nil {
		return erro
	}

	return nil
}

// Deletar apaga um usuário no banco de dados
func (repositorio Usuarios) Deletar(usuarioID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from usuarios where id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(
		usuarioID,
	); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail Busca um usuário por email e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (model.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = ?",
		email,
	)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario model.Usuario
	if linhas.Next() {
		if erro := linhas.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return model.Usuario{}, erro
		}
	}

	return usuario, nil
}
