package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade  int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r * http.Request)  {

	produtos := []Produto{
		{
			Nome:       "Camiseta",
			Descricao:  "Camiseta zul",
			Preco:      39.36,
			Quantidade: 5,
		},
		{
			Nome:       "Tenis",
			Descricao:  "Confortavel",
			Preco:      80.99,
			Quantidade: 5,
		},
		{
			Nome:       "Fone",
			Descricao:  "Bem alto",
			Preco:      45.50,
			Quantidade: 8,
		},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}