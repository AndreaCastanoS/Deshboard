<script setup lang="ts">
async function savedFilter() {
  const filtersUserSession = {
    page: Number(sessionStorage.getItem('page')),
    sort_by: sessionStorage.getItem('sortField'),
    order: sessionStorage.getItem('sortOrder'),
    filters: JSON.parse(sessionStorage.getItem('filters') || '{}')
  }

  try {
    const sessionToken = sessionStorage.getItem('token')
    if (!sessionToken) {
      console.error('Error: No se encontró el token de sesión.')
      return
    }

    const response = await fetch('http://18.191.10.34:3333/save-config', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${sessionToken}`
      },
      body: JSON.stringify(filtersUserSession)
    })

    if (response.ok) {
      const data = await response.json()
      console.log('Configuración guardada exitosamente:', data)
    } else if (response.status === 401) {
      console.error('Error: Usuario no autenticado.')
    } else {
      const errorMsg = await response.text()
      console.error('Error al guardar la configuración:', errorMsg)
    }
  } catch (err) {
    console.error('Error en la solicitud:', err)
  }
}
</script>

<template>
  <button @click="savedFilter">Save my filters</button>
</template>

<style>
button {
  height: 50px;
  border-radius: 10px;
  padding: 10px;
  width: 100px;
}
</style>
