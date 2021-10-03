import { createApp } from "vue";

import App from "./App.vue";
import router from "./router";

import { Database } from "vuex-typed-modules";
import { store as myStore } from "./store";
import { Store } from "vuex";

const database = new Database({ logger: true });
const store = new Store({ plugins: [database.deploy([myStore])] });

createApp(App).use(router).use(store).mount("#app");
