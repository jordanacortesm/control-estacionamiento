import Vue from 'vue'
import App from './App.vue'
import '@mdi/font/css/materialdesignicons.css'
import router from "./router"

import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
Vue.use(Buefy);
Vue.config.productionTip = false
// Filtros
const formateador = new Intl.DateTimeFormat('es-MX', { dateStyle: 'medium', timeStyle: 'medium' });
Vue.filter("formatearFecha", fecha => formateador.format(new Date(fecha)));

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
