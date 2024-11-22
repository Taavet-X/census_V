<template>
    <div class="home">
      <h1>Bienvenido a la Visualización de Datos Demográficos</h1>
      
      <div class="filters">
        <FilterPanel />
      </div>
  
      <div class="statistics">
        <Summary />
      </div>
  
      <div class="visualizations">
        <Visualizations />
      </div>
  
      <div class="export">
        <button @click="exportData">Exportar Datos</button>
      </div>
    </div>
  </template>
  
  <script>
  import FilterPanel from './FilterPanel.vue';
  import Summary from './Summary.vue';
  import Visualizations from './Visualizations.vue';
  
  export default {
    components: {
      FilterPanel,
      Summary,
      Visualizations
    },
    methods: {
      async exportData() {
        try {
          const token = localStorage.getItem('authToken');
          if (!token) {
            alert('Por favor, inicie sesión para exportar datos');
            return;
          }
  
          const response = await this.$axios.get('/api/export', {
            headers: {
              Authorization: `Bearer ${token}`
            },
            responseType: 'blob'
          });
  
          const url = window.URL.createObjectURL(new Blob([response.data]));
          const link = document.createElement('a');
          link.href = url;
          link.setAttribute('download', 'datos_censales.csv');
          document.body.appendChild(link);
          link.click();
        } catch (error) {
          console.error('Error al exportar los datos:', error);
          alert('Hubo un error al intentar exportar los datos');
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .home {
    text-align: center;
    padding: 20px;
  }
  
  h1 {
    font-size: 2rem;
    margin-bottom: 20px;
  }
  
  .filters {
    margin-bottom: 20px;
  }
  
  .statistics {
    margin-top: 20px;
  }
  
  .export {
    margin-top: 20px;
  }
  
  button {
    padding: 10px 20px;
    background-color: #28a745;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #218838;
  }
  </style>