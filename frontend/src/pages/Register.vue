<template>
    <div>
      <form @submit.prevent="register">
        <h2>Registro</h2>
        <input type="text" v-model="username" placeholder="Nombre de usuario" required />
        <input type="email" v-model="email" placeholder="Correo electrónico" required />
        <input type="password" v-model="password" placeholder="Contraseña" required />
        <button type="submit">Registrar</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        username: '',
        email: '',
        password: ''
      };
    },
    methods: {
      async register() {
        try {
          const response = await this.$axios.post('/api/register', {
            username: this.username,
            email: this.email,
            password: this.password
          });
          console.log(response.data);
          this.$router.push('/login');
        } catch (error) {
          console.error(error);
        }
      }
    }
  };
  </script>