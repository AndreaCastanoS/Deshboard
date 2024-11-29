<script setup lang="ts">
import { useCensusRecords } from '../store/record'

const records = useCensusRecords()

function getBarHeight(count: number): number {
  const maxCount = Math.max(...Object.values(records.educationCounts))
  return (count / maxCount) * 200
}
</script>

<template>
  <div>
    <h3>Educations with income >50</h3>
    <div class="bar-chart">
      <div
        v-for="(count, education) in records.educationCounts"
        :key="education"
        class="bar"
        :style="{ height: getBarHeight(count) + 'px' }"
      >
        <span class="bar-label">{{ education }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.bar-chart {
  display: flex;
  justify-content: space-around;
  align-items: flex-end;
  flex-wrap: wrap;
  height: 250px;
  border: 1px solid #ccc;
  margin-top: 20px;
  padding: 10px;
  gap: 10px;
}

h3 {
  text-align: center;
}

.bar {
  flex: 1;
  max-width: 60px;
  min-width: 20px;
  background-color: #ffffb5;
  text-align: center;
  color: black;
  border-radius: 5px;
  position: relative;
  border: 1px solid rgb(195, 190, 190);
}

.bar-label {
  position: absolute;
  bottom: -50px;
  right: 0;
  left: 0;
  text-align: center;
  margin-bottom: 10px;
  font-size: 12px;
  overflow-wrap: break-word;
}

@media (max-width: 768px) {
  .bar-chart {
    height: 200px;
  }
  .bar {
    max-width: 40px;
  }
  .bar-label {
    font-size: 10px;
  }
}

@media (max-width: 480px) {
  .bar {
    max-width: 30px;
    min-width: 15px;
  }
  .bar-label {
    font-size: 8px;
  }
}
</style>
