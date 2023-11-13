<template>
 <div Header >
    <h1>{{ msg }}</h1>
 </div>    
 <form @submit.prevent="login">  
      <div>  
        <input type="text" id="username" placeholder="请输入用户名" v-model="username" required>  
      </div>  
      <div>  
        <input type="password" id="password" placeholder="请输入密码:" v-model="password" required>  
      </div>  
      <button type="submit">登陆</button>  
    </form>  
</template>
<script>
import axios from 'axios'
import CryptoJS from 'crypto-js'
export default {
  name: 'Login_Page',
  props: {
    msg: String
  },  
  data() {  
    return {  
      username: '',  
      password: ''  
    };  
  }, 

  methods: {  
    login() {
      const hashedPassword = CryptoJS.SHA256(this.password).toString(); 
      const data = {
      username: this.username,
      password: hashedPassword
      };    
      // 创建并配置 axios 实例  
      const instance = axios.create({    
        withCredentials: true,
      });  
      // 将用户名和密码发送到后端    
      instance.post('http://127.0.0.1:8088/user/login', data)    
      .then(response => {    
        // 处理成功的响应    
        console.log(response.data);    
        // 在这里可以进行页面跳转或其他操作 
      const message = response.data.message;  
      const count = response.data.count;  
      const popup = window.open("", "", "width=200,height=200");  
      popup.document.write(`<h1>${message}</h1><p>Count: ${count}</p>`);  
      popup.document.close();  
     })      
      .catch(error => {    
        // 处理错误    
        console.error(error);    
      });  
    }    
  } 
};  
</script>