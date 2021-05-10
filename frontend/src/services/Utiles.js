const DIA_EN_MILISEGUNDOS = 1000 * 60 * 60 * 24;
const HORA_INICIO_DE_DIA = "00:00:00";
const HORA_FIN_DE_DIA = "23:59:59";
const formateadorFechaYHora = new Intl.DateTimeFormat('es-MX', { dateStyle: 'medium', timeStyle: 'short' });
const formateadorFecha = new Intl.DateTimeFormat('es-MX', { dateStyle: 'medium' });
const formateadorDinero = new Intl.NumberFormat("en", { style: "currency", "currency": "MXN" });
const Utiles = {
    formatearDinero(cantidad) {
        return formateadorDinero.format(cantidad);
    },
    minutosAHorasYMinutos(minutos) {
        const horas = Math.floor(minutos / 60);
        const minutosSobrantes = Math.floor(minutos % 60);
        let resultado = "";
        if (horas > 0) {
            resultado += `${horas} hora(s) `;
        }
        if (horas < 1 || minutosSobrantes > 0) {
            resultado += `${minutosSobrantes} minuto(s)`;
        }
        return resultado;
    },
    milisegundosAMinutos(milisegundos) {
        return parseInt(milisegundos / 1000 / 60);
    },
    restarFechaConFechaActual(fechaComoCadena) {
        const ahora = new Date();
        const fecha = new Date(fechaComoCadena);
        return ahora - fecha;
    },
    restarFechaComoCadenaConFechaDate(fechaComoCadena, otraFecha) {
        const fecha = new Date(fechaComoCadena);
        return otraFecha - fecha;
    },
    restarFechasComoCadenas(a, b) {
        const fechaA = new Date(a);
        const fechaB = new Date(b);
        return Math.abs(fechaA - fechaB);
    },
    obtenerFechaYHoraActual(separador) {
        return this.obtenerFechaYHora(new Date(), separador);
    },
    /*
    Regresa la fecha (debe ser una instancia de Date) formateada como
    2021-04-05 14:02:02
    El separador debería ser " " o "T", la "T" debería ser usada para guardar la fecha, y el espacio
    para mostrar la fecha al usuario
    La T es importante porque solo así es compatible con iOS
     */
    obtenerFechaYHora(fecha, separador) {
        separador = separador || " ";
        const cadenaFecha = Utiles.obtenerCadenaFecha(fecha);
        const cadenaHora = `${this.agregarCeroSiEsNecesario(fecha.getHours())}:${this.agregarCeroSiEsNecesario(fecha.getMinutes())}:${this.agregarCeroSiEsNecesario(fecha.getSeconds())}`;
        return cadenaFecha + separador + cadenaHora;
    },
    obtenerCadenaFecha(fecha) {
        const mes = fecha.getMonth() + 1; // Ya que los meses los cuenta desde el 0
        const dia = fecha.getDate();
        const cadenaFecha = `${fecha.getFullYear()}-${this.agregarCeroSiEsNecesario(mes)}-${this.agregarCeroSiEsNecesario(dia)}`;
        return cadenaFecha;
    },
    agregarCeroSiEsNecesario(valor) {
        if (valor < 10) {
            return '0'.concat(valor);
        }
        return valor.toString();
    },
    formatearFechaAInicioDeDia(fecha, separador) {
        separador = separador || "";
        return Utiles.obtenerCadenaFecha(fecha) + separador + HORA_INICIO_DE_DIA;
    },
    formatearFechaAFinDeDia(fecha, separador) {
        separador = separador || "";
        return Utiles.obtenerCadenaFecha(fecha) + separador + HORA_FIN_DE_DIA;
    },
    sumarDiasAFecha(fecha, dias) {
        return new Date(fecha.getTime() + (DIA_EN_MILISEGUNDOS * dias));
    },
    formatearFechaYHoraSegunLocale(fecha) {
        return formateadorFechaYHora.format(fecha);
    },
    formatearFechaSegunLocale(fecha) {
        return formateadorFecha.format(fecha);
    }
};
export default Utiles;