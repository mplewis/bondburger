<script setup lang="ts">
import "bulma/bulma.sass";
import Navbar from "./components/Navbar.vue"
import { useStore } from './store'

const store = useStore()

function click() {
  store.actions.addCountAsync(1)
}
</script>

<template>
  <div class="container p-3">
    <Navbar class="mb-3" />
    <p>{{ store.state.count }}</p>
    <button @click="click">+</button>
    <router-view />
  </div>
</template>
