package main

import (
	"encoding/json"
	"net/http"

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

	enrutador.HandleFunc("/costos", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			return obtenerPagosDeVehiculos()
		})
	}).Methods(http.MethodGet)
}
