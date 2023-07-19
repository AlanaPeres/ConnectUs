package autenticacao

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(UsuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = UsuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissoes)
	return token.SignedString([]byte("Secret"))
}
