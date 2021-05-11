package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func configurarRutas(enrutador *mux.Router) {
	configurarRutasVehiculos(enrutador)
	configurarRutasAjustes(enrutador)
	configurarRutasPagos(enrutador)
	configurarRutasEscritorio(enrutador)
	/*
	   ____ ___  ____  ____
	  / ___/ _ \|  _ \/ ___|
	 | |  | | | | |_) \___ \
	 | |__| |_| |  _ < ___) |
	  \____\___/|_| \_\____/

	*/
	if !EstamosEnProduccion {
		// Necesitamos el método OPTIONS porque JS manda una petición antes de que se haga un DELETE
		enrutador.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", DominioPermitidoCORS)
		}).Methods(http.MethodOptions)
		enrutador.Use(middlewareCors)
	}
	if EstamosEnProduccion {
		enrutador.PathPrefix(PrefijoRutaServirContenidoEstatico).
			Handler(
				http.StripPrefix(PrefijoRutaServirContenidoEstatico,
					http.FileServer(
						http.Dir(DirectorioContenidoEstatico),
					),
				),
			)
	}
}

func middlewareCors(siguienteManejador http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", DominioPermitidoCORS)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			siguienteManejador.ServeHTTP(w, req)
		})
}
