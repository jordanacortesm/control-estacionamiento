import HttpService from "./HttpService";
const PagosService = {
    async guardarPago(idVehiculo, pago, minutos, fechaSalida) {
        const payload = {
            fechaSalida,
            pagoDeVehiculo: {
                idVehiculo, pago, minutos
            }
        };
        return await HttpService.post("/pago", payload);
    },
    async obtenerPagos(fechaInicio, fechaFin) {
        return await HttpService.get(`/pagos_vehiculos?fechaInicio=${fechaInicio}&fechaFin=${fechaFin}`);
    },
    async obtenerPagoPorIdVehiculo(idVehiculo) {
        return await HttpService.get(`/pago_vehiculo?id=${idVehiculo}`);
    },
};
export default PagosService;