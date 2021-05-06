<template>
  <div class="columns">
    <div class="column">
      <router-link
        :to="{ name: 'RegistrarVehiculo' }"
        class="button is-success mb-2"
      >
        Registrar nuevo
      </router-link>
      <b-field
        grouped
        message="Selecciona el rango de fechas para ver registros"
      >
        <b-field label="Inicio">
          <b-datepicker
            ref="seleccionadorFechaInicio"
            :append-to-body="true"
            v-model="fechaInicio"
            @input="onFechaInicioCambiada()"
            locale="es-MX"
            placeholder="Selecciona una fecha"
            :date-formatter="formateadorFecha"
            icon="calendar-today"
            close-on-click
          >
          </b-datepicker>
        </b-field>
        <b-field label="Fin">
          <b-datepicker
            :date-formatter="formateadorFecha"
            ref="seleccionadorFechaFin"
            :append-to-body="true"
            v-model="fechaFin"
            @input="onFechaFinCambiada()"
            locale="es-MX"
            placeholder="Selecciona una fecha"
            icon="calendar-today"
            close-on-click
          >
          </b-datepicker>
        </b-field>
      </b-field>
      <b-table :data="vehiculos" :loading="cargando" :mobile-cards="true">
        <b-table-column
          searchable
          field="descripcion"
          label="Descripción"
          v-slot="props"
        >
          {{ props.row.descripcion }}
        </b-table-column>
        <b-table-column searchable field="placas" label="Placas" v-slot="props">
          {{ props.row.placas }}
        </b-table-column>
        <b-table-column field="fechaEntrada" label="Entrada" v-slot="props">
          {{ props.row.fechaEntrada | formatearFecha }}
        </b-table-column>
        <b-table-column field="fechaSalida" label="Salida" v-slot="props">
          <b-button
            @click="marcarSalida(props.row)"
            v-show="!props.row.fechaSalida"
            type="is-info"
            >Marcar salida</b-button
          >
          <template v-if="props.row.fechaSalida">
            {{ props.row.fechaSalida | formatearFecha }}
            <br />Aquí mostrar pago y tiempo
          </template>
        </b-table-column>
        <b-table-column field="id" label="Eliminar" v-slot="props">
          <b-button @click="eliminarVehiculo(props.row.id)" type="is-danger"
            >Eliminar</b-button
          >
        </b-table-column>
        <template #empty>
          <div class="has-text-centered">No hay registros</div>
        </template>
      </b-table>
    </div>
  </div>
</template>
<script>
import VehiculosService from "../services/VehiculosService";
import CostosService from "../services/CostosService";
import Utiles from "../services/Utiles";
export default {
  data: () => ({
    vehiculos: [],
    costos: [],
    cargando: false,
    fechaInicio: new Date(),
    fechaFin: new Date(),
    formateadorFecha: Utiles.formatearFechaSegunLocale,
  }),
  async mounted() {
    await this.obtenerCostos();
    await this.obtenerVehiculos();
  },
  methods: {
    async obtenerCostos() {
      this.costos = await CostosService.obtenerCostos();
    },
    onFechaInicioCambiada() {
      // Hay que ocultarlo cada que se selecciona una fecha, porque no se oculta automáticamente
      this.$refs.seleccionadorFechaInicio.toggle();
      this.onFechaCambiada();
    },
    onFechaFinCambiada() {
      // Hay que ocultarlo cada que se selecciona una fecha, porque no se oculta automáticamente
      this.$refs.seleccionadorFechaFin.toggle();
      this.onFechaCambiada();
    },
    onFechaCambiada() {
      this.obtenerVehiculos();
    },
    async eliminarVehiculo(id) {
      console.log("Eliminado el vehículo con id %d", id);
    },
    async marcarSalida(vehiculo) {
      const { fechaEntrada } = vehiculo;
      const minutosTranscurridos = Utiles.milisegundosAMinutos(
        Utiles.restarFechaConFechaActual(fechaEntrada)
      );
      const costo = CostosService.calcularCostoSegunTiempo(
        minutosTranscurridos,
        this.costos
      );
      console.log({ costo,minutosTranscurridos });
    },
    async obtenerVehiculos() {
      this.cargando = true;
      const fechaInicio = Utiles.formatearFechaAInicioDeDia(
        this.fechaInicio,
        "T"
      );
      const fechaFin = Utiles.formatearFechaAFinDeDia(this.fechaFin, "T");
      this.vehiculos = await VehiculosService.obtenerVehiculos(
        fechaInicio,
        fechaFin
      );
      this.cargando = false;
    },
  },
};
</script>