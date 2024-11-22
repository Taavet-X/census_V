<template>
    <div class="summary">
      <h2>Resumen de Datos</h2>
      <div v-if="summaryData">
        <p><strong>Edad Promedio:</strong> {{ summaryData.averageAge }} años</p>
        <p><strong>Ingreso Promedio:</strong> ${{ summaryData.averageIncome }}</p>
        <p><strong>Horas de Trabajo Promedio:</strong> {{ summaryData.averageHours }} horas</p>
        <p><strong>Distribución de Ingresos:</strong></p>
        <ul>
          <li v-for="(incomeRange, index) in summaryData.incomeDistribution" :key="index">
            {{ incomeRange.range }}: {{ incomeRange.count }} personas
          </li>
        </ul>
      </div>
      <div v-else>
        <p>Cargando estadísticas...</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Summary',
    data() {
      return {
        summaryData: null,
      };
    },
    created() {
      this.fetchSummaryData();
    },
    methods: {
      async fetchSummaryData() {
        try {
          const response = await fetch('/api/stats');
          if (response.ok) {
            this.summaryData = await response.json();
          } else {
            console.error('Error al obtener los datos');
          }
        } catch (error) {
          console.error('Error en la solicitud:', error);
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .summary {
    padding: 1rem;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
  
  h2 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
  }
  
  ul {
    list-style: none;
    padding: 0;
  }
  
  li {
    margin: 0.5rem 0;
  }
  </style>