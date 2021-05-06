package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func obtenerBD() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", NombreBaseDeDatos)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func crearTablas() error {
	db, err := obtenerBD()
	if err != nil {
		return err
	}
	defer db.Close()
	tablas := []string{
		`CREATE TABLE IF NOT EXISTS vehiculos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			placas TEXT NOT NULL,
			descripcion TEXT NOT NULL,
			fecha_entrada TEXT NOT NULL,
			fecha_salida TEXT NOT NULL DEFAULT ""
		);`,
		`CREATE TABLE IF NOT EXISTS costos_por_tiempo(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			minimo INTEGER NOT NULL,
			maximo INTEGER NOT NULL,
			costo REAL NOT NULL
		);`,
	}
	for _, tabla := range tablas {
		_, err = db.Exec(tabla)
		if err != nil {
			return err
		}
	}
	return nil
}
