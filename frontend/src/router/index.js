import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/LoginView.vue')
  },
  {
    path: '/balance',
    name: 'balance',
    component: () => import('../views/BalanceView')
  },
  {
    path: '/top-up',
    name: 'top-up',
    component: () => import('../views/TopupView')
  },
  {
    path: '/token',
    name: 'token',
    component: () => import('../views/TokenView')
  },
  {
    path: '/registration',
    name: 'registration',
    component: () => import('../views/RegistrationView')
  },
  {
    path: '/html',
    name: 'html',
    component: () => import('../views/HtmlView')
  },
  {
    path: '/url',
    name: 'url',
    component: () => import('../views/UrlView.vue')
  },
  {
    path: '/result',
    name: 'result',
    component: () => import('../views/ResultView')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
