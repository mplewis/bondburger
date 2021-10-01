import { createRouter, createWebHashHistory } from "vue-router";

import About from "./components/About.vue";
import Home from "./components/Home.vue";
import Quiz from "./components/Quiz.vue";
import Results from "./components/Results.vue";

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: "/", component: Home },
    { path: "/about", component: About },
    { path: "/quiz", component: Quiz },
    { path: "/results", component: Results },
  ],
});
