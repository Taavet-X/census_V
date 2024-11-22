<template>
    <div>
      <form @submit.prevent="login">
        <h2>Iniciar sesión</h2>
        <input 
          type="email" 
          v-model="email" 
          placeholder="Correo electrónico" 
          required 
          autofocus 
        />
        <input 
          type="password" 
          v-model="password" 
          placeholder="Contraseña" 
          required 
        />
        <button type="submit">Iniciar sesión</button>
        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
      </form>
      <p>No tienes cuenta? <router-link to="/register">Regístrate aquí</router-link></p>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        email: '',
        password: '',
        errorMessage: ''
      };
    },
    methods: {
      async login() {
        try {
          const response = await this.$axios.post('/api/login', {
            email: this.email,
            password: this.password
          });
          
          // Guardar el token JWT en localStorage
          localStorage.setItem('authToken', response.data.token);
          
          // Redirigir a la página principal o página de configuración, etc.
          this.$router.push('/dashboard'); // Ajusta la ruta a la que desees redirigir después del login
        } catch (error) {
          // Mostrar mensaje de error si las credenciales son incorrectas
          this.errorMessage = 'Correo electrónico o contraseña incorrectos';
          console.error(error);
        }
      }
    }
  };
  </script>
  
  <style scoped>
  form {
    display: flex;
    flex-direction: column;
    max-width: 300px;
    margin: 0 auto;
  }
  
  input {
    padding: 10px;
    margin: 10px 0;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  button {
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #0056b3;
  }
  
  .error {
    color: red;
    font-size: 14px;
    margin-top: 10px;
  }
  
  p {
    text-align: center;
  }
  
  a {
    color: #007bff;
    text-decoration: none;
  }
  
  a:hover {
    text-decoration: underline;
  }
  </style>