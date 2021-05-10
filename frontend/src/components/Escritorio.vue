<template>
  <div>
    <div class="columns">
      <div class="column">
        <p>Escritorio</p>
        {{ estadisticasHoy }}
        <strong>Mes:</strong>
        {{ estadisticasMes }}
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