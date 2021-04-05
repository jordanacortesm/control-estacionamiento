const RUTA_SERVIDOR = process.env.VUE_APP_RUTA_API;
const manejarRespuesta = async (respuesta) => {
    const respuestaDecodificada = await respuesta.json();
    if (respuestaDecodificada.error) {
        throw new Error(JSON.stringify(respuestaDecodificada.error));
    }
    return respuestaDecodificada.data;
};
const HttpService = {
    "post": async (ruta, datos) => {
        const respuestaRaw = await fetch(RUTA_SERVIDOR + ruta, {
            credentials: 'include',
            method: "POST",
            body: JSON.stringify(datos)
        });
        return await manejarRespuesta(respuestaRaw);
    },
    "put": async (ruta, datos) => {
        const respuestaRaw = await fetch(RUTA_SERVIDOR + ruta, {
            credentials: 'include',
            method: "PUT",
            body: JSON.stringify(datos)
        });
        return await manejarRespuesta(respuestaRaw);
    },
    "get": async (ruta) => {
        const respuestaRaw = await fetch(RUTA_SERVIDOR + ruta, {
            credentials: 'include',
        });
        return await manejarRespuesta(respuestaRaw);
    },
    "delete": async (ruta) => {
        const respuestaRaw = await fetch(RUTA_SERVIDOR + ruta, {
            credentials: 'include',
            method: 'DELETE',
        });
        return await manejarRespuesta(respuestaRaw);
    },
};
export default HttpService;