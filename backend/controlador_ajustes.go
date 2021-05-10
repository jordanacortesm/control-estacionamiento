package main

type AjusteCosto struct {
	CostoHora        float64 `json:"costoHora"`
	MinutosRedondear int64   `json:"minutosRedondear"`
	Tolerancia       int64   `json:"tolerancia"`
}

func guardarAjustes(ajuste AjusteCosto) error {
	err := eliminarAjustesCosto()
	if err != nil {
		return err
	}
	bd, err := obtenerBD()
	if err != nil {
		return err
	}

	defer bd.Close()
	_, err = bd.Exec(`INSERT INTO ajustes_costo(costo_hora, minutos_redondear, minutos_tolerancia)
	VALUES
	(?, ?, ?)`, ajuste.CostoHora, ajuste.MinutosRedondear, ajuste.Tolerancia)
	if err != nil {
		return err
	}
	return nil
}

func eliminarAjustesCosto() error {
	bd, err := obtenerBD()
	if err != nil {
		return err
	}

	defer bd.Close()
	_, err = bd.Exec(`DELETE FROM ajustes_costo`)
	return err
}

func obtenerAjusteCosto() (AjusteCosto, error) {
	var ajuste AjusteCosto
	bd, err := obtenerBD()
	if err != nil {
		return ajuste, err
	}

	defer bd.Close()
	filas, err := bd.Query(`SELECT costo_hora, minutos_redondear, minutos_tolerancia FROM ajustes_costo`)
	if err != nil {
		return ajuste, err
	}
	defer filas.Close()
	for filas.Next() {
		err := filas.Scan(&ajuste.CostoHora, &ajuste.MinutosRedondear, &ajuste.Tolerancia)
		if err != nil {
			return ajuste, err
		}
	}
	return ajuste, nil
}
