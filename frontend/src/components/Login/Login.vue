<template>
<head>
<link rel="stylesheet" type="text/css" href="./Login.css">
</head>
 <div Header >
    <h1>{{ msg }}</h1>
 </div>    
 <div class="container">
	<div class="row clearfix">
		<div class="col-md-12 column">
			<form role="form">
				<div class="Usr_name">
					<label for="exampleInputUser_name"></label><input :input="username" type="text" placeholder="请输入用户名" class="form-control" id="InputUser_name" />
				</div>
				<div class="Ps_word">
					<label for="InputPassword"></label><input :input="password" type="password" class="form-control" placeholder="清输入密码" id="InputPassword" />
				</div>
				<button type="登录" @click="login" class="btn_Login">登陆</button>
			</form>
		</div>
	</div>
</div>
</template>
<script>
import axios from 'axios'
	
export default {
  name: 'Login_Page',
  props: {
    msg: String
  },
  methods:{
	data(){
		return {
		username:'',
		password:''
		};
	},
	login() {  
      // 将用户名和密码发送到后端  
      axios.post('127.0.0.1:8088/api/login', {  
        username: this.username,  
        password: this.password  
      })  
      .then(response => {  
        // 处理成功的响应  
        console.log(response.data);  
        // 在这里可以进行页面跳转或其他操作  
      })  
      .catch(error => {  
        // 处理错误  
        console.error(error);  
      });  
    }  
  }
}

</script>