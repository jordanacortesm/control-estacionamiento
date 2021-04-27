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
          v-model="detalles.fechaEntrada"
          rounded
          placeholder="Clic aquí para seleccionar"
          icon="calendar-today"
          locale="es-MX"
          :datetime-formatter="formatearFecha"
          :timepicker="{ enableSeconds: false, hourFormat: '24' }"
          inline
        >
          <template #left>
            <b-button
              label="Ahora"
              type="is-primary"
              icon-left="clock"
              @click="detalles.fechaEntrada = new Date()"
            />
          </template>

          <template #right>
            <b-button
              label="Limpiar"
              type="is-danger"
              icon-left="close"
              outlined
              @click="detalles.fechaEntrada = null"
            />
          </template>
        </b-datetimepicker>
      </b-field>
      <b-field>
        <b-button @click="guardar()" type="is-success">Guardar</b-button>
        <router-link :to="{ name: 'Vehiculos' }" class="button is-info ml-2">
          Volver
        </router-link>
      </b-field>
    </div>
  </div>
</template>
<script>
import Utiles from "../services/Utiles";
import DialogosService from "../services/DialogosService";
import VehiculosService from "../services/VehiculosService";
export default {
  data: () => ({
    fechaYHoraActual: null,
    usarFechaYHoraActual: true,
    detalles: {
      placas: "",
      descripcion: "",
      fechaEntrada: null,
    },
  }),
  mounted() {
    this.detalles.fechaEntrada = new Date();
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
    async guardar() {
      if (!this.usarFechaYHoraActual && !this.detalles.fechaEntrada) {
        return DialogosService.mostrarNotificacionError(
          "Selecciona una fecha y hora"
        );
      }
      const cargaUtil = {
        placas: this.detalles.placas,
        descripcion: this.detalles.descripcion,
        fechaEntrada: Utiles.obtenerFechaYHora(this.detalles.fechaEntrada, "T"),
      };
      if (this.usarFechaYHoraActual) {
        cargaUtil.fechaEntrada = Utiles.obtenerFechaYHoraActual("T");
      }
      const respuesta = await VehiculosService.agregarVehiculo(cargaUtil);
      if (respuesta) {
        DialogosService.mostrarNotificacionExito("Vehículo registrado");
        this.detalles = {
          placas: "",
          descripcion: "",
          fechaEntrada: new Date(),
        };
      }
    },
  },
};
</script>