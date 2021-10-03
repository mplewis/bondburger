<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { pick } from '../logic/films'
import { Question, buildQuestion } from '../logic/questions';// TODO: Quiz and Question need to be totally reworked and renamed

const choices = 4
const sentencePadding = 1;

var question = ref<Question>(null)

onMounted(async () => {
  const { q } = buildQuestion(pick(4))
  options.
})
</script>

<template>
  <Question :question="question" :options="options" />
</template>
