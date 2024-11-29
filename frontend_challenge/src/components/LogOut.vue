<script setup lang="ts">
import { useAuthenticated } from '@/store/auth';
import { useRouter} from 'vue-router';

const isAuthenticated = useAuthenticated()
const router = useRouter();

async function logout() {
  try {
    const sessionToken = sessionStorage.getItem('token')
    const response = await fetch('http://localhost:3333/auth/logout', {
      method: 'POST',
      headers: {
      Authorization: `Bearer ${sessionToken}`,
      }
    });

    if (response.ok) {
      console.log('Successfully logged out');
      isAuthenticated.clearSession()
      router.push('/'); 
    } else {
      const errorMsg = await response.text();
      console.error('Error logging out:', errorMsg);
    }
  } catch (error) {
    console.error('Error in the request:', error);
  }
}

</script>

<template>
    <button @click="logout">Cerrar sesion</button>
</template>

<style scoped>
button {
  width: 50%;
  padding: 10px;
  margin-left: 30%;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #0056b3;
}

</style>