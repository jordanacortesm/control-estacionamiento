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
              <div v-show="cargandoEstadisticasHoy">
                <b-skeleton width="60%" animated height="50px"></b-skeleton>
                <div class="mt-5"></div>
                <b-skeleton width="50%" animated height="45px"></b-skeleton>
              </div>
              <div v-show="!cargandoEstadisticasHoy">
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
      </div>
      <div class="column">
        <div class="card">
          <header class="card-header">
            <p class="card-header-title">Este mes</p>
          </header>
          <div class="card-content">
            <div class="content">
              <div v-show="cargandoEstadisticasMes">
                <b-skeleton width="60%" animated height="50px"></b-skeleton>
                <div class="mt-5"></div>
                <b-skeleton width="50%" animated height="45px"></b-skeleton>
              </div>
              <div v-show="!cargandoEstadisticasMes">
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
    </div>
    <div class="columns">
      <div class="column">
        <div class="card">
          <header class="card-header">
            <p class="card-header-title">Recaudado este mes</p>
          </header>
          <div class="card-content">
            <div v-show="cargandoGraficaMes">
              <b-skeleton
                width="30%"
                animated
                position="is-centered"
              ></b-skeleton>
              <b-skeleton height="180px" animated></b-skeleton>
              <b-skeleton animated position="is-centered"></b-skeleton>
            </div>
            <canvas
              v-show="!cargandoGraficaMes"
              id="graficaMes"
              width="400"
              height="200"
            ></canvas>
          </div>
        </div>
      </div>
      <div class="column">
        <div class="card">
          <header class="card-header">
            <p class="card-header-title">Recaudado este año</p>
          </header>
          <div class="card-content">
            <div v-show="cargandoGraficaAnio">
              <b-skeleton
                width="30%"
                animated
                position="is-centered"
              ></b-skeleton>
              <b-skeleton height="180px" animated></b-skeleton>
              <b-skeleton animated position="is-centered"></b-skeleton>
            </div>
            <canvas
              v-show="!cargandoGraficaAnio"
              id="graficaAnio"
              width="400"
              height="200"
            ></canvas>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import EscritorioService from "../services/EscritorioService";
import Utiles from "../services/Utiles";
import Chart from "chart.js/auto";
export default {
  data: () => ({
    cargandoEstadisticasHoy: true,
    cargandoEstadisticasMes: true,
    cargandoGraficaMes: true,
    cargandoGraficaAnio: true,
    estadisticasHoy: {
      conteo: 0,
      recaudado: 0,
    },
    estadisticasMes: {
      conteo: 0,
      recaudado: 0,
    },
    formateadorDinero: null,
  }),
  async mounted() {
    this.formateadorDinero = this.$options.filters.dinero;
    this.obtenerEstadisticasHoy();
    this.obtenerEstadisticasMes();
    this.obtenerEstadisticasMesActualGrafica();
    this.obtenerEstadisticasAnioActualGrafica();
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
      this.cargandoEstadisticasHoy = false;
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
      this.cargandoEstadisticasMes = false;
    },
    async obtenerEstadisticasMesActualGrafica() {
      const hoy = new Date();
      const fechaInicioMes = new Date(hoy.getFullYear(), hoy.getMonth(), 1);
      const fechaFinMes = new Date(hoy.getFullYear(), hoy.getMonth() + 1, 0);
      const inicio = Utiles.formatearFechaAInicioDeDia(fechaInicioMes, "T");
      const fin = Utiles.formatearFechaAInicioDeDia(fechaFinMes, "T");
      const datos = await EscritorioService.recaudadoVehiculosAgrupado(
        inicio,
        fin
      );
      const formateador = this.formateadorDinero;
      new Chart(document.querySelector("#graficaMes"), {
        type: "line",
        data: {
          labels: datos.map((dato) => dato.dia),
          datasets: [
            {
              label: "Recaudado",
              data: datos.map((dato) => dato.recaudado),
              backgroundColor: [
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
              ],
              borderColor: [
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
              ],
              borderWidth: 1,
            },
          ],
        },
        options: {
          plugins: {
            tooltip: {
              callbacks: {
                title(contextos) {
                  const contexto = contextos[0];
                  return "Día " + contexto.label;
                },
                label(contexto) {
                  const dineroSinFormato = contexto.parsed.y;
                  return formateador(dineroSinFormato);
                },
              },
            },
          },
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
      this.cargandoGraficaMes = false;
    },
    async obtenerEstadisticasAnioActualGrafica() {
      const hoy = new Date();
      const fechaInicioAnio = new Date(hoy.getFullYear(), 0, 1);
      const fechaFinAnio = new Date(hoy.getFullYear(), 11, 31);
      const inicio = Utiles.formatearFechaAInicioDeDia(fechaInicioAnio, "T");
      const fin = Utiles.formatearFechaAInicioDeDia(fechaFinAnio, "T");
      const datos = await EscritorioService.recaudadoVehiculosAgrupadoPorMes(
        inicio,
        fin
      );
      const formateador = this.formateadorDinero;
      new Chart(document.querySelector("#graficaAnio"), {
        type: "bar",
        data: {
          labels: datos.map((dato) => dato.mes),
          datasets: [
            {
              label: "Recaudado",
              data: datos.map((dato) => dato.recaudado),
              backgroundColor: [
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
                "rgba(54, 162, 235, 0.2)",
                "rgba(255, 206, 86, 0.2)",
                "rgba(75, 192, 192, 0.2)",
                "rgba(153, 102, 255, 0.2)",
                "rgba(255, 159, 64, 0.2)",
                "rgba(255, 99, 132, 0.2)",
              ],
              borderColor: [
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
                "rgba(54, 162, 235, 1)",
                "rgba(255, 206, 86, 1)",
                "rgba(75, 192, 192, 1)",
                "rgba(153, 102, 255, 1)",
                "rgba(255, 159, 64, 1)",
                "rgba(255, 99, 132, 1)",
              ],
              borderWidth: 1,
            },
          ],
        },
        options: {
          plugins: {
            tooltip: {
              callbacks: {
                title(contextos) {
                  const contexto = contextos[0];
                  // Que dios me perdone por esto
                  return "Enero Febrero Marzo Abril Mayo Junio Julio Agosto Septiembre Octubre Noviembre Diciembre".split(
                    " "
                  )[parseInt(contexto.label) - 1];
                },
                label(contexto) {
                  const dineroSinFormato = contexto.parsed.y;
                  return formateador(dineroSinFormato);
                },
              },
            },
          },
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
      this.cargandoGraficaAnio = false;
    },
  },
};
</script>