package main

type PagoDeVehiculo struct {
	Id         int64   `json:"id"`
	IdVehiculo int64   `json:"idVehiculo"`
	Pago       float64 `json:"pago"`
	Minutos    int64   `json:"minutos"`
}

// Struct para recibir el JSON que se envía del cliente para poder guardar la fecha de salida además del pago
type PagoDeVehiculoConFechaSalida struct {
	PagoDeVehiculo PagoDeVehiculo `json:"pagoDeVehiculo"`
	FechaSalida    string         `json:"fechaSalida"`
}

type VehiculoConPago struct {
	Vehiculo       Vehiculo       `json:"vehiculo"`
	PagoDeVehiculo PagoDeVehiculo `json:"pagoDeVehiculo"`
}

func obtenerVehiculosConPagos(fechaInicio, fechaFin string) ([]VehiculoConPago, error) {
	vehiculos := []VehiculoConPago{}
	bd, err := obtenerBD()
	if err != nil {
		return vehiculos, err
	}

	defer bd.Close()
	filas, err := bd.Query(`select v.placas, v.descripcion, v.fecha_entrada, v.fecha_salida, pv.pago, pv.minutos from pagos_vehiculos pv
inner join vehiculos v
on v.id = pv.id_vehiculo
where v.fecha_entrada >= ?
AND v.fecha_entrada <= ?`, fechaInicio, fechaFin)
	if err != nil {
		return vehiculos, err
	}
	defer filas.Close()
	var v VehiculoConPago
	for filas.Next() {
		err := filas.Scan(&v.Vehiculo.Placas, &v.Vehiculo.Descripcion, &v.Vehiculo.FechaEntrada, &v.Vehiculo.FechaSalida, &v.PagoDeVehiculo.Pago, &v.PagoDeVehiculo.Minutos)
		if err != nil {
			return vehiculos, err
		}
		vehiculos = append(vehiculos, v)
	}
	return vehiculos, nil
}

func guardarPago(pago PagoDeVehiculoConFechaSalida) error {
	err := establecerFechaSalida(pago.PagoDeVehiculo.IdVehiculo, pago.FechaSalida)
	if err != nil {
		return err
	}
	bd, err := obtenerBD()
	if err != nil {
		return err
	}

	defer bd.Close()
	_, err = bd.Exec(`INSERT INTO pagos_vehiculos(id_vehiculo, pago, minutos)
	VALUES
	(?, ?, ?)`, pago.PagoDeVehiculo.IdVehiculo, pago.PagoDeVehiculo.Pago, pago.PagoDeVehiculo.Minutos)
	if err != nil {
		return err
	}
	return nil
}

func obtenerPagosDeVehiculos() ([]PagoDeVehiculo, error) {
	pagos := []PagoDeVehiculo{}
	bd, err := obtenerBD()
	if err != nil {
		return pagos, err
	}

	defer bd.Close()
	filas, err := bd.Query(`SELECT id, id_vehiculo, pago, minutos FROM pagos_vehiculos`)
	if err != nil {
		return pagos, err
	}
	defer filas.Close()
	var pago PagoDeVehiculo
	for filas.Next() {
		err := filas.Scan(&pago.Id, &pago.IdVehiculo, &pago.Pago, &pago.Minutos)
		if err != nil {
			return pagos, err
		}
		pagos = append(pagos, pago)
	}
	return pagos, nil
}
