package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bonita", Preco: 50.5, Quantidade: 5},
		{"Tenis", "Confortável", 98, 3},
		{"Fone", "Razr", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
