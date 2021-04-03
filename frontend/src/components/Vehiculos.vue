<template>
  <div class="columns">
    <div class="column">
      <b-field label="Placas">
        <b-input placeholder="Placas"></b-input>
      </b-field>
      <b-field label="Descripción del vehículo">
        <b-input placeholder="Descripción del vehículo"></b-input>
      </b-field>

      <b-field>
        <b-switch v-model="usarHoraActual"
          >Usar hora actual ({{ horaActual }})</b-switch
        >
      </b-field>

      <b-field label="Select datetime" v-show="!usarHoraActual">
        {{ fecha instanceof Date }}
        <b-datetimepicker
          v-model="fecha"
          rounded
          placeholder="Click to select..."
          icon="calendar-today"
          locale="es-MX"
          :timepicker="{ enableSeconds: true, hourFormat: '12' }"
          horizontal-time-picker
        >
        </b-datetimepicker>
      </b-field>
    </div>
  </div>
</template>
<script>
export default {
  data: () => ({
    fecha: null,
    horaActual: "Soy la hora actual",
    usarHoraActual: true,
  }),
};
</script>