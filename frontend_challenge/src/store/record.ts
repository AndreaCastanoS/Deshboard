import { defineStore } from 'pinia'
import type { DataRecord } from '@/types'
import { ref, watch, toRaw } from 'vue'
import { useAuthenticated } from '@/store/auth'

export const useCensusRecords = defineStore('records', () => {
  const occupationCounts = ref<Record<string, number>>({})
  const educationCounts = ref<Record<string, number>>({})
  const isAuthenticated = useAuthenticated()
  const promediumAge = ref<number>(0)
  const promediumHours = ref<number>(0)
  const data = ref<DataRecord[]>([])
  const sortField = ref<string | null>(null)
  const sortOrder = ref<'asc' | 'desc'>('asc')
  const filters = ref<Record<string, string>>({})
  const page = ref<number>(1)
  const itemsPerPage = ref<number>(100)
  const totalPages = ref(1)
  const highIncomeOccupations = ref<string[]>([])
  const highIncomeEducation = ref<string[]>([])

  const dataKeys = ref<(keyof DataRecord)[]>([
    'Age',
    'Education',
    'MaritalStatus',
    'Occupation',
    'Income'
  ])

  function updateOccupationCounts() {
    occupationCounts.value = highIncomeOccupations.value.reduce(
      (acc: Record<string, number>, occupation: string) => {
        acc[occupation] = (acc[occupation] || 0) + 1
        return acc
      },
      {}
    )
  }

  function updateEducationCounts() {
    educationCounts.value = highIncomeEducation.value.reduce(
      (acc: Record<string, number>, education: string) => {
        acc[education] = (acc[education] || 0) + 1
        return acc
      },
      {}
    )
  }

  async function fetchData(
    page: number = 1,
    sortField: string | null,
    sortOrder: 'asc' | 'desc',
    filters: Record<string, string>
  ) {
    try {
      const sessionToken = sessionStorage.getItem('token')
      const url = 'http://18.191.10.34:3333/records'
      const queryParams: Record<string, string> = {
        sort_by: sortField || '',
        order: sortOrder,
        page: page.toString(),
        limit: itemsPerPage.value.toString(),
        ...filters
      }

      const urlParams = new URLSearchParams(queryParams)
      const response = await fetch(`${url}?${urlParams.toString()}`, {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${sessionToken}`
        }
      })

      if (response.ok) {
        const result = await response.json()
        data.value = result.data
        totalPages.value = result.totalPages
        let totalAge = 0
        let totalHours = 0
        let count = 0
        result.data.forEach((e: any) => {
          ;(totalAge += e.Age), (totalHours += e.HoursPerWeek || 0)
          count++

          if (e.Income) {
            if (e.Income.includes('>50K')) {
              highIncomeOccupations.value.push(e.Occupation)
              highIncomeEducation.value.push(e.Education)
            }
          }
        })
        promediumAge.value = count > 0 ? totalAge / count : 0
        promediumHours.value = count > 0 ? totalHours / count : 0
        updateOccupationCounts()
        updateEducationCounts()
      } else {
        console.error('Error al obtener los datos:', await response.text())
      }
    } catch (err) {
      console.error('Error de conexión:', err)
    }
  }

  function deepToRaw(obj: any): any {
    if (Array.isArray(obj)) {
      return obj.map(deepToRaw)
    } else if (obj && typeof obj === 'object') {
      const rawObj = toRaw(obj)
      return Object.fromEntries(
        Object.entries(rawObj).map(([key, value]) => [key, deepToRaw(value)])
      )
    }
    return obj
  }

  function saveStateToSession() {
    const rawFilters = deepToRaw(filters.value)
    sessionStorage.setItem('filters', JSON.stringify(rawFilters))
    sessionStorage.setItem('sortField', sortField.value || '')
    sessionStorage.setItem('sortOrder', sortOrder.value)
    sessionStorage.setItem('page', page.value.toString())
    console.log('Estado guardado en sessionStorage:', {
      filters: rawFilters,
      sortField: sortField.value,
      sortOrder: sortOrder.value
    })
  }

  function loadStateFromSession() {
    const storedFilters = sessionStorage.getItem('filters')
    const storedSortField = sessionStorage.getItem('sortField')
    const storedSortOrder = sessionStorage.getItem('sortOrder')
    const storedPage = sessionStorage.getItem('page')

    if (storedFilters) {
      filters.value = JSON.parse(storedFilters)
    }

    if (storedSortField) {
      sortField.value = storedSortField
    } else if (isAuthenticated.sortBy) {
      sortField.value = isAuthenticated.sortBy
    }

    if (storedSortOrder) {
      sortOrder.value = storedSortOrder as 'asc' | 'desc'
    } else if (isAuthenticated.order) {
      sortOrder.value = isAuthenticated.order as 'asc' | 'desc'
    }

    if (storedPage) {
      page.value = parseInt(storedPage, 10)
    } else if (isAuthenticated.currentPage) {
      page.value = parseInt(isAuthenticated.currentPage)
    } else {
      page.value = parseInt('1')
    }
  }

  watch([filters, sortField, sortOrder, page], saveStateToSession, { deep: true })
  watch(
    () => totalPages.value,
    (newTotalPages) => {
      if (page.value > newTotalPages) {
        page.value = newTotalPages
      }
    }
  )

  return {
    dataKeys,
    sortField,
    sortOrder,
    data,
    filters,
    fetchData,
    loadStateFromSession,
    page,
    totalPages,
    promediumAge,
    promediumHours,
    highIncomeOccupations,
    occupationCounts,
    updateOccupationCounts,
    educationCounts
  }
})
