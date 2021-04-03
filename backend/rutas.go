package main

import (
	"github.com/gorilla/mux"
)

func configurarRutas(enrutador *mux.Router) {
	configurarRutasVehiculos(enrutador)
}
