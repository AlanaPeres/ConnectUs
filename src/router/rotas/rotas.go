package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rotas representa a estrutura de todas as rotas da Api, todas as rotas da Api serão construídas a partir desse modelo
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request) //recebe uma requisição e retorna uma resposta
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
// recebe um router q não tem nenhuma rota dentro, vai configurar todas as rotas utilizando o handlefunc e devolver o router pronto.
func Configurar(r *mux.Router) *mux.Router {

	rotas := rotasUsuarios
	//itero por cada uma das rotas e dou um HandleFunc nelas passando a Uri, a funçao e o método.
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
