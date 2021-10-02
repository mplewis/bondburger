<script setup lang="ts">
import "bulma/bulma.sass";
import Navbar from "./components/Navbar.vue"
import { films } from './data/films';

document.addEventListener('DOMContentLoaded', async () => {
  const f = await films();
  console.log(f)
})
</script>

<template>
  <div class="container p-3">
    <Navbar class="mb-3" />
    <router-view />
  </div>
</template>
