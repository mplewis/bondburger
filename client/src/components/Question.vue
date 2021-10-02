<script setup lang="ts">
import { FilmsBySlug } from '../data/films'

defineProps<{
  question: string;
  options: FilmsBySlug;
}>()
</script>

<template>
  <p class="question mb-3">{{ question }}</p>
  <button v-for="option in options" class="button mb-3">
    {{ option.title }}
    ({{ option.year }}, {{ option.actor }})
  </button>
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
