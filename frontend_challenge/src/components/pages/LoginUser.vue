<script setup lang="ts">
import { useRouter } from 'vue-router'
import { reactive, ref } from 'vue'
import { useAuthenticated } from '@/store/auth'

const router = useRouter()

const userCredentials = reactive({
  email: '',
  password: ''
})

const isAuthenticated = useAuthenticated()
const error = ref<string | null>(null)

async function loginUser() {
  try {
    const response = await fetch('http://18.191.10.34:3333/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(userCredentials)
    })

    if (response.status === 200) {
      const data = await response.json()
      console.log('Inicio de sesión exitoso:', data)
      //isAuthenticated.isAuthenticated = true
      isAuthenticated.setSession(data.token)
      isAuthenticated.sortBy = data.display_config.sort_by
      isAuthenticated.order = data.display_config.order
      isAuthenticated.currentPage = data.display_config.page
      router.push('/records')
    } else if (response.status === 401) {
      error.value = 'Credenciales inválidas'
    } else {
      const errorMsg = await response.text()
      error.value = `Error en el inicio de sesión: ${errorMsg}`
      console.error('Error en login:', errorMsg)
    }
  } catch (err) {
    error.value = 'No se pudo conectar con el servidor'
    console.error('Error en la solicitud:', err)
  }
}
</script>

<template>
  <div class="login-container">
    <form @submit.prevent="loginUser" class="login-form">
      <h2>Iniciar Sesión</h2>

      <div class="form-group">
        <label for="email">Email</label>
        <input
          v-model="userCredentials.email"
          type="email"
          id="email"
          placeholder="Ingresa tu correo electrónico"
          required
        />
      </div>

      <div class="form-group">
        <label for="password">Contraseña</label>
        <input
          v-model="userCredentials.password"
          type="password"
          id="password"
          placeholder="Ingresa tu contraseña"
          required
        />
      </div>
      <button type="submit">Iniciar Sesión</button>
      <p>No tienes cuenta todavia?</p>
      <RouterLink to="/register"><p class="signUp">Registrarse</p></RouterLink>
    </form>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 70vh;
}

.login-form {
  width: 100%;
  max-width: 400px;
  padding: 20px;
  background-color: #f7f7f7;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  text-align: center;
}

h2 {
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
  text-align: left;
}

label {
  font-weight: bold;
  font-size: 0.9rem;
}

input[type='email'],
input[type='password'] {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 1rem;
}

button {
  width: 100%;
  padding: 10px;
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

.signUp {
  color: #007bff;
}

.signUp:hover {
  background-color: #0056b3;
}
</style>
