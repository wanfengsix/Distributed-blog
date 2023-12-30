<template>
	<head>
		<meta charset="utf-8">
		
	</head>

	<body>
		<div class="head">
			<div class="main">
				<div class="head1">
					<div class="logo">
						<img src="../../../public/img/logo.png">
					</div>

					<div class="geren">
						<ul>
							<li><a class="a1" href="/">首页</a></li>
							<div class="clear">

							</div>
						</ul>
					</div>


					<div class="input">
						<input class="b1" type="text" name="" id="" value="" />
						<a class="b2" href="###"><img src="../../../public/img/fdj.png"></a>
					</div>



					<div class="san">
						<a href="###">创作者中心</a>
					</div>

					<div v-if="isLoggedIn" class="loggedI">
        <div class="navgationbarItemAvat">
          <img class="imgAvat" :src="imageSrc" alt="Image from backend" />
        </div>

		
		<div class="navigationbarItemLo" style="position: relative;">  
          <button @click="logout" style="width: 200px; height: 50px;font-size: 30px;">
            退出
          </button>
        </div>
      </div>
      <div v-else class="nologgedI">
        <div class="navgationbarItemLog">
          <a href="login">登录</a>
        </div>
        <div class="navgationbarItemRegister">
          <a href="regist">注册</a>
        </div>
      </div>


					<div class="clear">

					</div>

				</div>
			</div>
		</div>



		<div class="centen">
			<div class="main">

				<div class="centen1">

					<div class="left">
						<span class="aw1">通知</span>
						<p>“用户123456"给你的文章“鸿蒙OS起步照官网练习创建第一页面和跳转的第二页面需要注意的点”点赞了。<br><br>
							“用户123456”给你的文章“鸿蒙OS起步照官网练习创建第一页面和跳转的第二页面需要注意的点”点赞了。<br><br>
							“用户123456”给你的文章“鸿蒙OS起步照官网练习创建第一页面和跳转的第二页面需要注意的点”点赞了。<br><br>
							“用户123456”给你的文章“鸿蒙OS起步照官网练习创建第—页面和跳转的第二页面需要注意的点”点赞了。<br><br>
							“用户123456”给你的文章“鸿蒙OS起步照官网练习创建第一页面和跳转的第二页面需要注意的点”点赞了。<br><br>
							“用户123456”给你的文章“鸿蒙OS起步照官网练习创建第一页面和跳转的第二页面需要注意的点”点赞了。<br><br>
							“用户123456”给你的文章“鸿蒙OS起步照官网练习创建第一页面和跳转的第二页面需要注意的点”点赞了。<br><br>
							“用户123456”给你的文章“鸿蒙OS起步照官网练习创建第一页面和跳转的第二页面需要注意的点”点赞了。<br><br>
						</p>
					</div>


					<div class="zhong">
						<div class="shang">
							<div class="zuo">
								<img :src="imageSrca" >
							</div>
							
							<div class="zhong">  
								<h3>{{u_name}}</h3>  
								<h3>LV.{{level}}</h3>  
								<span>个性签名</span>  
								<input v-model="signature" v-if="isSetting" type="text" placeholder="请输入个性签名">  
								<p v-else>{{signature}}</p>  
							  </div>  
							  <div class="you">  
								<button v-if="!isSetting" @click="isSetting = true">设置</button>  
								<button v-else @click="sendSignature()">完成</button>  
							  </div>  
							
							<div class="clear">
								
						</div>
							
					</div>
						
						
						<div class="xia">
						
							<div class="articleRecommendedLis">
								<!-- 通过循环生成文章列表 -->
								<div
								  class="article"
								  v-for="(item, index) in articleList.article_list"
								  :key="index"
								>
								  <h3>
									<a :href="'article/' + item.article_id">{{ item.head }}</a>
								  </h3>
								  <p>文章摘要或内容简介...</p>
								</div>
							</div>
							<div class="clear">

							</div>
						
							
						</div>
						
					</div>

					<div class="right">
						<div class="shang">
							<h3>关注了 &nbsp;&nbsp;丨&nbsp;&nbsp; 关注者<br></h3>
							<span>{{ following }}</span> <span>{{ followed }}</span>
						</div>
						
						<div class="xia">  
							  
							<div class="lists-container">  
							  <div class="list followers-list">  
								<h4>关注列表</h4>  
								<ul>  
								  <li v-for="follower in followersList" :key="follower.UID">  
									<a :href="follower.UID">{{ follower.u_name }}</a>
								  </li>  
								</ul>  
							  </div>  
							  <div class="list fans-list">  
								<h4>粉丝列表</h4>  
								<ul>  
								  <li v-for="fan in fansList" :key="fan.UID">  
									<a :href="fan.UID">{{ fan.u_name }}</a>
								  </li>  
								</ul>  
							  </div>  
							</div>  
						  </div>
						
						
					</div>

					



				</div>

			</div>


		</div>


	</body>

</template>
<script>  
import axios, { Axios } from 'axios';
export default {  
	name: 'Host_Page',
	data() {
    return {
	  isLoggedIn: localStorage.getItem("isLoggedIn"),
      username: localStorage.getItem("username"),
      imageSrc:"",
	  imageSrca:"",
	  signature:"",
	  isSetting: false ,
	  articleId: "0",
	  articleList: "",
	  level:0,
	  following: 0,
	  followed: 0,
	  followersList:[], // 关注列表数据  
      fansList:[] , // 粉丝列表数据  
	  uid: this.$route.params.uid,
	  u_name:"",
    };
  },
  created() {
	if (localStorage.getItem("isLoggedIn")!=false){
    this.fetchAvatar(); // 在页面加载时调用fetchAvatar方法

    //this.fetchArticleList();
	//this.getfollowing();
	//this.getfollowed();
	this.getfollowers_list();
	this.getfans_list();
	this.getu_name();
	

	}
  },
  methods: {  
	updateImage() {  
      // 在这里更新 imageSrc 的值   
      this.imageStyle = {  
        width: '10px',  
        borderRadius: '75%', 
      };  
    } , 
	sendSignature() {  
      const data ={
		name:this.u_name,
		resource_type:"signature",
		post_data: this.signature,

	  };
	  const instance = axios.create({
        withCredentials: true,
      });
	  //将signature发送到后端
	  instance
.post(`http://127.0.0.1:8088/user/signature/${this.u_name}`,data)
        .then((response) => {
          // 处理成功的响应
          
          if (response.data.Success == true) {
			this.isSetting = false; // 更新isSetting的值，使文本框变为不可编辑状态  

            alert("您已修改成功");
            
          } else {
            alert("修改失败");
          }
		  this.$forceUpdate(); // 强制Vue组件重新渲染  
        })
        .catch((error) => {
          // 处理错误
          console.error(error);
        });
    },  
    logout() {  
		this.isLoggedIn=false;
		localStorage.setItem("isLoggedIn",false)
      // 在这里编写退出登录的逻辑，比如清除本地存储的用户信息等。  
      // 这里只是一个示例，具体的实现取决于您的应用需求。  
    }  ,
	fetchAvatar() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance.get(`http://127.0.0.1:8088/user/avatar/${this.username}`) // 使用get请求获取头像图片文件
        .then(async response => {
         
          this.imageSrc = "data:img/png;base64,"+response.data.data; // 更新imageSrc以显示头像  
        })
        .catch(error => {
          console.error(error);
        });
  },
  fetchAvatara() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance.get(`http://127.0.0.1:8088/user/avatar/${this.u_name}`) // 使用get请求获取头像图片文件
        .then(async response => {
        
          this.imageSrca = "data:img/png;base64,"+response.data.data; // 更新imageSrc以显示头像  
		  
        })
        .catch(error => {
          console.error(error);
        });
  },
  getSignature(){
	const instance = axios.create({
        withCredentials: true,
      });
      instance.get(`http://127.0.0.1:8088/user/signature/${this.u_name}`) // 
        .then(async response => {
         
          this.signature = response.data.data; // 更新
        })
        .catch(error => {
          console.error(error);
        });
  },
  /*
  getfollowing(){
	const instance = axios.create({
        withCredentials: true,
      });
	  
      instance.get(`http://127.0.0.1:8088/userinfo/u_name/${this.username}`) 
        .then(async response => {
          console.log(response.data)
          this.following = response.data.data.following; // 更新关注者数量
		  this.followed=response.data.data.followed  
        })
        .catch(error => {
          console.error(error);
        });
  },
  getfollowed(){
	const instance = axios.create({
        withCredentials: true,
      });
	  console.log(this.u_name)
      instance.get(`http://127.0.0.1:8088/userinfo/u_name/${this.u_name}`) 
        .then(async response => {
          console.log(response.data)
          this.followed = response.data.data.followed; // 更新粉丝数量
        })
        .catch(error => {
          console.error(error);
        });
  },
  */
	fetchArticleList() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/article-list/${this.u_name}`) // 使用get请求获取文章标题
        .then(async (response) => {
         
          this.articleList = response.data; // 显示文章内容

          // this.articleList = make([]interface{},len(response.data.articleList))
        })
        .catch((error) => {
          console.error(error);
        });
    },
	getfollowers_list() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/userinfo/following-list/${this.username}`) // 使用get请求获取关注列表
        .then(async (response) => {
         
          this.followersList = response.data.List; // 显示关注列表
        })
        .catch((error) => {
          console.error(error);
        });
    },
	getfans_list() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/userinfo/followed-list/${this.username}`) // 使用get请求获取粉丝列表
        .then(async (response) => {
        
          this.fansList = response.data.List; // 显示粉丝列表
        })
        .catch((error) => {
          console.error(error);
        });
    },
	getu_name() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/userinfo/uid/${this.uid}`) // 使用get请求获取u_name
        .then(async (response) => {
          this.u_name = response.data.data.u_name.String; //返回u_name值
		  this.followed=response.data.data.followed;
		  this.following =response.data.data.following;
		  this.level=response.data.data.level;
		  this.getSignature();    
		  this.fetchArticleList();
		  this.fetchAvatara();
		 
        })
        .catch((error) => {
          console.error(error);
        });
    },
  }  
}  
</script>
<style>
/*通用类*/
* {
	margin: 0;
	padding: 0;
}

body {
	min-width: 1400px;
	margin: 0 auto;
	font-size: 14px;
	background-color: #F5E9D9;
	text-align: justify;
	color: #000;
}

h1,
h2,
h3 {
	font-weight: 500;
}

img {
	border: none;
}

a {
	cursor: pointer;
	color: #000;
	text-decoration: none;
	outline: none;
}

a:hover {
	color: #f60;
}

ul {
	list-style-type: none;
}

div.clear {
	font: 0px Arial;
	line-height: 0;
	height: 0;
	overflow: hidden;
	clear: both;
}



/*公用CSS 开始*/
.main {
	width: 1400px;
	margin: 0 auto;
}

.head {
	width: 100%;
	background-color: #EFEFEF;
}

.head1 {
	padding: 20px 0;
}

.head1 .logo {
	width: 10%;
	float: left;
}

.head1 .geren {
	width: 18%;
	float: left;
}

.head1 .geren ul li {
	width: 49%;
	float: left;
}

.head1 .geren ul li a {
	background-color: #FFFFFF;
	height: 50px;
	display: block;
	text-align: center;
	line-height: 50px;
	border: 1px solid #515151;
}

.head1 .geren ul li .a1 {
	/*background: url(../images/fz.png)no-repeat left;
	*/background-color: #FFFFFF;
	background-size: 18px;
	background-position-x: 5px;
}

.head1 .geren ul li .a2 {
	/*background: url(../images/grzx.png)no-repeat left;*/
	background-color: #FFFFFF;
	background-size: 18px;
	background-position-x: 5px;
}

.head1 .geren ul li a:hover {
	background-color: #E7F8FF;
}

.head1 .input {
	width: 25%;
	float: left;
	margin: 0 20px;
	margin-top: 10px;
}

.head1 .input .b1 {
	width: 310px;
	float: left;
	height: 30px;
}

.head1 .input .b2 {
	width: 20px;
	float: right;
}

.head1 .san {
	width: 10%;
	float: left;
	margin: 0 30px;
	margin-top: 10px;
}

.head1 .san a {
	display: block;
	background-color: #B63927;
	color: #FFFFFF;
	text-align: center;
	height: 30px;
	line-height: 30px;
	text-align: center;
}
.navgationbarItemAvat{  
	width: 50px; /* 设置头像的宽度 */  
	height: 50px; /* 设置头像的高度 */  
	border-radius: 50%; /* 设置头像的边框半径为50%，使其呈现圆形 */  
	margin-left: -100px; /* 将头像向左移动10像素 */
  }  
	
  .imgAvat {  
	width: 100%; /* 设置图片的宽度为100% */  
	height: 100%; /* 设置图片的高度为100% */  
  }
  
  .navigationbarItemLo button {  
	position: relative;  
	top: -50px; /* 调整为适合你的距离 */  
  }
.head1 .nologgedI{
	width: 30%;
	float: right;
	text-align: center;
}
.nologgedI {  
	display: flex;  
	justify-content: space-between;  
	align-items: center;
	position: relative;  
	top: -30px; /* 将整个导航栏向上移动20像素 */  
  }  
	

  .navgationbarItemLog {
	display: inline-block;  
	padding: 10px;  
	font-size: 16px;  
	width: 40%; /* 你可以根据需要调整宽度 */  
  }
  .navgationbarItemRegister {  
	display: inline-block;  
	padding: 10px;  
	font-size: 16px;  
	width: 40%; /* 你可以根据需要调整宽度 */  
  }  
	
 

  .head1 .loggedI{
	width: 10%;
	float: right;
	text-align: center;
}
.si /*img[src="path/to/image.png"]*/ {  
  width: 60px;  
  border-radius: 50%;  
}



.head1 .wu {
	width: 10%;
	float: left;
	margin-top: 1000px;
}
.articleRecommendedLis {  
	margin-top: 50px; /* 调整这个值来控制向下移动的距离 */  
  }

.head1 .wu a {
	display: block;
	background-color: #4E93DE;
	color: #000;
	text-align: center;
	height: 40px;
	line-height: 40px;
	text-align: center;
	border-radius: 15px;
}

.lists-container {  
	display: flex; /* 使用Flex布局 */  
	justify-content: space-between; /* 使两个列表水平排列，并在它们之间添加间距 */  
  }  
	
  .list {  
	width: 45%; /* 控制每个列表的宽度 */  
  }
.centen {
	background-color: #F5E9D9;
	margin-top: 50px;
}

.centen1 .left {
	width: 20%;
	float: left;
}

.centen1 .left .aw1 {
	text-align: center;
	margin: 0 auto;
	margin-bottom: 40px;
}

.centen1 .left span {
	width: 200px;
	height: 60px;
	line-height: 60px;
	color: #A0404E;
	border: 1px solid #A0404E;
	display: block;
	text-align: center;
	background-color: #FFFFFF;
	border-radius: 15px;
}

.centen1 .left p {
	background-color: #FFFFFF;
	padding: 20px;
	box-sizing: border-box;
	font-size: 17px;
}

.centen1 .zhong {
	width: 50%;
	float: left;
	margin: 0 40px;
}

.centen1 .zhong .shang {
	background-color: #FFFFFF;
	height: 300px;
}

.centen1 .zhong .shang .zuo {
	width: 20%;
	float: left;
	text-align: center;
}

.centen1 .zhong .shang .zuo img {
	margin-top: 80px;
	width: 100px;
}

.centen1 .zhong .shang .zhong {
	width: 50%;
	float: left;
	margin: 0 10px;
}

.centen1 .zhong .shang .zhong h3 {
	margin-top: 50px;
}

.centen1 .zhong .shang .zhong img {
	display: block;
	
	margin: 5px 0;
}

.centen1 .zhong .shang .zhong span {
	display: inline-block;
	width: 86px;
	height: 20px;
	border: 1px solid #ccc;
	margin: 5px 0;
	color: #4E93DE;
	font-size: 18px;
	text-align: center;
	font-weight: bold;
}

.centen1 .zhong .shang .zhong p {
	border: 1px solid #515151;
	padding: 30px;
	margin-top: 20px;
}

.centen1 .zhong .shang .you {
	width: 20%;
	float: left;
	margin-left: 20px;
	margin-top: 190px;
}

.centen1 .zhong .shang .you a {
	display: inline-block;
	width: 100px;
	height: 50px;
	line-height: 50px;
	border: 1px solid #ABCCDD;
	color: #ABCCDD;
	text-align: center;
	font-size: 20px;
	font-weight: bold;
}


.centen1 .zhong .xia{background-color: #FFFFFF;margin-top: 30px;padding-bottom: 100px;}
.centen1 .zhong .xia ul li{padding: 15px;box-sizing: border-box;}
.centen1 .zhong .xia ul li .yi{width: 50%;float: left;}
.centen1 .zhong .xia ul li .yi h3{font-size: 14px;font-weight: bold;}
.centen1 .zhong .xia ul li .yi p{font-size: 12px;color: #515151;margin: 10px 0;}
.centen1 .zhong .xia ul li .yi span{display: inline-block;width: 45px;}

.centen1 .zhong .xia ul li .er{width: 35%;float: left;margin: 0 15px;}
.centen1 .zhong .xia ul li .er img{width:100%;}
.centen1 .zhong .xia ul li .san{width: 8%;float: left;}
.centen1 .zhong .xia ul li .san a{display: block;margin-top: 20px;}

.centen1 .right{width: 20%;float: right;}

.centen1 .right .shang{background-color: #FFFFFF;padding: 20px 0;}
.centen1 .right .shang h3{text-align: center;font-size: 22px;}
.centen1 .right .shang  span{font-size: 22px;width: 48%;text-align: center;display: inline-block;}

.centen1 .right .xia{text-align: center;font-size: 22px;margin: 30px 0;}

.centen1 .right .xia .bg1{background-color: #FFFFFF;padding: 15px;margin-top: 20px;height: 800px;}
.centen1 .right .xia .bg1 ul li{width: 100%;margin-top: 15px;}
.centen1 .right .xia .bg1 ul li a{display: block;background: url/*(../images/grzx.png)*/no-repeat left;
padding-left: 20px;
background-size: 19px;
text-align: left;padding-left: 30px;}	
</style>
