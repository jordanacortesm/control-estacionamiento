<template>
  <div>
    <b-navbar type="is-warning">
      <template #brand>
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img :src="logo" />
        </b-navbar-item>
      </template>
      <template #start>
        <b-navbar-item tag="router-link" :to="{ name: 'Vehiculos' }"
          >Vehículos</b-navbar-item
        >
        <b-navbar-item tag="router-link" :to="{ name: 'Ajustes' }"
          >Ajustar costos</b-navbar-item
        >
        <b-navbar-item tag="router-link" :to="{ name: 'ReporteVehiculos' }"
          >Reporte</b-navbar-item
        >
        <b-navbar-item tag="router-link" :to="{ name: 'Escritorio' }"
          >Escritorio</b-navbar-item
        >
      </template>

      <template #end>
        <b-navbar-item tag="div">
          <div class="buttons">
            <a class="button is-primary">
              <b-icon icon="star"></b-icon>
              <strong>Ayuda y soporte</strong>
            </a>
          </div>
        </b-navbar-item>
      </template>
    </b-navbar>
    <section class="section">
      <router-view></router-view>
    </section>
  </div>
</template>

<script>
import Logo from "./assets/parzibyte.png";
export default {
  name: "app",
  data: () => ({
    logo: Logo,
  }),
  methods: {},
};
</script>