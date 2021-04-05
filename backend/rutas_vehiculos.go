package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func configurarRutasVehiculos(enrutador *mux.Router) {
	enrutador.HandleFunc("/vehiculo", func(w http.ResponseWriter, r *http.Request) {
		responderHttpExitoso(true, w, r)
	}).Methods(http.MethodGet)
	enrutador.HandleFunc("/vehiculo", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var vehiculo Vehiculo
			err := json.NewDecoder(r.Body).Decode(&vehiculo)
			if err != nil {
				return nil, err
			}
			err = registrarNuevoVehiculo(vehiculo)
			return true, err
		})
	}).Methods(http.MethodPost)
}
