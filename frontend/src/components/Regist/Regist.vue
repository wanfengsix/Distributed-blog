<template>
    <div class="container">
    <div class="Header" >
      <h1>注册用户</h1>
    </div>    
    <form @submit.prevent="regist">  
          <div class="form-item">  
            <input type="text" id="username" placeholder="请输入用户名" v-model="username" required>  
          </div>  
          <div class="form-item">  
            <input type="password" id="password" placeholder="请输入密码:" v-model="password" required>  
          </div>
          <div class="form-item">  
            <input type="password" id="password" placeholder="请输入密码:" v-model="again_password" required>  
          </div>
          <div class="form-item">  
            <input type="text" id="pass_word_protection" placeholder="你的生日是什么" v-model="password_protection1" required>  
          </div>
          <div class="form-item">  
            <input type="text" id="pass_word_protection" placeholder="你喜欢吃的水果是什么" v-model="password_protection2" required>  
          </div>
          <div class="form-item">  
            <input type="text" id="pass_word_protection" placeholder="你的妈妈叫什么名字" v-model="password_protection3" required>  
          </div>        
          <div class="auth">
          <button type="submit">注册并登录</button>
          </div>  
    </form>  
    </div>
    </template>
    <script>
    import axios from 'axios'
    export default {
      name: 'Regist_Page',
      props: {
        msg: String
      },  
      data() {  
        return {  
          username: '',  
          password: '',
          again_password:'',
          password_protection1:'',
          password_protection2:'',
          password_protection3:''  
        };  
      }, 
    
      methods: {  
      regist() {
      const data = {
      username: this.username,
      password: this.password,
      password_protection1:this.password_protection1,
      password_protection2:this.password_protection2,
      password_protection3:this.password_protection3
      };
      if(this.password!=this.again_password){
      const message = "两个密码不一致";   
      const popup = window.open("", "", "width=800,height=200");  
      popup.document.write(`<h1>${message}</h1>`);  
      popup.document.close();
      // 设置3秒后关闭窗口  
      setTimeout(function() {  
        popup.close();  
      }, 3000);  
      }
      // 创建并配置 axios 实例  
      const instance = axios.create({    
        withCredentials: true,
      });  
      // 将用户名和密码发送到后端    
      instance.post('http://127.0.0.1:8088/user/regist', data)    
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
    <style>
      body{
        background-image: url("../星空.jpg");
      }
      .container{
        text-align: center;
        justify-content: center;
        
      }
      .Header{
        font-family: "黑体";
        font-size: 46px;
        color:aliceblue;
      }
      input{
        font-size: 40px;
      }
      .form-item{
      margin-bottom: 20px;
      }
      button{
        font-size: 45px;
      }
      .auth{
        justify-content: center;
        display: flex;
        gap: 300px;
      }
    </style>