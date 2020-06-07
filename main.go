package main

import (
	"log"

	"github.com/Okan830305/Twitter/bd"
	"github.com/Okan830305/Twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexion a la DB")
		return
	}
	handlers.Manejadores()
}
