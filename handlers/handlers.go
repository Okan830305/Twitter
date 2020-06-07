package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Okan830305/Twitter/middlew"
	"github.com/Okan830305/Twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores stero del puerto y ecucha del servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
