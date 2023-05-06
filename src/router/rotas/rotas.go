package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rotas representa a estrutura de todas as rotas da Api, todas as rotas da Api serão construídas a partir desse modelo
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {

	rotas := rotasUsuarios
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
