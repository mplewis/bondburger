import { createVuexModule } from "vuex-typed-modules";

export const [store, useStore] = createVuexModule({
  name: "store",
  state: {
    count: 1,
  },
  mutations: {
    addCount(state, number: number) {
      state.count += number;
    },
  },
  actions: {
    async addCountAsync(_, count: number): Promise<void> {
      store.mutations.addCount(count);
    },
  },
});
