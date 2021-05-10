import HttpService from "./HttpService";

const EscritorioService = {
    async conteoVehiculos(fechaInicio, fechaFin) {
        return await HttpService.get(`/conteo_vehiculos?fechaInicio=${fechaInicio}&fechaFin=${fechaFin}`);
    },
    async recaudadoVehiculos(fechaInicio, fechaFin) {
        return await HttpService.get(`/recaudado_vehiculos?fechaInicio=${fechaInicio}&fechaFin=${fechaFin}`);
    },
    async recaudadoVehiculosAgrupado(fechaInicio, fechaFin) {
        return await HttpService.get(`/recaudado_vehiculos_agrupado?fechaInicio=${fechaInicio}&fechaFin=${fechaFin}`);
    },
    async recaudadoVehiculosAgrupadoPorMes(fechaInicio, fechaFin) {
        return await HttpService.get(`/recaudado_vehiculos_agrupado_mes?fechaInicio=${fechaInicio}&fechaFin=${fechaFin}`);
    },
};
export default EscritorioService;