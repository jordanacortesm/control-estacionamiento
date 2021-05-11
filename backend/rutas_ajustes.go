package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func configurarRutasAjustes(enrutador *mux.Router) {
	enrutador.HandleFunc("/ajustes_impresora", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			return obtenerImpresora()
		})
	}).Methods(http.MethodGet)

	enrutador.HandleFunc("/ajustes_impresora", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var impresora string
			err := json.NewDecoder(r.Body).Decode(&impresora)
			if err != nil {
				return nil, err
			}
			err = guardarImpresora(impresora)
			return true, err
		})
	}).Methods(http.MethodPost)

	enrutador.HandleFunc("/ajustes_costos", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			var ajuste AjusteCosto
			err := json.NewDecoder(r.Body).Decode(&ajuste)
			if err != nil {
				return nil, err
			}
			err = guardarAjustes(ajuste)
			return true, err
		})
	}).Methods(http.MethodPost)

	enrutador.HandleFunc("/ajustes_costos", func(w http.ResponseWriter, r *http.Request) {
		responderHttpConFuncion(w, r, func() (interface{}, error) {
			return obtenerAjusteCosto()
		})
	}).Methods(http.MethodGet)
}
