<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { pick, Film, Slug, FilmsBySlug } from '../data/films'
import { random, randInt } from '../random'
import Question from './Question.vue'

const choices = 4
const sentencePadding = 1;

var options = ref<FilmsBySlug>({})
var question = ref<string>("")
var answer = ref<Slug>("")

function buildQuestion(film: Film, padding: number): string {
  const min = padding
  const max = film.plot.length - 1 - padding
  const base = randInt(min, max)
  const q = film.plot.slice(base - padding, base + padding + 1)
  return q.join(' ')
}

onMounted(async () => {
  const ch = await pick(choices)
  options.value = ch
  const an = random(Object.keys(ch))
  const anf = ch[an]
  answer.value = an
  question.value = buildQuestion(anf, sentencePadding)
})
</script>

<template>
  <Question :question="question" :options="options" />
</template>
