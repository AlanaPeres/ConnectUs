package models //nesse pacote é armazenado as entidades de usuário e publicações e os métodos q eles podem ter

import "time"

//Usuario representa um usuário que utiliza a rede social
type Usuario struct {
	ID       uint64    `json: "id,omitempty"` //omitempty, se o campo estiver vazio ele não vai ser passado para o json, ele vai omitir
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "senha,omitempty"`
	CriadoEm time.Time `json: "CriadoEm,omitempty"`
}
