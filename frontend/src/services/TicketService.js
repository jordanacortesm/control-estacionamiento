import Impresora from "./Impresora";
import Utiles from "./Utiles";
const TicketService = {
    async imprimirTicketPrueba(impresora){
        const i = new Impresora();
        i.setAlign("center");
        i.write("Si puede leer esto, ha configurado correctamente su impresora. No olvide guardar los ajustes :)");
        i.feed(1);
        i.imprimirEnImpresora(impresora);
    },
    async obtenerImpresoras() {
        return await Impresora.getImpresoras();
    },
    async imprimirTicketEntrada(vehiculo, nombreImpresora) {
        const i = new Impresora();
        i.setAlign("center");
        i.setFontSize(1, 2);
        i.write("Comprobante de ingreso\n\n");
        i.setFontSize(1, 1);
        i.setAlign("left");
        i.setEmphasize(1);
        i.write("ID: ");
        i.setEmphasize(0);
        i.write(vehiculo.id);
        i.feed(1);
        i.setEmphasize(1);
        i.write("Placas: ");
        i.setEmphasize(0);
        i.write(vehiculo.placas);
        i.feed(1);
        i.setEmphasize(1);
        i.write("Entrada: ");
        i.setEmphasize(0);
        i.write(Utiles.formatearFechaYHoraSegunLocale(new Date(vehiculo.fechaEntrada)));
        i.feed(2);
        i.setFontSize(1, 2);
        i.setAlign("center");
        i.write("Por favor, conserve este comprobante")
        i.setFontSize(1, 1);
        i.feed(3);
        i.cut();
        i.cutPartial();
        return await i.imprimirEnImpresora(nombreImpresora);
    },
    async imprimirTicketSalida(vehiculo, nombreImpresora, minutos, costo) {
        const i = new Impresora();
        i.setAlign("center");
        i.setFontSize(1, 2);
        i.write("Comprobante de pago\n\n");
        i.setFontSize(1, 1);
        i.setAlign("left");
        i.setEmphasize(1);
        i.write("ID: ");
        i.setEmphasize(0);
        i.write(vehiculo.id);
        i.feed(1);
        i.setEmphasize(1);
        i.write("Placas: ");
        i.setEmphasize(0);
        i.write(vehiculo.placas);
        i.feed(1);
        i.setEmphasize(1);
        i.write("Entrada: ");
        i.setEmphasize(0);
        i.write(Utiles.formatearFechaYHoraSegunLocale(new Date(vehiculo.fechaEntrada)));
        i.feed(1);
        i.setEmphasize(1);
        i.write("Salida: ");
        i.setEmphasize(0);
        i.write(Utiles.formatearFechaYHoraSegunLocale(new Date(vehiculo.fechaSalida)));
        i.feed(1);
        i.setEmphasize(1);
        i.write("Tiempo: ");
        i.setEmphasize(0);
        i.write(Utiles.minutosAHorasYMinutos(minutos));
        i.feed(2);
        i.setEmphasize(1);
        i.setAlign("right");
        i.setFontSize(2, 1);
        i.write("Total:\n");
        i.setEmphasize(0);
        i.write(Utiles.formatearDinero(costo));
        i.feed(2);
        i.setFontSize(1, 2);
        i.setAlign("center");
        i.write("Gracias por su preferencia");
        i.setFontSize(1, 1);
        i.feed(3);
        i.cut();
        i.cutPartial();
        i.cash();
        return await i.imprimirEnImpresora(nombreImpresora);
    },
};
export default TicketService;