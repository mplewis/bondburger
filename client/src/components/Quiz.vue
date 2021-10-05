<script setup lang="ts">
import { useStore } from '../store';

const store = useStore();
</script>

<template>
  <p v-if="!store.getters.question">Loading...</p>
  <div v-else>
    <p class="question mb-3">{{ store.getters.question.text }}</p>
    <button v-for="option in store.getters.question.options" class="button mb-3">
      {{ option.title }}
      ({{ option.year }}, {{ option.actor }})
    </button>
  </div>
</template>

<style scoped lang="scss">
button {
  display: block;
  background-color: #e67e22;
  color: white !important;
  border: none;

  &:hover {
    background-color: #d35400;
  }
}
</style>
