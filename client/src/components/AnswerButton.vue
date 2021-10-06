<script setup lang="ts">
import { defineProps, computed } from 'vue'
import { useStore } from '../store';
import { Film, Slug } from '../types';

const props = defineProps<{ film: Film }>()

const store = useStore()

function submit(slug: Slug) {
  store.actions.submitAnswer(slug)
}

const qa = computed(() => store.getters.qa)
const expected = computed(() => qa.value?.question.answer.slug)
const actual = computed(() => qa.value?.answer)
const disabled = computed(() => !!actual.value)
const klass = computed(() => {
  if (!actual.value) return 'option'
  if (props.film.slug === expected.value) return 'correct'
  if (props.film.slug === actual.value) return 'incorrect'
})
</script>

<template>
  <div>
    <button :class="klass" @click="submit(film.slug)" :disabled="disabled">
      {{ film.title }}
      ({{ film.year }}, {{ film.actor }})
    </button>
    <span v-if="props.film.slug === actual">
      <span v-if="klass === 'correct'" class="ml-3 annotation correct">✅ Correct!</span>
      <span v-else-if="klass === 'incorrect'" class="ml-3 annotation incorrect">❌ Incorrect</span>
    </span>
  </div>
</template>

<style scoped lang="scss">
.annotation {
  font-weight: 700;
  &.correct {
    color: #27ae60;
  }
  &.incorrect {
    color: #c0392b;
  }
}
</style>
