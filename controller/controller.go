package controller

import (
	"encoding/json"
	"net/http"
	"zombie/config"
)

// APP ...
type APP struct {
	Versao    int64
	Descricao string
	Data      string
	Linguagem string
}

// Sobrevivente
type Sobreviventes struct {
	Codigo uint
	Nome   string
	Flag   string
}

var app []APP
var sobreviventes []Sobreviventes
var db = config.DB // var do banco
var versao int64
var descricao, data, linguagem string

// AtualizaAPP ...
func AtualizaAPP(v int64, d string, data string, l string) {
	app = append(app, APP{
		Versao:    v,
		Descricao: d,
		Data:      data,
		Linguagem: l,
	})
}

// HomeAPI ...
func HomeAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	versao = 1
	descricao = "Aplicação Sobrevivência Zumbi"
	data = "25/07/2018"
	linguagem = "Go"

	AtualizaAPP(versao, descricao, data, linguagem)

	json.NewEncoder(w).Encode(app)
}

// BuscarTodosSobrevivente ...
func BuscarTodosSobrevivente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sobreviventes = sobreviventes[:0]

	rows, err := db.Query("select * from sobreviventes")
	CheckError(err)

	for rows.Next() {
		sobrevivente := Sobreviventes{}
		rows.Scan(&sobrevivente.Codigo, &sobrevivente.Nome, &sobrevivente.Flag)
		sobrevivente = append(sobrevivente, Sobreviventes)
	}

}

// CheckError ...
func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
