package routes

import (
	"fmt"
	"log"
	"net/http"
	"sobrevivenciaZumbie/config"
	"sobrevivenciaZumbie/controller"

	"github.com/gorilla/mux"
)

var portaAplicacao string

func HandleFunc() {
	rotas := mux.NewRouter()
	config.TryConn()
	portaAplicacao = ":3000"
	fmt.Println("Aplicação ON: porta => ", portaAplicacao)

	rotas.HandleFunc("/api/", controller.HomeAPI).Methods("GET")

	log.Fatal(http.ListenAndServe(portaAplicacao, rotas))
}
