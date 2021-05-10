<template>
  <div>
    <div class="columns">
      <div class="column">
        <div class="card">
          <header class="card-header">
            <p class="card-header-title">Hoy</p>
          </header>
          <div class="card-content">
            <div class="content">
              <h2 class="is-size-3">
                <b-icon icon="cash" size="is-large"></b-icon>
                {{ estadisticasHoy.recaudado | dinero }}
              </h2>

              <h2 class="is-size-4">
                <b-icon icon="car" size="is-medium"></b-icon>
                {{ estadisticasHoy.conteo }} vehículo(s)
              </h2>
            </div>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="card">
          <header class="card-header">
            <p class="card-header-title">Este mes</p>
          </header>
          <div class="card-content">
            <div class="content">
              <h2 class="is-size-3">
                <b-icon icon="cash" size="is-large"></b-icon>
                {{ estadisticasMes.recaudado | dinero }}
              </h2>
              <h2 class="is-size-4">
                <b-icon icon="car" size="is-medium"></b-icon>
                {{ estadisticasMes.conteo }} vehículo(s)
              </h2>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column">
        <div class="card">
          <header class="card-header">
            <p class="card-header-title">Recaudado este mes</p>
          </header>
          <div class="card-content"></div>
        </div>
      </div>
      <div class="column">
        <div class="card">
          <header class="card-header">
            <p class="card-header-title">Recaudado este año</p>
          </header>
          <div class="card-content"></div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import EscritorioService from "../services/EscritorioService";
import Utiles from "../services/Utiles";
export default {
  data: () => ({
    estadisticasHoy: {
      conteo: 0,
      recaudado: 0,
    },
    estadisticasMes: {
      conteo: 0,
      recaudado: 0,
    },
  }),
  async mounted() {
    await this.obtenerEstadisticasHoy();
    await this.obtenerEstadisticasMes();
    await this.obtenerEstadisticasMesActualGrafica();
    await this.obtenerEstadisticasAnioActualGrafica();
  },
  methods: {
    async obtenerEstadisticasHoy() {
      const hoy = new Date();
      const hoyInicioDia = Utiles.formatearFechaAInicioDeDia(hoy, "T");
      const hoyFinDia = Utiles.formatearFechaAFinDeDia(hoy, "T");
      this.estadisticasHoy.conteo = await EscritorioService.conteoVehiculos(
        hoyInicioDia,
        hoyFinDia
      );
      this.estadisticasHoy.recaudado = await EscritorioService.recaudadoVehiculos(
        hoyInicioDia,
        hoyFinDia
      );
    },
    async obtenerEstadisticasMes() {
      const hoy = new Date();
      const fechaInicioMes = new Date(hoy.getFullYear(), hoy.getMonth(), 1);
      const fechaFinMes = new Date(hoy.getFullYear(), hoy.getMonth() + 1, 0);
      const inicio = Utiles.formatearFechaAInicioDeDia(fechaInicioMes, "T");
      const fin = Utiles.formatearFechaAInicioDeDia(fechaFinMes, "T");
      this.estadisticasMes.conteo = await EscritorioService.conteoVehiculos(
        inicio,
        fin
      );
      this.estadisticasMes.recaudado = await EscritorioService.recaudadoVehiculos(
        inicio,
        fin
      );
    },
    async obtenerEstadisticasMesActualGrafica() {
      const hoy = new Date();
      const fechaInicioMes = new Date(hoy.getFullYear(), hoy.getMonth(), 1);
      const fechaFinMes = new Date(hoy.getFullYear(), hoy.getMonth() + 1, 0);
      const inicio = Utiles.formatearFechaAInicioDeDia(fechaInicioMes, "T");
      const fin = Utiles.formatearFechaAInicioDeDia(fechaFinMes, "T");
      const r = await EscritorioService.recaudadoVehiculosAgrupado(inicio, fin);
      console.log({ r });
    },
    async obtenerEstadisticasAnioActualGrafica() {
      const hoy = new Date();
      const fechaInicioAnio = new Date(hoy.getFullYear(), 0, 1);
      const fechaFinAnio = new Date(hoy.getFullYear(), 11, 31);
      const inicio = Utiles.formatearFechaAInicioDeDia(fechaInicioAnio, "T");
      const fin = Utiles.formatearFechaAInicioDeDia(fechaFinAnio, "T");
      const anioActual = await EscritorioService.recaudadoVehiculosAgrupadoPorMes(
        inicio,
        fin
      );
      console.log({ anioActual });
    },
  },
};
</script>