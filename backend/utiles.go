package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type RespuestaHttp struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func obtenerVariableEnteraDeUrl(variablesDePeticion map[string]string, variable string) (int64, error) {
	numero, err := strconv.ParseInt(variablesDePeticion[variable], 0, 64)
	if err != nil {
		return 0, err
	}
	return numero, err
}

func responderHttpExitoso(valor interface{}, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  valor,
		Error: "",
	})
}
func responderHttpConError(err error, w http.ResponseWriter, r *http.Request) {
	log.Printf("Error al servir respuesta para %s: %v", r.RemoteAddr, err)
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(RespuestaHttp{
		Data:  nil,
		Error: err.Error(),
	})
}

func responder(datos interface{}, err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		responderHttpConError(err, w, r)
		return
	}
	responderHttpExitoso(datos, w, r)
}
func responderHttpConFuncion(w http.ResponseWriter, r *http.Request, f func() (interface{}, error)) {
	datos, err := f()
	if err != nil {
		responderHttpConError(err, w, r)
		return
	}
	responderHttpExitoso(datos, w, r)
}
