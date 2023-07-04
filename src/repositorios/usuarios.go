package repositorios

import (
	"api/src/models"
	"database/sql"
)

// Representa um repositório de usuários
type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um novo usuário no banco de dados
func (u usuarios) Criar(usuarios models.Usuario) (uint64, error) {
	return 0, nil
}
