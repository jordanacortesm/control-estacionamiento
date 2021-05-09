import Vue from 'vue'
import Router from 'vue-router'
import Vehiculos from '@/components/Vehiculos'
import RegistrarVehiculo from '@/components/RegistrarVehiculo'
import Ajustes from '@/components/Ajustes'
import ReporteVehiculos from '@/components/ReporteVehiculos'

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Vehiculos',
            component: Vehiculos
        },
        {
            path: '/agregar-vehiculo',
            name: 'RegistrarVehiculo',
            component: RegistrarVehiculo,
        },
        {
            path: '/ajustes',
            name: 'Ajustes',
            component: Ajustes,
        },
        {
            path: '/reporte-vehiculos',
            name: 'ReporteVehiculos',
            component: ReporteVehiculos,
        },
    ]
});