package main

type RecaudadoPorDia struct {
	Recaudado float64 `json:"recaudado"`
	Dia       int8    `json:"dia"`
}
type RecaudadoPorMes struct {
	Recaudado float64 `json:"recaudado"`
	Mes       int8    `json:"mes"`
}

func obtenerConteoVehiculosAtendidosEnRango(fechaInicio, fechaFin string) (int64, error) {
	bd, err := obtenerBD()
	if err != nil {
		return 0, err
	}

	defer bd.Close()
	filas, err := bd.Query(`select COUNT(*) from vehiculos where fecha_salida >= ? AND fecha_salida <= ?`, fechaInicio, fechaFin)
	if err != nil {
		return 0, err
	}
	defer filas.Close()
	var conteo int64
	for filas.Next() {
		err := filas.Scan(&conteo)
		if err != nil {
			return 0, err
		}
	}
	return conteo, nil
}

func obtenerDineroRecaudadoVehiculosAtendidosEnRango(fechaInicio, fechaFin string) (float64, error) {
	bd, err := obtenerBD()
	if err != nil {
		return 0, err
	}

	defer bd.Close()
	filas, err := bd.Query(`select coalesce(sum(pv.pago),0) from pagos_vehiculos pv
	inner join vehiculos v on v.id = pv.id_vehiculo
	and v.fecha_salida >= ? and v.fecha_salida <= ?`, fechaInicio, fechaFin)
	if err != nil {
		return 0, err
	}
	defer filas.Close()
	var recaudado float64
	for filas.Next() {
		err := filas.Scan(&recaudado)
		if err != nil {
			return 0, err
		}
	}
	return recaudado, nil
}
func obtenerDineroRecaudadoVehiculosAtendidosEnRangoPorDia(fechaInicio, fechaFin string) ([]RecaudadoPorDia, error) {
	datos := []RecaudadoPorDia{}
	bd, err := obtenerBD()
	if err != nil {
		return datos, err
	}

	defer bd.Close()
	filas, err := bd.Query(`select coalesce(sum(pv.pago),0), strftime("%d", v.fecha_salida) as dia from pagos_vehiculos pv
	inner join vehiculos v on v.id = pv.id_vehiculo
	and v.fecha_salida >= ? and v.fecha_salida <= ?
	group by dia`, fechaInicio, fechaFin)
	if err != nil {
		return datos, err
	}
	defer filas.Close()
	var recaudado RecaudadoPorDia
	for filas.Next() {
		err := filas.Scan(&recaudado.Recaudado, &recaudado.Dia)
		if err != nil {
			return datos, err
		}
		datos = append(datos, recaudado)
	}
	return datos, nil
}
func obtenerDineroRecaudadoVehiculosAtendidosEnRangoPorMes(fechaInicio, fechaFin string) ([]RecaudadoPorMes, error) {
	datos := []RecaudadoPorMes{}
	bd, err := obtenerBD()
	if err != nil {
		return datos, err
	}

	defer bd.Close()
	filas, err := bd.Query(`select coalesce(sum(pv.pago),0), strftime("%m", v.fecha_salida) as mes from pagos_vehiculos pv
	inner join vehiculos v on v.id = pv.id_vehiculo
	and v.fecha_salida >= ? and v.fecha_salida <= ?
	group by mes`, fechaInicio, fechaFin)
	if err != nil {
		return datos, err
	}
	defer filas.Close()
	var recaudado RecaudadoPorMes
	for filas.Next() {
		err := filas.Scan(&recaudado.Recaudado, &recaudado.Mes)
		if err != nil {
			return datos, err
		}
		datos = append(datos, recaudado)
	}
	return datos, nil
}
