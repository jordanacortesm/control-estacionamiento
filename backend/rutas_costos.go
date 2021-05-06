package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func configurarRutasCostos(enrutador *mux.Router) {
	enrutador.HandleFunc("/costos", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var costos []CostoPorTiempo
			err := json.NewDecoder(r.Body).Decode(&costos)
			if err != nil {
				return nil, err
			}
			err = guardarCostos(costos)
			return true, err
		})
	}).Methods(http.MethodPost)

	enrutador.HandleFunc("/costos", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			return obtenerCostos()
		})
	}).Methods(http.MethodGet)
}
