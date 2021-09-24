package main

import (
	"go_store/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe("localhost:8000", nil)
}
