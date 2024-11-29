
<script setup lang="ts">
import { onBeforeMount,  } from 'vue';
import {useCensusRecords} from '../../store/record';
import SaveFilters from '../SaveFilters.vue'
import ChartOccupations from '../ChartOccupations.vue';
import ChartEducation from '../ChartEducation.vue';


const records = useCensusRecords()

  onBeforeMount(() => {
   records.loadStateFromSession();
   records.fetchData(records.page, records.sortField, records.sortOrder, records.filters);
   });


  function handleSort(field: string) {
     if (records.sortField === field) {
      records.sortOrder = records.sortOrder === 'asc' ? 'desc' : 'asc';
    } else {
      records.sortField = field;
      records.sortOrder = 'asc';
    }
  
    records.fetchData(records.page, records.sortField, records.sortOrder, records.filters);
  
    const urlParams = new URLSearchParams(window.location.search);
    urlParams.set('sort_by', records.sortField || '');
    urlParams.set('order', records.sortOrder);
    window.history.pushState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
  }
  
  
  const ageOptions = [
  17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
  37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55,
  56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
  76, 80, 90,
];

const educationOptions = [
  "Bachelors", "HS-grad", "11th", "Masters", "9th", "Some-college",
  "Assoc-acdm", "Assoc-voc", "7th-8th", "Doctorate", "Prof-school",
  "5th-6th", "10th"
];

const maritalStatusOptions = [
  "Never-married", 
  "Married-civ-spouse", 
  "Divorced", 
  "Married-spouse-absent", 
  "Separated", 
  "Married-AF-spouse"
];

const occupationOptions = [
  "Adm-clerical", 
  "Exec-managerial", 
  "Handlers-cleaners", 
  "Prof-specialty", 
  "Other-service", 
  "Sales", 
  "Craft-repair", 
  "Transport-moving", 
  "Farming-fishing", 
  "Machine-op-inspct", 
  "Tech-support", 
  "Protective-serv"
];


  function handleFilterChange(key: string, event: Event) {
    const select = event.target as HTMLSelectElement;
    const value = select.value;
    const mappedKey = key === 'MaritalStatus' ? 'marital_status' : key.toLowerCase();

    const mappedValue = value === "less-than-50K" ? "<=50K" : value;

    records.filters[key] = mappedValue;
  
    if (value === '') {
      delete records.filters[mappedKey];
    } else {
      records.filters[mappedKey.toLowerCase()] = value;
    }
  
    records.fetchData(records.page, records.sortField, records.sortOrder, records.filters);

    records.promediumAge
    const urlParams = new URLSearchParams(window.location.search);
    for (const [key, value] of Object.entries(records.filters)) {
      urlParams.set(key, value);
    }
    window.history.pushState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
 
  }

   function handlePageChange(page: number) {
  if (page >= 1 && page <= records.totalPages) {
    records.page = page;
    console.log("Cambio de página:", page);
    records.fetchData(records.page, records.sortField, records.sortOrder, records.filters);
    const urlParams = new URLSearchParams(window.location.search);
    urlParams.set('page', page.toString());
    window.history.pushState({}, '', `${window.location.pathname}?${urlParams.toString()}`);
  }
} 

</script>

<template>
  <div class="table-container">
    <h1>Census Data</h1>
    <div class="filters">
      <SaveFilters></SaveFilters>
      <div class="filter">
        <label for="age">Age</label>
        <select id="age" @change="handleFilterChange('Age', $event)">
          <option value="">All</option>
          <option v-for="age in ageOptions" :key="age" :value="age">
            {{ age }}
          </option>
        </select>
      </div>
      
      <div class="filter">
        <label for="education">Education</label>
        <select id="education" @change="handleFilterChange('Education', $event)">
          <option value="">All</option>
          <option v-for="option in educationOptions" :key="option" :value="option">
            {{ option }}
          </option>
        </select>
      </div>
      
      
      <div class="filter">
        <label for="marital-status">Marital Status</label>
        <select id="marital-status" @change="handleFilterChange('MaritalStatus', $event)">
          <option value="">All</option>
          <option v-for="option in maritalStatusOptions" :key="option" :value="option">
            {{ option }}
          </option>
        </select>
      </div>
      
      <div class="filter">
        <label for="occupation">Occupation</label>
        <select id="occupation" @change="handleFilterChange('Occupation', $event)">
          <option value="">All</option>
          <option v-for="option in occupationOptions" :key="option" :value="option">
            {{ option }}
          </option>
        </select>
      </div>
      
      <div class="filter">
        <label for="income">Income</label>
        <select id="income" @change="handleFilterChange('Income', $event)">
          <option value="">All</option>
          <option value="&lt;=50K">&lt;=50K</option>
          <option value=">50K">>50K</option>
        </select>
      </div>
   
   

      <div class="promediums">
        <div class="promediumAge">
        <h4>Promedium Age</h4>
        <p>{{records.promediumAge }}</p>
      </div>
      <div class="promediumHours">
        <h4>Promedium Hours</h4>
        <p>{{ records.promediumHours}}</p>
      </div>
      </div>
    
    </div>
    <div class="filters"></div>
    <div class="table-wrapper">
      <table>
        <thead>
          <tr>
            <th
              v-for="title in records.dataKeys"
              :key="title"
              @click="handleSort(title.toLowerCase())"  
              :style="{ cursor: 'pointer' }" 
            >
              {{ title }}
              <span v-if="records.sortField === title.toLowerCase()">
                {{ records.sortOrder === 'asc' ? '↑' : '↓' }}
              </span>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="record in records.data" :key="record.ID">
            <td>{{ record.Age }}</td>
            <td>{{ record.Education }}</td>
            <td>{{ record.MaritalStatus }}</td>
            <td>{{ record.Occupation }}</td>
            <td>{{ record.Income }}</td>
          </tr>
        </tbody>
      </table>
    </div>
     <div class="pagination">
      <button :disabled="records.page === 1"  @click="handlePageChange(records.page - 1)">Previous</button>
      <span>Page {{ records.page }} of {{ records.totalPages }}</span>
      <button :disabled="records.page === records.totalPages" @click="handlePageChange(records.page + 1)">Next</button>
    </div> 
    <div class="chart">
      <div class="chartOccupation">
    <ChartOccupations></ChartOccupations>
  </div>
  <div class="chartOccupation">
    <ChartEducation></ChartEducation>
  </div>
  </div>
  </div>
</template>


<style scoped>
.chartOccupation {
  width: 100%;
  max-width: 800px; 
  margin-top: 20px;
}

.chart {
  display: flex;
  flex-direction: column; 
  gap: 100px;
}

@media (min-width: 768px) {
  .chart {
    flex-direction: row; 
  }
}

.table-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 20px;
  font-family: Arial, sans-serif;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.table-wrapper {
  flex: 1; 
  max-height: 400px; 
  overflow-y: auto; 
  border: 1px solid #ccc; 
  margin-top: 20px;
}


table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ccc;
  padding: 10px;
  text-align: left;
  font-size: 14px; 
}

th {
  background-color: #f4f4f4;
}

tbody tr:nth-child(even) {
  background-color: #f5f1f5;
}

.filters {
  display: flex;
  flex-wrap: wrap; 
  gap: 10px;
  margin-bottom: 20px;
}

.filter {
  display: flex;
  text-align: center;
  flex-direction: column;
  min-width: 120px; 
}

select{
  height: 40px;
  border-radius: 10px;
  width: 150px;
  border: 1px solid #cbaacb
}

.promediums {
  display: flex;
  gap: 10px;
  margin-top: 10px;
 
}

@media (min-width: 768px) {
  .promediums {
    flex-direction: row; 
    gap: 20%;
    margin-left: 10px;
  }
}

.promediumAge {
  background-color: #cbaacb;
  justify-content: center;
  text-align: center;
  color: white;
  font-weight: bold;
  border-radius: 10px;
  padding: 10px;
  border: 1px solid black;
}

.promediumHours {
  background-color: #abdee6;
  justify-content: center;
  text-align: center;
  color: white;
  border-radius: 10px;
  padding: 10px;
  font-weight: bold;
  border: 1px solid black;
}

.saveFilters {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap; 
}

.pagination {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 10px;
}

button {
  padding: 5px 10px;
  border: none;
  background-color: #ffccb6;
  color: rgb(63, 63, 63);
  cursor: pointer;
   font-size: 14px;
  border: 1px solid rgb(195, 190, 190)
}

button:hover {
  background-color: #abdee6;
  cursor: pointer;
}
</style>
