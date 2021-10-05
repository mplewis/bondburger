<script setup lang="ts">
import "bulma/bulma.sass";
import { onMounted } from "vue";
import { store } from "./store";
import Navbar from "./components/Navbar.vue"
import { films } from "./logic/films";

onMounted(() => {
  films() // preload
  store.actions.newQuestion()
})
</script>

<template>
  <div class="container p-3">
    <Navbar class="mb-3" />
    <router-view />
  </div>
</template>
