import HttpService from "./HttpService";
const CostosService = {
    async guardarCostos(costos) {
        return await HttpService.post("/costos", costos);
    },
    async obtenerCostos() {
        return await HttpService.get("/costos");
    },
    calcularCostoSegunTiempo(minutos, costos) {
        for (const costo of costos) {
            if (minutos >= costo.minimo && minutos <= costo.maximo) {
                return costo.costo;
            }
        }
        return 0;
    },
};
export default CostosService;