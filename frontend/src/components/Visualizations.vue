<template>
    <div class="visualizations">
      <h2>Visualización de Datos Demográficos</h2>
      <div class="charts">
        <div id="age-distribution" class="chart"></div>
        <div id="income-distribution" class="chart"></div>
      </div>
    </div>
  </template>
  
  <script>
  import { onMounted } from 'vue';
  import { createChart } from 'chart.js';
  
  export default {
    name: 'Visualizations',
    setup() {
      onMounted(() => {
        renderAgeDistribution();
        renderIncomeDistribution();
      });
  
      function renderAgeDistribution() {
        const ctx = document.getElementById('age-distribution').getContext('2d');
        new createChart(ctx, {
          type: 'bar',
          data: {
            labels: ['18-25', '26-35', '36-45', '46-60', '60+'],
            datasets: [
              {
                label: 'Distribución de Edad',
                data: [20, 30, 25, 15, 10], // Example data
                backgroundColor: '#42a5f5',
              },
            ],
          },
        });
      }
  
      function renderIncomeDistribution() {
        const ctx = document.getElementById('income-distribution').getContext('2d');
        new createChart(ctx, {
          type: 'pie',
          data: {
            labels: ['<50K', '>=50K'],
            datasets: [
              {
                label: 'Distribución de Ingresos',
                data: [60, 40], // Example data
                backgroundColor: ['#ff7043', '#42f57b'],
              },
            ],
          },
        });
      }
    },
  };
  </script>
  
  <style scoped>
  .visualizations {
    padding: 1rem;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
  
  h2 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
  }
  
  .charts {
    display: flex;
    gap: 2rem;
  }
  
  .chart {
    width: 45%;
    height: 300px;
  }
  </style>