import HttpService from "./HttpService";
const CostosService = {
    async guardarCostos(costos) {
        return await HttpService.post("/costos", costos);
    },
    async obtenerCostos() {
        return await HttpService.get("/costos");
    },
};
export default CostosService;