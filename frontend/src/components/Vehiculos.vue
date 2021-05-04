<template>
  <div class="columns">
    <div class="column">
      <router-link
        :to="{ name: 'RegistrarVehiculo' }"
        class="button is-success mb-2"
      >
        Registrar nuevo
      </router-link>
      <b-field label="Filtrar por fecha">
        <b-datepicker
          v-model="fechaSeleccionada"
          locale="es-MX"
          placeholder="Selecciona una fecha"
          icon="calendar-today"
          close-on-click
        >
        </b-datepicker>
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
          <b-button v-show="!props.row.fechaSalida" type="is-info"
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
export default {
  data: () => ({
    vehiculos: [],
    cargando: false,
    fechaSeleccionada: new Date(),
  }),
  async mounted() {
    // this.fechaSeleccionada = new Date();
    await this.obtenerVehiculos();
  },
  methods: {
    async eliminarVehiculo(id) {
      console.log("Eliminado el vehículo con id %d", id);
    },
    async obtenerVehiculos() {
      this.cargando = true;
      this.vehiculos = await VehiculosService.obtenerVehiculos();
      this.cargando = false;
    },
  },
};
</script>