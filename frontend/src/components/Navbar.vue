<template>
    <nav class="navbar">
      <router-link to="/">Home</router-link>
      <router-link to="/login">Login</router-link>
      <router-link to="/register">Register</router-link>
    </nav>
  </template>
  
  <style scoped>
  .navbar {
    display: flex;
    gap: 1rem;
    background-color: #f8f9fa;
    padding: 1rem;
  }
  </style>