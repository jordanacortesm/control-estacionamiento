package main

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func configurarRutasEscritorio(enrutador *mux.Router) {

	enrutador.HandleFunc("/conteo_vehiculos", func(w http.ResponseWriter, r *http.Request) {
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
			return obtenerConteoVehiculosAtendidosEnRango(fechaInicio, fechaFin)
		})
	}).Methods(http.MethodGet)

	enrutador.HandleFunc("/recaudado_vehiculos", func(w http.ResponseWriter, r *http.Request) {
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
			return obtenerDineroRecaudadoVehiculosAtendidosEnRango(fechaInicio, fechaFin)
		})
	}).Methods(http.MethodGet)

	enrutador.HandleFunc("/recaudado_vehiculos_agrupado", func(w http.ResponseWriter, r *http.Request) {
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
			return obtenerDineroRecaudadoVehiculosAtendidosEnRangoPorDia(fechaInicio, fechaFin)
		})
	}).Methods(http.MethodGet)
	enrutador.HandleFunc("/recaudado_vehiculos_agrupado_mes", func(w http.ResponseWriter, r *http.Request) {
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
			return obtenerDineroRecaudadoVehiculosAtendidosEnRangoPorMes(fechaInicio, fechaFin)
		})
	}).Methods(http.MethodGet)
}
