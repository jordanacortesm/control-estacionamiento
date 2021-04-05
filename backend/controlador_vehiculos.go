package main

type Vehiculo struct {
	Descripcion  string `json:"descripcion"`
	Placas       string `json:"placas"`
	FechaEntrada string `json:"fechaEntrada"`
	FechaSalida  string `json:"fechaSalida"`
}

func registrarNuevoVehiculo(vehiculo Vehiculo) error {
	bd, err := obtenerBD()
	if err != nil {
		return err
	}

	defer bd.Close()
	_, err = bd.Exec(`INSERT INTO vehiculos(descripcion,placas,fecha_entrada)
	VALUES
	(?, ?, ?)`, vehiculo.Descripcion, vehiculo.Placas, vehiculo.FechaEntrada)
	if err != nil {
		return err
	}
	return nil
}
