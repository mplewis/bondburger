import { createVuexModule } from "vuex-typed-modules";
import { pick } from "./logic/films";
import { buildQuestion, Question } from "./logic/questions";
import { choices } from "./config";
import { Slug } from "./types";

export type QuestionAnswer = {
  question: Question;
  answer?: Slug;
};

type State = {
  qas: QuestionAnswer[];
};

const initialState: State = {
  qas: [],
};

export const [store, useStore] = createVuexModule({
  name: "store",
  state: initialState,
  getters: {
    qa(state): QuestionAnswer | undefined {
      if (state.qas.length === 0) return;
      const i = state.qas.length - 1;
      return state.qas[i];
    },
    score(state): { correct: number; total: number } {
      const correct = state.qas.reduce(
        (acc, qa) => (qa.answer === qa.question.answer.slug ? acc + 1 : acc),
        0
      );
      const total = state.qas.length;
      return { correct, total };
    },
  },
  mutations: {
    setCurrentQuestion(state, question: Question) {
      state.qas.push({ question });
    },
    setCurrentAnswer(state, answer: Slug) {
      const i = state.qas.length - 1;
      state.qas[i].answer = answer;
    },
  },
  actions: {
    async newQuestion(_) {
      const q = buildQuestion(await pick(choices));
      store.mutations.setCurrentQuestion(q);
    },
    submitAnswer(_, answer: Slug) {
      store.mutations.setCurrentAnswer(answer);
    },
  },
});
