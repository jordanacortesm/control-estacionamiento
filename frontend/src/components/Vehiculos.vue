<template>
  <div class="columns">
    <div class="column">
      <router-link
        :to="{ name: 'RegistrarVehiculo' }"
        class="button is-success"
      >
        Registrar nuevo
      </router-link>
    </div>
  </div>
</template>
<script>
export default {
  data: () => ({}),
  mounted() {},
  methods: {},
};
</script>