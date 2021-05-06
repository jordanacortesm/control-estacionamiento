package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func configurarRutasVehiculos(enrutador *mux.Router) {
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
	enrutador.HandleFunc("/vehiculos", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			variablesGet := r.URL.Query()
			fechaInicioArreglo := variablesGet["fechaInicio"]
			fechaFinArreglo := variablesGet["fechaFin"]
			if len(fechaInicioArreglo) <= 0 {
				return nil, errors.New("no hay fecha inicio")
			}
			if len(fechaFinArreglo) <= 0 {
				return nil, errors.New("no hay fecha de fin")
			}
			fechaInicio, fechaFin := fechaInicioArreglo[0], fechaFinArreglo[0]
			return obtenerVehiculos(fechaInicio, fechaFin)
		})
	}).Methods(http.MethodGet)
}
