package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em Json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um erro em formato em JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro`
	}{
		Erro: erro.Error(),
	})

}