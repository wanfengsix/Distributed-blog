<template>
  <router-view></router-view>
</template>

<script>
import Login from './components/Login/Login.vue'
import Regist from './components/Regist/Regist.vue'
import { createRouter, createWebHistory } from 'vue-router'
const routes = [
    { path: '/login', component: Login },
    { path: '/regist', component: Regist },
]
  
  // 3. 创建路由实例并传递 `routes` 配置
  // 你可以在这里输入更多的配置，但我们在这里
  // 暂时保持简单
const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHistory(),
    routes, // `routes: routes` 的缩写
})
export default {
  name: 'App',
  components: {
  },
  router
}
</script>

<style>

</style>
