package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Okan830305/Twitter/bd"
	"github.com/Okan830305/Twitter/models"
)

/*Registro crear el registro de usuario en la BD*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email Requerido", 400)
		return
	}
	if len(t.Password) < 0 {
		http.Error(w, "Password Requerido mayor a 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Usuario ya existe con ese Email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al interntar insertar "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
