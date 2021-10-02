<script setup lang="ts">
import Question from './Question.vue'
import { pick, Film, Slug, FilmsBySlug } from '../data/films'
import { ref, onMounted } from 'vue';

const choices = 4
const sentencePadding = 1;

var options = ref<FilmsBySlug>({})
var question = ref<string>("")
var answer = ref<Slug>("")

function random<T>(items: T[]): T {
  return items[Math.floor(Math.random() * items.length)]
}

function randInt(min: number, max: number) {
  return Math.floor(Math.random() * (max - min + 1)) + min
}

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
