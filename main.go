package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=go_store password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
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
