<template>
  <div class="columns">
    <div class="column">
      <b-field label="Placas">
        <b-input v-model="detalles.placas" placeholder="Placas"></b-input>
      </b-field>
      <b-field label="Descripción del vehículo">
        <b-input
          v-model="detalles.descripcion"
          placeholder="Descripción del vehículo"
        ></b-input>
      </b-field>

      <b-field>
        <b-switch v-model="usarFechaYHoraActual"
          >Usar fecha y hora actual ({{ fechaYHoraActual }})</b-switch
        >
      </b-field>

      <b-field label="Selecciona manualmente" v-show="!usarFechaYHoraActual">
        <b-datetimepicker
          v-model="detalles.fecha"
          rounded
          placeholder="Clic aquí para seleccionar"
          icon="calendar-today"
          locale="es-MX"
          :datetime-formatter="formatearFecha"
          :timepicker="{ enableSeconds: false, hourFormat: '24' }"
          horizontal-time-picker
        >
        </b-datetimepicker>
      </b-field>
      <b-field>
        <b-button @click="guardar()" type="is-success">Guardar</b-button>
      </b-field>
    </div>
  </div>
</template>
<script>
import Utiles from "../services/Utiles";
import DialogosService from "../services/DialogosService";
export default {
  data: () => ({
    fechaYHoraActual: null,
    usarFechaYHoraActual: true,
    detalles: {
      placas: "",
      descripcion: "",
      fecha: null,
    },
  }),
  mounted() {
    this.detalles.fecha = new Date();
    this.refrescarFechaYHoraActual();
    setInterval(this.refrescarFechaYHoraActual, 500);
  },
  methods: {
    refrescarFechaYHoraActual() {
      this.fechaYHoraActual = Utiles.obtenerFechaYHoraActual();
    },
    formatearFecha(fecha) {
      return Utiles.obtenerFechaYHora(fecha);
    },
    guardar() {
      if (!this.usarFechaYHoraActual && !this.detalles.fecha) {
        return DialogosService.mostrarNotificacionError(
          "Selecciona una fecha y hora"
        );
      }
      const cargaUtil = {
        placas: this.detalles.placas,
        descripcion: this.detalles.descripcion,
        fecha: Utiles.obtenerFechaYHora(this.detalles.fecha, "T"),
      };
      if (this.usarFechaYHoraActual) {
        cargaUtil.fecha = Utiles.obtenerFechaYHoraActual("T");
      }
      console.log({cargaUtil});
    },
  },
};
</script>