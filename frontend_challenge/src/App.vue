<script setup lang="ts">
import { RouterView } from 'vue-router'
import { useAuthenticated } from '@/store/auth';
import LogOut from './components/LogOut.vue';
import {onBeforeMount} from 'vue'

const isAuthenticated = useAuthenticated()

 onBeforeMount(isAuthenticated.init)


</script>

<template>
        <div v-if="isAuthenticated.isLoggedIn">
      <LogOut></LogOut>
    </div>
        <RouterView></RouterView>
</template>

<style scoped>
/*  */
</style>
