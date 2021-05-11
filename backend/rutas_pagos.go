package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func configurarRutasPagos(enrutador *mux.Router) {
	enrutador.HandleFunc("/pago", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var pago PagoDeVehiculoConFechaSalida
			err := json.NewDecoder(r.Body).Decode(&pago)
			if err != nil {
				return nil, err
			}
			err = guardarPago(pago)
			return true, err
		})
	}).Methods(http.MethodPost)

	enrutador.HandleFunc("/pagos_vehiculos", func(w http.ResponseWriter, r *http.Request) {
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
			return obtenerVehiculosConPagos(fechaInicio, fechaFin)
		})
	}).Methods(http.MethodGet)
	enrutador.HandleFunc("/pago_vehiculo", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			variablesGet := r.URL.Query()
			idVehiculoArreglo := variablesGet["id"]
			if len(idVehiculoArreglo) <= 0 {
				return nil, errors.New("no hay id de vehículo")
			}
			idVehiculo, _ := strconv.ParseInt(idVehiculoArreglo[0], 10, 64)
			return obtenerPagoDeVehiculoPorId(idVehiculo)
		})
	}).Methods(http.MethodGet)
}
