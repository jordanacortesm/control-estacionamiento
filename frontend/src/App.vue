<template>
  <div class="container">
    <div class="columns">
      <div class="column">
        <section>
          <h1 class="is-size-1">Probando Buefy</h1>
          <button class="button is-medium" @click="mostrarToast">
            Mostrar un toast
          </button>
        </section>
        <Vehiculos></Vehiculos>
      </div>
    </div>
  </div>
</template>

<script>
import Vehiculos from "./components/Vehiculos";
export default {
  name: "app",
  components: { Vehiculos },
  methods: {
    mostrarToast() {
      this.$buefy.toast.open("¡Hola, Buefy! Parzibyte.me");
    },
  },
};
</script>