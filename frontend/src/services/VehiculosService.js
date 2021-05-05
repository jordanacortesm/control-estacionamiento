import HttpService from "./HttpService";
const VehiculosService = {
    async agregarVehiculo(vehiculo) {
        return await HttpService.post("/vehiculo", vehiculo);
    },
    async obtenerVehiculos(fechaInicio, fechaFin) {
        return await HttpService.get(`/vehiculos?fechaInicio=${fechaInicio}&fechaFin=${fechaFin}`);
    },
};
export default VehiculosService;