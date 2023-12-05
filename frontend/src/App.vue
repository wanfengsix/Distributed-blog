<template>
  <router-view></router-view>
</template>

<script>
import Host from './components/Host/Host.vue'
import Login from './components/Login/Login.vue'
import Article from './components/Article/Article.vue'
import Regist from './components/Regist/Regist.vue'
import { createRouter, createWebHistory } from 'vue-router'
import Individual from './components/Individual_center/Individual.vue'
const routes = [
    {path: '/', component:Host},
    { path: '/login', component: Login },
    { path: '/regist', component: Regist },
    {path:'/article',component:Article},
    {path:'/individual',component:Individual}
]
  
  // 3. 创建路由实例并传递 `routes` 配置
  // 你可以在这里输入更多的配置，但我们在这里
  // 暂时保持简单
const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHistory(),
    routes, // `routes: routes` 的缩写
})
// 添加全局前置守卫
router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const isLoggedIn = localStorage.getItem('isLoggedIn')   

  if (requiresAuth && !isLoggedIn) {
    next('/login')           //重定向到登录
  } else {
    next()
  }
})
export default {
  name: 'App',
  components: {
  },
  router,
  computed: {  
    bgStyle() {  
      return {  
        'background-image': 'none', // 这将覆盖所有子组件的背景图片样式  
      }  
    }  
  },  
}
</script>

<style scoped>
.app{
  background-color: aliceblue;
}
</style>
