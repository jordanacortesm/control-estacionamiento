<template>
  <div>
    <div class="columns">
      <div class="column">
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
        <h3 class="is-size-3">Total: {{ total() | dinero }}</h3>
        <b-table
          :data="vehiculos"
          :loading="cargando"
          :mobile-cards="true"
          hoverable
        >
          <b-table-column
            searchable
            field="vehiculo.descripcion"
            label="Descripción"
            v-slot="props"
            sortable
          >
            {{ props.row.vehiculo.descripcion }}
          </b-table-column>
          <b-table-column
            searchable
            field="vehiculo.placas"
            sortable
            label="Placas"
            v-slot="props"
          >
            {{ props.row.vehiculo.placas }}
          </b-table-column>
          <b-table-column
            field="vehiculo.fechaEntrada"
            label="Entrada"
            v-slot="props"
            sortable
          >
            {{ props.row.vehiculo.fechaEntrada | formatearFecha }}
          </b-table-column>
          <b-table-column
            field="vehiculo.fechaSalida"
            label="Salida"
            v-slot="props"
            sortable
          >
            {{ props.row.vehiculo.fechaSalida | formatearFecha }}
          </b-table-column>
          <b-table-column
            field="pagoDeVehiculo.minutos"
            label="Tiempo"
            v-slot="props"
            sortable
          >
            {{ props.row.pagoDeVehiculo.minutos | minutosAHorasYMinutos }}
          </b-table-column>
          <b-table-column
            field="pagoDeVehiculo.pago"
            label="Pago"
            v-slot="props"
            sortable
          >
            {{ props.row.pagoDeVehiculo.pago | dinero }}
          </b-table-column>
          <template #empty>
            <div class="has-text-centered">No hay registros</div>
          </template>
        </b-table>
      </div>
    </div>
  </div>
</template>
<script>
import PagosService from "../services/PagosService";
import Utiles from "../services/Utiles";
export default {
  data: () => ({
    formateadorFecha: Utiles.formatearFechaSegunLocale,
    fechaFin: new Date(),
    fechaInicio: new Date(),
    vehiculos: [],
    cargando: false,
  }),
  async mounted() {
    await this.obtenerReporteConFechasSeleccionadas();
  },
  methods: {
    total() {
      let total = 0;
      for (const vehiculo of this.vehiculos) {
        total += vehiculo.pagoDeVehiculo.pago;
      }
      return total;
    },
    onFechaFinCambiada() {
      this.$refs.seleccionadorFechaFin.toggle();
      this.onFechasCambiadas();
    },
    onFechaInicioCambiada() {
      this.$refs.seleccionadorFechaInicio.toggle();
      this.onFechasCambiadas();
    },
    onFechasCambiadas() {
      this.obtenerReporteConFechasSeleccionadas();
    },
    async obtenerReporteConFechasSeleccionadas() {
      this.cargando = true;
      const fechaInicio = Utiles.formatearFechaAInicioDeDia(
        this.fechaInicio,
        "T"
      );
      const fechaFin = Utiles.formatearFechaAFinDeDia(this.fechaFin, "T");
      this.vehiculos = await PagosService.obtenerPagos(fechaInicio, fechaFin);
      this.cargando = false;
    },
  },
};
</script>