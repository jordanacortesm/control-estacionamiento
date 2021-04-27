import Vue from 'vue'
import Router from 'vue-router'
import Vehiculos from '@/components/Vehiculos'
import RegistrarVehiculo from '@/components/RegistrarVehiculo'

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
    ]
});