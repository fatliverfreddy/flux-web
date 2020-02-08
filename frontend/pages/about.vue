<template>
  <dashboard-layout>
    <template slot="right-side">
      <span>About</span>
    </template>
  </dashboard-layout>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import DashboardLayout from "~/layouts/DashboardLayout.vue";

@Component({
  components: {
    DashboardLayout,
  }
})
export default class extends Vue {}
</script>

<style lang="scss"></style>
