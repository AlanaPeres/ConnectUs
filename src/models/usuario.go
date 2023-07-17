package models //nesse pacote é armazenado as entidades de usuário e publicações e os métodos q eles podem ter

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário que utiliza a rede social
type Usuario struct {
	ID       uint64    `json: "id,omitempty"` //omitempty, se o campo estiver vazio ele não vai ser passado para o json, ele vai omitir
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "senha,omitempty"`
	CriadoEm time.Time `json: "CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
