<template>
    <form @submit.prevent="handleSubmit">
      <input
        type="email"
        v-model="formData.email"
        placeholder="Email"
        required
      />
      <input
        type="password"
        v-model="formData.password"
        placeholder="Password"
        required
      />
      <button type="submit">{{ isLogin ? 'Login' : 'Register' }}</button>
    </form>
  </template>
  
  <script>
  import api from '../api';
  
  export default {
    props: {
      isLogin: Boolean,
    },
    data() {
      return {
        formData: {
          email: '',
          password: '',
        },
      };
    },
    methods: {
      async handleSubmit() {
        const endpoint = this.isLogin ? '/api/login' : '/api/register';
        try {
          await api.post(endpoint, this.formData);
          this.$router.push('/');
        } catch (error) {
          console.error(error);
        }
      },
    },
  };
  </script>