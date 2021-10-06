<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from '../store';
import AnswerButton from './AnswerButton.vue';

const store = useStore();

const qa = computed(() => store.getters.qa)
const score = computed(() => store.getters.score)

function nextQuestion() {
  store.actions.newQuestion()
}
</script>

<template>
  <p v-if="!qa">Loading...</p>
  <div v-else>
    <p class="mb-3">
      <span class="score">Your score: {{ score.correct }} of {{ score.total }}</span>
      <a class="ml-2" @click="store.actions.clearScore">Clear</a>
    </p>
    <p class="question mb-3">{{ qa.question.text }}</p>
    <AnswerButton v-for="option in qa.question.options" :film="option" />
    <button v-if="qa.answer" class="option" @click="nextQuestion">Continue &raquo;</button>
  </div>
</template>

<style scoped lang="scss">
.score {
  font-weight: 700;
}
</style>

<style lang="scss">
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
