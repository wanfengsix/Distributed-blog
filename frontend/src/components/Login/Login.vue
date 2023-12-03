<template>
  <div class="container">
    <div class="Header">
      <h1>dusha</h1>
    </div>
    <form @submit.prevent="login" class="form">
      <div class="form-item">
        <input type="text" id="username" placeholder="请输入用户名" v-model="username" required>
      </div>
      <div class="form-item">
        <input type="password" id="password" placeholder="请输入密码:" v-model="password" required>
      </div>
      <div class="auth">
        <button type="submit">登陆</button>
        <button @click.prevent="Go_Regist">注册</button>
      </div>
    </form>
  </div>
</template>

<style scoped>
.container {
  background-image: url("../星空.jpg");
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.Header {
  font-family: "黑体";
  font-size: 46px;
  color: aliceblue;
}

input {
  font-size: 40px;
}

.form {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.form-item {
  margin-bottom: 20px;
}

button {
  font-size: 45px;
}

.auth {
  display: flex;
  gap: 300px;
}
</style>

<script>
import axios from 'axios'
import CryptoJS from 'crypto-js'
export default {
  name: 'Login_Page',
  data() {  
    return { 
      username: '',  
      password: ''  
    };  
  }, 

  methods: {
    Go_Regist(){
      this.$router.push('/regist')
    },  
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
      instance.post('http://127.0.0.1:8088/user/login'+'/'+hashedPassword, data)    
      .then(response => {    
        // 处理成功的响应    
      console.log(response.data);
      if(response.data.Success==true){
        this.$router.push('/')   
      }else{
        // 在这里可以进行页面跳转或其他操作 
      const message = response.data.message;   
      const popup = window.open("", "", "width=800,height=200");  
      popup.document.write(`<h1>${message}</h1>`);  
      popup.document.close();
      // 设置3秒后关闭窗口  
      setTimeout(function() {  
        popup.close();  
      }, 3000);   
      }
     })      
      .catch(error => {    
        // 处理错误    
        console.error(error);    
      });  
    }    
  } 
};  
</script>