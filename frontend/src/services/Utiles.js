const Utiles = {
    restarHorarios(a, b) {
        const fecha = this.formatearFechaActual();
        const fechaYHoraA = new Date(fecha + " " + a);
        const fechaYHoraB = new Date(fecha + " " + b);
        return fechaYHoraA - fechaYHoraB;
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
        const mes = fecha.getMonth() + 1; // Ya que los meses los cuenta desde el 0
        const dia = fecha.getDate();
        const cadenaFecha = `${fecha.getFullYear()}-${this.agregarCeroSiEsNecesario(mes)}-${this.agregarCeroSiEsNecesario(dia)}`;
        const cadenaHora = `${this.agregarCeroSiEsNecesario(fecha.getHours())}:${this.agregarCeroSiEsNecesario(fecha.getMinutes())}:${this.agregarCeroSiEsNecesario(fecha.getSeconds())}`;
        return cadenaFecha + separador + cadenaHora;
    },
    agregarCeroSiEsNecesario(valor) {
        if (valor < 10) {
            return '0'.concat(valor);
        }
        return valor.toString();
    }
};
export default Utiles;