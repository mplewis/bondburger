<script setup lang="ts">
document.addEventListener('DOMContentLoaded', () => {
  const navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);
  navbarBurgers.forEach(el => {
    el.addEventListener('click', () => {
      const target = document.getElementById(el.dataset.target);
      if (!target) return;
      el.classList.toggle('is-active');
      target.classList.toggle('is-active');
    });
  });
});
</script>

<template>
  <nav class="navbar" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <router-link class="navbar-item is-tab has-text-weight-bold" to="/">Bondburger</router-link>
      <a
        role="button"
        class="navbar-burger"
        aria-label="menu"
        aria-expanded="false"
        data-target="navbar-menu"
      >
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>

    <div id="navbar-menu" class="navbar-menu">
      <div class="navbar-start">
        <router-link class="navbar-item is-tab" to="/quiz">Quiz</router-link>
        <router-link class="navbar-item is-tab" to="/results">Results</router-link>
        <router-link class="navbar-item is-tab" to="/about">About</router-link>
      </div>
    </div>
  </nav>
</template>
