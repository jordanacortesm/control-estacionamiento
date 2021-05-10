<template>
  <div class="modal-card" style="width: auto">
    <header class="modal-card-head">
      <p class="modal-card-title">Marcar salida</p>
    </header>
    <section class="modal-card-body">
      <b-notification type="is-info">
        <p>{{ vehiculo.descripcion }} | {{ vehiculo.placas }}</p>
      </b-notification>
      <strong>Entrada: </strong>
      <span class="is-size-4">{{
        vehiculo.fechaEntrada | formatearFecha
      }}</span>
      <br />
      <strong>Transcurrido: </strong>
      <span class="is-size-4">{{
        minutosTranscurridos | minutosAHorasYMinutos
      }}</span>
      <br />
      <strong>Costo sugerido: </strong>
      <span class="is-size-4">{{ costoSugerido | dinero }}</span>
      <b-field grouped>
        <b-field
          label="Costo final"
          message="Usted puede ajustar el costo final"
        >
          <b-input
            type="number"
            v-model.number="costo"
            placeholder="Costo"
          ></b-input>
        </b-field>
        <b-field label="Pago del cliente" message="Para el cálculo del cambio">
          <b-input
            @keydown.enter.native="cobrar()"
            ref="pagoDelCliente"
            type="number"
            v-model.number="pagoDelCliente"
            placeholder="Pago del cliente"
          ></b-input> </b-field
      ></b-field>
      <strong>Cambio: </strong> {{ cambio() | dinero }}
    </section>
    <footer class="modal-card-foot">
      <b-button :disabled="!puedeCobrar()" @click="cobrar()" type="is-success"
        >Cobrar</b-button
      >
      <b-button @click="cerrar()">Cerrar</b-button>
    </footer>
  </div>
</template>
<script>
import CostosService from "../services/CostosService";
import PagosService from "../services/PagosService";
import DialogosService from "../services/DialogosService";
import Utiles from "../services/Utiles";
export default {
  props: ["vehiculo"],
  data: () => ({
    ajustesCostos: {},
    costo: null,
    costoSugerido: null,
    minutosTranscurridos: null,
    pagoDelCliente: null,
    fechaSalida: null,
  }),
  async mounted() {
    // El mounted es invocado cuando el modal se muestra desde el padre. Aquí ya podemos acceder a los detalles del vehículo
    this.$refs.pagoDelCliente.focus();
    await this.obtenerAjustesCostos();
    this.calcularCosto();
  },
  methods: {
    async cobrar() {
      if (!this.puedeCobrar()) {
        DialogosService.mostrarNotificacionError("Ingrese el pago del cliente");
        return;
      }
      const respuesta = await PagosService.guardarPago(
        this.vehiculo.id,
        this.costo,
        this.minutosTranscurridos,
        this.fechaSalida
      );
      if (respuesta) {
        DialogosService.mostrarNotificacionExito(
          "Salida registrada correctamente"
        );
        this.$emit("cobrado");
        this.cerrar();
      } else {
        DialogosService.mostrarNotificacionError("Error registrando salida");
      }
    },
    puedeCobrar() {
      return this.cambio() >= 0;
    },
    cambio() {
      return this.pagoDelCliente - this.costo;
    },
    cerrar() {
      this.$parent.close();
    },
    async obtenerAjustesCostos() {
      this.ajustesCostos = await CostosService.obtenerAjustesCostos();
    },
    calcularCosto() {
      const { fechaEntrada } = this.vehiculo;
      const fechaActual = new Date();
      const milisegundosTranscurridos = Utiles.restarFechaComoCadenaConFechaDate(
        fechaEntrada,
        fechaActual
      );
      this.fechaSalida = Utiles.obtenerFechaYHora(fechaActual, "T");
      const minutosTranscurridos = Utiles.milisegundosAMinutos(
        milisegundosTranscurridos
      );
      this.minutosTranscurridos = minutosTranscurridos;
      this.costo = this.costoSugerido = CostosService.calcularCosto(
        minutosTranscurridos,
        this.ajustesCostos.costoHora,
        this.ajustesCostos.minutosRedondear,
        this.ajustesCostos.tolerancia
      );
    },
  },
};
</script>