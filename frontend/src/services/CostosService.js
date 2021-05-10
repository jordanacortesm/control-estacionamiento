const MINUTOS_EN_UNA_HORA = 60;
import HttpService from "./HttpService";

const CostosService = {
    async guardarAjustesCostos(costoHora, minutosRedondear, tolerancia) {
        const payload = { costoHora, minutosRedondear, tolerancia };
        return await HttpService.post("/ajustes_costos", payload);
    },
    async obtenerAjustesCostos() {
        return await HttpService.get("/ajustes_costos");
    },
    calcularCosto(minutos, costoPorHora, minutosRedondear, tolerancia) {
        let minutosVerdaderos = minutos;
        let diferencia = 0;
        if (minutosRedondear > 0) {
            diferencia = (minutos % minutosRedondear);
        }
        if (diferencia > tolerancia) {
            minutosVerdaderos = minutosVerdaderos - diferencia + minutosRedondear;
        } else {
            minutosVerdaderos = minutosVerdaderos - diferencia;
        }
        if (minutosVerdaderos < minutosRedondear) {
            minutosVerdaderos = minutosRedondear;
        }
        return (minutosVerdaderos / MINUTOS_EN_UNA_HORA) * costoPorHora;
    },
};
export default CostosService;