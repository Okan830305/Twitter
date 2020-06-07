package middlew

import (
	"net/http"

	"github.com/Okan830305/Twitter/bd"
)

/*ChequeoBD Midleware que permite ahcer el chequeo de la DB*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
