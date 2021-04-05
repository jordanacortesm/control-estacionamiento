import HttpService from "./HttpService";
const VehiculosService = {
    async agregarVehiculo(vehiculo) {
        return await HttpService.post("/vehiculo", vehiculo);
    }
};
export default VehiculosService;