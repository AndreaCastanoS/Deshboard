<script setup lang="ts">
import { RouterLink, useRouter } from 'vue-router'
import { ref, reactive } from 'vue'

const router = useRouter()

interface User {
  id: number
  email: string
  password: string
}

const newUser = reactive({
  email: '',
  password: ''
})

const users = ref<User[]>([])

async function createUser() {
  try {
    const response = await fetch('http://18.191.10.34:3333/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newUser)
    })
    if (response.status === 201) {
      const data = await response.json()
      users.value = [...users.value, data]
      console.log(response.status)
      console.log(data)
      router.push('/')
    } else if (response.status === 409) {
      console.log('El email ya está registrado')
    }
  } catch (err) {
    if (err && typeof err === 'object' && 'value' in err) {
      console.error((err as any).value) 
    } else {
      console.error('Error desconocido:', err)
    }
    console.error('Error en la solicitud:', err)
  }
}
</script>

<template>
  <div class="login-container">
    <form @submit.prevent="createUser" class="login-form">
      <h2>Registrarse</h2>

      <div class="form-group">
        <label for="email">Email</label>
        <input
          v-model="newUser.email"
          type="email"
          id="email"
          placeholder="Ingresa tu correo electrónico"
          required
        />
      </div>

      <div class="form-group">
        <label for="password">Contraseña</label>
        <input
          v-model="newUser.password"
          type="password"
          id="password"
          placeholder="Ingresa tu contraseña"
          required
        />
      </div>
      <button type="submit">Registrarse</button>
      <RouterLink to="/"> <p class="signUp">Iniciar Sesión</p></RouterLink>
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
  text-decoration: none;
}

.signUp:hover {
  background-color: #0056b3;
}
</style>
