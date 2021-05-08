import HttpService from "./HttpService";
const PagosService = {
    async guardarPago(idVehiculo, pago, minutos, fechaSalida) {
        const payload = {
            fechaSalida,
            pagoDeVehiculo:{
                idVehiculo,pago,minutos
            }
        };
        return await HttpService.post("/pago", payload);
    },
    async obtenerPagos() {
        return await HttpService.get("/pagos");
    },
};
export default PagosService;