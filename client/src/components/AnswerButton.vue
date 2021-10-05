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
  </div>
</template>

<style scoped lang="scss">
$button-color: #e67e22;
$button-color-hover: lighten($button-color, 10%);
$button-color-active: #d35400;

$button-color-correct: #2ecc71;
$button-color-incorrect: #e74c3c;

button {
  font-size: 1em;
  font-weight: 700;

  color: rgba(0, 0, 0, 0.8);
  background: none;

  padding: 0.75em;
  margin-bottom: 0.75rem;

  border: 3px solid rgba(0, 0, 0, 0.1);
  border-radius: 6px;

  &.option {
    background-color: $button-color;
    border-color: $button-color;
    color: white;

    &:hover {
      background-color: $button-color-hover;
      border-color: $button-color-hover;
    }

    &:active {
      background-color: $button-color-active;
      border-color: $button-color-active;
    }
  }

  &.correct {
    border-color: $button-color-correct;
  }

  &.incorrect {
    border-color: $button-color-incorrect;
  }
}
</style>
