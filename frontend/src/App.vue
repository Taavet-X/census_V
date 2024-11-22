<template>
    <div id="app">
      <Navbar />
      <main>
        <router-view />
      </main>
    </div>
  </template>
  
  <script>
  import Navbar from './components/Navbar.vue';
  
  export default {
    name: 'App',
    components: {
      Navbar,
    },
  };
  </script>
  
  <style>
  /* App-wide styles */
  #app {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
  }
  
  main {
    padding: 1rem;
    background-color: #f5f5f5;
    min-height: 100vh;
  }
  
  a {
    text-decoration: none;
    color: inherit;
  }
  
  a:hover {
    text-decoration: underline;
  }
  </style>