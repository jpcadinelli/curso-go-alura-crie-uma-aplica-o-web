package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto {
		{Nome: "Camiseta", Descricao: "Azul bem bonita", Preco: 39.00, Quantidade: 10},
		{Nome: "Tenis", Descricao: "Confortável", Preco: 89.99, Quantidade: 3},
		{Nome: "Fone", Descricao: "Muito bom", Preco: 59.00, Quantidade: 2},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}