<template>
    <div>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Age</th>
            <th>Education</th>
            <th>Income</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(row, index) in data" :key="index">
            <td>{{ row.name }}</td>
            <td>{{ row.age }}</td>
            <td>{{ row.education }}</td>
            <td>{{ row.income }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script>
  import api from '../api';
  
  export default {
    data() {
      return {
        data: [],
        loading: true,
      };
    },
    async created() {
      try {
        const response = await api.get('/api/census');
        this.data = response.data;
        this.loading = false;
      } catch (error) {
        console.error(error);
      }
    },
  };
  </script>