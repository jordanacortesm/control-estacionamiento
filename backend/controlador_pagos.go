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
