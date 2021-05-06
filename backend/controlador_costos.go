package main

type CostoPorTiempo struct {
	Id     int     `json:"id"`
	Minimo int64   `json:"minimo"`
	Maximo int64   `json:"maximo"`
	Costo  float64 `json:"costo"`
}

func guardarCostos(costos []CostoPorTiempo) error {
	err := eliminarTodosLosCostos()
	if err != nil {
		return err
	}
	bd, err := obtenerBD()
	if err != nil {
		return err
	}

	defer bd.Close()
	for _, costo := range costos {
		_, err = bd.Exec(`INSERT INTO costos_por_tiempo(minimo, maximo, costo)
	VALUES
	(?, ?, ?)`, costo.Minimo, costo.Maximo, costo.Costo)
	}
	if err != nil {
		return err
	}
	return nil
}

func eliminarTodosLosCostos() error {
	bd, err := obtenerBD()
	if err != nil {
		return err
	}

	defer bd.Close()
	_, err = bd.Exec(`DELETE FROM costos_por_tiempo`)
	return err
}

func obtenerCostos() ([]CostoPorTiempo, error) {
	costos := []CostoPorTiempo{}
	bd, err := obtenerBD()
	if err != nil {
		return costos, err
	}

	defer bd.Close()
	filas, err := bd.Query(`SELECT id, minimo, maximo, costo FROM costos_por_tiempo`)
	if err != nil {
		return costos, err
	}
	defer filas.Close()
	var costo CostoPorTiempo
	for filas.Next() {
		err := filas.Scan(&costo.Id, &costo.Minimo, &costo.Maximo, &costo.Costo)
		if err != nil {
			return costos, err
		}
		costos = append(costos, costo)
	}
	return costos, nil
}
