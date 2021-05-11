import HttpService from "./HttpService";

const AjustesImpresoraService = {
    async guardarImpresora(impresora) {
        return await HttpService.post("/ajustes_impresora", impresora);
    },
    async obtenerImpresora() {
        return await HttpService.get("/ajustes_impresora");
    },
};
export default AjustesImpresoraService;