package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("ConnectUs")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":4000", r))

}
