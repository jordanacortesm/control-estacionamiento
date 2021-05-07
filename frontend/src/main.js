import Vue from 'vue'
import App from './App.vue'
import '@mdi/font/css/materialdesignicons.css'
import router from "./router"
import Utiles from "./services/Utiles"

import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
Vue.use(Buefy);
Vue.config.productionTip = false
// Filtros
Vue.filter("formatearFecha", fechaComoCadena => Utiles.formatearFechaYHoraSegunLocale(new Date(fechaComoCadena)));
Vue.filter("minutosAHorasYMinutos", Utiles.minutosAHorasYMinutos);
Vue.filter("dinero", Utiles.formatearDinero);

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
