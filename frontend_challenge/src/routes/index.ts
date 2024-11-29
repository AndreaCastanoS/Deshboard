import LoginUser from "../components/pages/LoginUser.vue"
import RegisterUser from "@/components/pages/RegisterUser.vue"
import CensusRecords from "@/components/pages/CensusRecords.vue"

export const routes = [
  { path: '/', component: LoginUser },
  { path: '/register', component: RegisterUser},
  { path: '/records', component: CensusRecords},
  { path: '/:pathMatch(.*)*', component: LoginUser }

]
