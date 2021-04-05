package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	err := crearTablas()
	if err != nil {
		log.Printf("Error creando tablas: %v", err)
		return
	}
	enrutador := mux.NewRouter()
	configurarRutas(enrutador)

	servidor := &http.Server{
		Handler: enrutador,
		Addr:    PuertoHttp,
		// Timeouts para evitar que el servidor se quede "colgado" por siempre
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	ruta := "http://localhost" + PuertoHttp
	fmt.Printf("Escuchando en %s\n", ruta)
	log.Fatal(servidor.ListenAndServe())
}
