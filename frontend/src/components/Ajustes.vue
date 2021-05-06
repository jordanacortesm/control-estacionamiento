<template>
  <div>
    <div class="columns">
      <div class="column">
        <h2 class="is-size-2">Ajustes</h2>
        <b-field v-for="(costo, i) in costos" :key="i" grouped>
          <b-field label="Si el tiempo en minutos está entre:">
            <b-input
              v-model.number="costo.minimo"
              type="number"
              placeholder="Mínimo"
            ></b-input>
          </b-field>
          <b-field label="y:">
            <b-input
              v-model.number="costo.maximo"
              type="number"
              placeholder="Máximo"
            ></b-input>
          </b-field>
          <b-field label="El costo es:">
            <b-input
              v-model.number="costo.costo"
              type="number"
              placeholder="Costo"
            ></b-input>
          </b-field>
          <b-field label="‎">
            <b-button
              :disabled="!deberiaHabilitarBotonEliminar()"
              @click="eliminarCosto(i)"
              type="is-danger"
            >
              <b-icon icon="delete"></b-icon>
            </b-button>
          </b-field>
          <b-field v-if="deberiaMostrarBotonAgregar(i)" label="‎">
            <b-button
              @click="agregarCosto()"
              :disabled="!deberiaHabilitarBotonAgregar()"
              type="is-primary"
            >
              <b-icon icon="plus"></b-icon>
            </b-button>
          </b-field>
        </b-field>
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
    costos: [],
  }),
  async mounted() {
    await this.obtenerCostos();
  },
  methods: {
    async guardarCostos() {
      await CostosService.guardarCostos(this.costos);
      DialogosService.mostrarNotificacionExito("Costos guardados");
    },
    eliminarCosto(indice) {
      this.costos.splice(indice, 1);
    },
    agregarCosto() {
      this.costos.push({
        minimo: null,
        maximo: null,
        costo: null,
      });
    },
    deberiaMostrarBotonAgregar(indice) {
      return indice === this.costos.length - 1;
    },
    deberiaHabilitarBotonEliminar() {
      return this.costos.length > 1;
    },
    deberiaHabilitarBotonAgregar() {
      if (this.costos.length <= 0) {
        return false;
      }
      const ultimoCosto = this.costos[this.costos.length - 1];
      if (isNaN(parseInt(ultimoCosto.minimo))) {
        return false;
      }
      if (isNaN(parseInt(ultimoCosto.maximo))) {
        return false;
      }
      if (isNaN(parseFloat(ultimoCosto.costo))) {
        return false;
      }

      return true;
    },
    async obtenerCostos() {
      this.costos = await CostosService.obtenerCostos();
      // Si no hay costos, agregamos uno vacío para que se muestre el formulario
      if (this.costos.length <= 0) {
        this.agregarCosto();
      }
    },
  },
};
</script>