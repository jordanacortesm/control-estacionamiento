<template>
  <div>
    <div class="columns">
      <div class="column">
        <h2 class="is-size-2">Costos</h2>
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
          >Guardar costo</b-button
        >
        <hr />
        <h2 class="is-size-2">Impresión</h2>
        <b-field grouped>
          <b-field label="Impresora para tickets">
            <b-select
              v-model="impresoraSeleccionada"
              :loading="cargandoImpresoras"
            >
              <option
                v-for="(impresora, i) in impresoras"
                :value="impresora"
                :key="i"
              >
                {{ impresora }}
              </option>
            </b-select>
            <p class="control">
              <b-button @click="probarConImpresoraSeleccionada()" type="is-info"
                >Probar</b-button
              >
            </p>
          </b-field>
        </b-field>
        <b-notification type="is-warning">
          Es normal que aparezca una ventana negra por unos segundos
          <br />
          Recuerde que el plugin se debe estar ejecutando para que la lista sea
          obtenida. Más información y descarga
          <a href="https://parzibyte.me/plugin-impresora-termica-v1/">aquí</a>.
          <br />
          <strong>Nota:</strong> solo seleccione impresoras térmicas y
          compartidas desde el panel de control. No seleccione impresoras
          virtuales o de tinta
        </b-notification>
        <b-button
          :disabled="!deberiaHabilitarBotonGuardarImpresora()"
          type="is-success"
          @click="guardarImpresora()"
          >Guardar impresora</b-button
        >
      </div>
    </div>
  </div>
</template>
<script>
import CostosService from "../services/CostosService";
import DialogosService from "../services/DialogosService";
import TicketService from "../services/TicketService";
import AjustesImpresoraService from "../services/AjustesImpresoraService";
export default {
  data: () => ({
    impresoraSeleccionada: null,
    cargandoImpresoras: false,
    impresoras: [],
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
    await this.obtenerListaImpresoras();
    await this.obtenerImpresora();
  },
  methods: {
    async probarConImpresoraSeleccionada() {
      try {
        await TicketService.imprimirTicketPrueba(this.impresoraSeleccionada);
        DialogosService.mostrarNotificacionExito(
          "Si la impresora está conectada, debería haber impreso un ticket"
        );
      } catch (e) {
        DialogosService.mostrarNotificacionError(
          "Error imprimiendo. Asegúrese de haber descargado y ejecutado el plugin. Recuerde seleccionar impresoras térmicas y compartidas desde el Panel de control"
        );
      }
    },
    async guardarImpresora() {
      await AjustesImpresoraService.guardarImpresora(
        this.impresoraSeleccionada
      );
      DialogosService.mostrarNotificacionExito("Impresora guardada");
    },
    async obtenerImpresora() {
      this.impresoraSeleccionada = await AjustesImpresoraService.obtenerImpresora();
    },
    deberiaHabilitarBotonGuardarImpresora() {
      return !this.cargandoImpresoras && this.impresoraSeleccionada;
    },
    async obtenerListaImpresoras() {
      this.cargandoImpresoras = true;
      try {
        this.impresoras = await TicketService.obtenerImpresoras();
      } catch (e) {
        DialogosService.mostrarNotificacionError(
          "No se pudo obtener la lista de impresoras. Asegúrese de que el plugin está ejecutándose"
        );
      } finally {
        this.cargandoImpresoras = false;
      }
    },
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