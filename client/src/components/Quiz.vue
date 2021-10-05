<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from '../store';
import AnswerButton from './AnswerButton.vue';

const store = useStore();

const qa = computed(() => store.getters.qa)
</script>

<template>
  <p v-if="!qa">Loading...</p>
  <div v-else>
    <p class="score mb-3">Your score: 0 of 0</p>
    <p class="question mb-3">{{ qa.question.text }}</p>
    <AnswerButton v-for="option in qa.question.options" :film="option" />
  </div>
</template>

<style scoped lang="scss">
.score {
  font-weight: 700;
}
</style>
