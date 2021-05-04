import HttpService from "./HttpService";
const VehiculosService = {
    async agregarVehiculo(vehiculo) {
        return await HttpService.post("/vehiculo", vehiculo);
    },
    async obtenerVehiculos() {
        return await HttpService.get("/vehiculos");
    },
};
export default VehiculosService;