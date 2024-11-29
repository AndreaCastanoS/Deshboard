import type { JWTPayload } from '@/types'
import { jwtDecode } from 'jwt-decode'
import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

export const useAuthenticated = defineStore('authenticated', () => {
  const session = ref('')
  const sortBy = ''
  const order = ''
  const currentPage = ''

  const isLoggedIn = computed(() => !!session.value)

  async function init() {
    const tokenStr = sessionStorage.getItem('token')
    if (tokenStr) {
      setSession(tokenStr)
    }
  }

  function setSession(tokenStr: string) {
    const payload = jwtDecode(tokenStr) as JWTPayload
    const now = new Date()
    const diff = payload.MapClaims.eat * 1000 - now.getTime()

    sessionStorage.setItem('token', tokenStr)
    session.value = payload.session
    setTimeout(() => {
      clearSession()
    }, diff)
  }

  function clearSession() {
    session.value = ''
  }
  return { init, isLoggedIn, clearSession, setSession, sortBy, order, currentPage }
})
