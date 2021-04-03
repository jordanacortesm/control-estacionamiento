package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func configurarRutasVehiculos(enrutador *mux.Router) {
	enrutador.HandleFunc("/vehiculo", func(w http.ResponseWriter, r *http.Request) {
		responderHttpExitoso(true, w, r)
	}).Methods(http.MethodGet)
}
