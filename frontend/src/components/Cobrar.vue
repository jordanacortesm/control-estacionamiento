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
      <span class="is-size-4">{{ costo | dinero }}</span>
      <b-field grouped>
        <b-field
          label="Costo final"
          message="Usted puede ajustar el costo final"
        >
          <b-input type="number" v-model="costo" placeholder="Costo"></b-input>
        </b-field>
        <b-field label="Pago del cliente" message="Para el cálculo del cambio">
          <b-input
            @keydown.enter.native="cobrar()"
            ref="pagoDelCliente"
            type="number"
            v-model="pagoDelCliente"
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
import DialogosService from "../services/DialogosService";
import Utiles from "../services/Utiles";
export default {
  props: ["vehiculo"],
  data: () => ({
    costo: null,
    minutosTranscurridos: null,
    pagoDelCliente: null,
  }),
  async mounted() {
    // El mounted es invocado cuando el modal se muestra desde el padre. Aquí ya podemos acceder a los detalles del vehículo
    this.$refs.pagoDelCliente.focus();
    await this.obtenerCostos();
    this.calcularCosto();
  },
  methods: {
    cobrar() {
      if (!this.puedeCobrar()) {
        DialogosService.mostrarNotificacionError("Ingrese el pago del cliente");
        return;
      }
      console.log("Se cobra");
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
    async obtenerCostos() {
      this.costos = await CostosService.obtenerCostos();
    },
    calcularCosto() {
      const { fechaEntrada } = this.vehiculo;
      const minutosTranscurridos = Utiles.milisegundosAMinutos(
        Utiles.restarFechaConFechaActual(fechaEntrada)
      );
      this.minutosTranscurridos = minutosTranscurridos;
      this.costo = CostosService.calcularCostoSegunTiempo(
        minutosTranscurridos,
        this.costos
      );
    },
  },
};
</script>