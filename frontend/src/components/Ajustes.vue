<template>
  <div>
    <div class="columns">
      <div class="column">
        <h2 class="is-size-2">Ajustes</h2>
        <b-field grouped>
          <b-field label="Costo por hora">
            <b-input
              placeholder="Ingrese el costo por hora"
              type="number"
              v-model.number="ajustes.costoHora"
            ></b-input>
          </b-field>
          <b-field label="Redondear">
            <b-select v-model.number="ajustes.minutosRedondear">
              <option
                :value="opcion"
                v-for="(opcion, i) in opcionesRedondeo"
                :key="i"
              >
                Próximos {{ opcion }} minuto(s)
              </option>
            </b-select>
          </b-field>
          <b-field label="Tolerancia en minutos:">
            <b-input
              type="number"
              v-model.number="ajustes.tolerancia"
              placeholder="Tolerancia en minutos"
            ></b-input>
          </b-field>
        </b-field>
        <b-field label="Puede mover el deslizador para ver una simulación">
          <b-slider
            type="is-info"
            v-model="minutosSimulador"
            :min="0"
            :max="300"
          ></b-slider>
        </b-field>
        <b-notification type="is-primary">
          Según los ajustes, por
          <strong>{{ minutosSimulador | minutosAHorasYMinutos }}</strong> el
          costo es de
          <strong>{{ costoSimulacion() | dinero }}</strong>
        </b-notification>
        <b-button
          :disabled="!deberiaHabilitarBotonAgregar()"
          type="is-success"
          @click="guardarCostos()"
          >Guardar</b-button
        >
      </div>
    </div>
  </div>
</template>
<script>
import CostosService from "../services/CostosService";
import DialogosService from "../services/DialogosService";
export default {
  data: () => ({
    opcionesRedondeo: [0, 10, 15, 30, 60],
    ajustes: {
      costoHora: null,
      minutosRedondear: null,
      tolerancia: null,
    },
    minutosSimulador: 60,
  }),
  async mounted() {
    this.ajustes.minutosRedondear = this.opcionesRedondeo[0];
    await this.obtenerAjustesCostos();
  },
  methods: {
    costoSimulacion() {
      const costo = CostosService.calcularCosto(
        this.minutosSimulador,
        this.ajustes.costoHora,
        this.ajustes.minutosRedondear,
        this.ajustes.tolerancia
      );
      return costo;
    },

    async guardarCostos() {
      await CostosService.guardarAjustesCostos(
        this.ajustes.costoHora,
        this.ajustes.minutosRedondear,
        this.ajustes.tolerancia
      );
      DialogosService.mostrarNotificacionExito("Costos guardados");
    },
    deberiaHabilitarBotonAgregar() {
      return true;
    },
    async obtenerAjustesCostos() {
      this.ajustes = await CostosService.obtenerAjustesCostos();
    },
  },
};
</script>