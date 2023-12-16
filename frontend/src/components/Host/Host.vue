<template>
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>个人博客系统</title>
  </head>

  <body>
    <div class="navgationbar">
      <img class="imgLogo" src="img/logo.jpg" />
      <div class="navgationbarItemHome">
        <img class="imgHome" src="img/home.png" alt="" />
        <a href="/">首页</a>
      </div>
      <div class="navgationbarItemNotifications">
        <img class="imgNotifications" src="img/notifications.png" alt="" />
        <a href="#notifications">通知</a>
      </div>
      <div class="navgationbarItemProfile">
        <img class="imgProfile" src="img/profile.png" alt="" />
        <a href="/individual">个人中心</a>
      </div>
      <div class="navgationbarItemSearch">
        <img class="imgSearch" src="img/search.png" />
        <input type="text" style="width: 400px" />
      </div>
      <div class="navgationbarItemCreaterCenter">
        <p style="color: white">创作者中心</p>
      </div>
      <div v-if="isLoggedIn" class="loggedIn">
        <div class="navgationbarItemAvator">
          <img class="imgAvator" :src="imageSrc" alt="Image from backend" />
        </div>

        <button
          @click="logout"
          class="navgationbarItemLogOut"
          style="width: 200px; height: 50px; font-size: 20px  margin-left: 300px; "
        >
          退出
        </button>
      </div>
      <div v-else class="nologgedIn">
        <div class="navgationbarItemRight">
          <div class="navgationbarItemLog">
            <a href="login">登录</a>
          </div>
          <div class="navgationbarItemRegister">
            <a href="regist">注册</a>
          </div>
        </div>
      </div>
    </div>
    <section id="home">
      <div class="articleList">
        <!-- 通过循环生成文章列表 -->

        <ol>
          <div class="articleListTiTle">
            <img src="img/list.png" style="height: 20px" />
            文章榜
          </div>

          <li class="articleListItem">
            <h3><a class="articleItemTitle" href="#article1">文章标题</a></h3>
          </li>
          <li class="articleListItem">
            <h3><a class="articleItemTitle" href="#article1">文章标题</a></h3>
          </li>
          <li class="articleListItem">
            <h3><a class="articleItemTitle" href="#article1">文章标题</a></h3>
          </li>
          <li class="articleListItem">
            <h3><a class="articleItemTitle" href="#article1">文章标题</a></h3>
          </li>
          <li class="articleListItem">
            <h3><a class="articleItemTitle" href="#article1">文章标题</a></h3>
          </li>
          <li class="articleListItem">
            <h3><a class="articleItemTitle" href="#article1">文章标题</a></h3>
          </li>
        </ol>
        <li class="articleListMore">...查看更多</li>
        <!-- 可以添加更多文章项 -->
      </div>
      <div class="authorList">
        <!-- 通过循环生成文章列表 -->
        <ol>
          <div class="authorListTiTle">
            <img src="img/list.png" style="height: 20px" />
            作者榜
          </div>
          <li class="authorListItem">
            <h3><a class="authorItemName" href="#article1">廖显东</a></h3>
          </li>
          <li class="authorListItem">
            <h3><a class="authorItemName" href="#article1">廖雪峰</a></h3>
          </li>
          <li class="authorListItem">
            <h3><a class="authorItemName" href="#article1">廖显东</a></h3>
          </li>
          <li class="authorListItem">
            <h3><a class="authorItemName" href="#article1">廖雪峰</a></h3>
          </li>
          <li class="authorListItem">
            <h3><a class="authorItemName" href="#article1">廖显东</a></h3>
          </li>
          <li class="authorListItem">
            <h3><a class="authorItemName" href="#article1">廖雪峰</a></h3>
          </li>
          <li class="authorListMore">...查看更多</li>
        </ol>

        <!-- 可以添加更多文章项 -->
      </div>

      <div class="articleRecommendedList">
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

        <!-- 可以添加更多文章项 -->
      </div>
    </section>

    <!-- <section id="notifications">
      <h2>通知</h2>
      通知内容可以在这里添加 
    </section>

    <section id="profile">
      <h2>个人中心</h2>
       个人中心内容可以在这里添加 
    </section> -->

    <footer>
      <p>&copy; 2023 个人博客系统</p>
    </footer>
  </body>
</template>
<script>
import axios from "axios";
export default {
  name: "Host_Page",
  data() {
    return {
      isLoggedIn: localStorage.getItem("isLoggedIn"),
      username: localStorage.getItem("username"),
      imageSrc: "",

      articleId: "0",
      // articleList:[Head:0,Article_ID:0]
      articleList: "",
    };
  },
  created() {
    this.fetchAvatar(); // 在页面加载时调用fetchAvatar方法
    this.fetchArticleHead(); // 在页面加载时调用fetchArticle方法
    this.fetchArticleList();
    this.getLoggedIn();
  },
  methods: {
    getLoggedIn(){
      this.isLoggedIn=localStorage.getItem("isLoggedIn");
      console.log( localStorage.getItem("isLoggedIn"));
      

    },
    
    logout() {
      this.isLoggedIn = false;
      localStorage.setItem('isLoggedIn',this.isLoggedIn)
      console.log( localStorage.getItem("isLoggedIn"));
      

    },
    fetchAvatar() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/avatar/${this.username}`) // 使用get请求获取头像图片文件
        .then(async (response) => {
          console.log(response.data);
          this.imageSrc = "data:image/png;base64," + response.data.data; // 更新imageSrc以显示头像
        })
        .catch((error) => {
          console.error(error);
        });
    },
    fetchArticleHead() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/head/${this.articleId}`) // 使用get请求获取文章标题
        .then(async (response) => {
          console.log(response.data);
          this.articleHead = response.data.data; // 显示文章内容
        })
        .catch((error) => {
          console.error(error);
        });
    },
    fetchArticleList() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/article-list/${this.articleId}`) // 使用get请求获取文章标题
        .then(async (response) => {
          console.log(response.data);
          this.articleList = response.data; // 显示文章内容

          // this.articleList = make([]interface{},len(response.data.articleList))
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>
<style lang="less">
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
  background-color: #f2f2f2;
}
.navgationbarItemAvator {
  height: auto;
  .imgAvator {
    height: 100%;
  }
  margin-right: 30px;
}
.navgationbarItemLog {
  width: 20%;
  border: 1px solid rgb(128, 128, 128);
  float: left;
  text-align: center;
  line-height: 100%;
  margin-right: 20px;
}
.navgationbarItemRegister {
  width: 20%;
  border: 1px solid rgb(128, 128, 128);
  float: left;
  text-align: center;
  line-height: 100%;
  margin-right: 70px;
}
header {
  background-color: #333;
  color: #fff;
  padding: 10px;
  text-align: center;
  imgProfile {
  }
}

nav {
  background-color: #555;
  overflow: hidden;
}

nav a {
  float: left;
  display: block;
  color: #fff;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
}

nav a:hover {
  background-color: #ddd;
  color: #333;
}

section {
  padding: 20px;
  display: flex;
  justify-content: space-between;
}

.article {
  background-color: #fff;
  margin-bottom: 10px;
  padding: 15px;
  border-radius: 8px;
}

.article a {
  text-decoration: none;
  color: #333;
  font-weight: bold;
}

.article a:hover {
  color: #555;
}

.navgationbar {
  background-color: #f3f3f3;
  list-style: none;
  display: flex;
  justify-content: flex-start;
  width: 100%;
  height: 50px;
  padding: 0;

  .imgLogo {
    margin-right: 25px;
    margin-left: 10px;
  }

  .navgationbarItemHome {
    width: 5%;
    border: 1px solid rgb(128, 128, 128);
    float: left;
    text-align: center;
    line-height: 100%;

    .imgHome {
      height: 40%;
    }
  }
  .navgationbarItemNotifications {
    width: 5%;
    border: 1px solid rgb(128, 128, 128);
    float: left;
    text-align: center;

    line-height: 100%;
    .imgNotifications {
      height: 40%;
    }
  }
  .navgationbarItemProfile {
    width: 5%;
    border: 1px solid rgb(128, 128, 128);
    float: left;
    text-align: center;
    line-height: 100%;
    margin-right: 300px;
    .imgProfile {
      height: 40%;
    }
  }
  .navgationbarItemSearch {
    position: absolute;
    top: 0;
    left: 36%;
    height: 50px;
    .imgSearch {
      width: auto;
      height: 70%;
      margin-right: 20px;
    }

    display: flex;
    justify-content: flex-start;
    input {
      border-radius: 20%;
    }
  }
  .navgationbarItemCreaterCenter {
    background-color: red;
    text-align: center;
    width: 100px;
    margin-right: 30px;
    height: 50px;

    position: absolute;
    top: 0;
    left: 65%;
    height: 50px;
  }
  .loggedIn {
    display: flex;
    justify-content: center;

    .navgationbarItemAvator {
      .imgAvator {
        width: 50px;
        height: 50px;
        right: 400px;
        position: absolute;
        top: 0;
        left: 75%;
        height: 50px;
      }
    }
    .navgationbarItemLogOut {
      position: absolute;
      top: 0;
      left: 85%;
      height: 50px;
    }
  }
  .nologgedIn {
    display: flex;
    justify-content: space-between;
    .navgationbarItemRight {
      display: flex;
      width: 50px;
      height: 50px;
      position: absolute;
      top: 0;
      left: 85%;
      height: 50px;
      .navgationbarItemLog{

      }
    }
  }
}
.articleRecommendedList {
  width: 50%;
  height: 70%;
  background: #eee;
  border: 1px solid #ccc;
  position: absolute;
  /* display: flex;
      justify-content: center;
      align-items: center; */

  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  margin-left: -50px;
  margin-top: -50px;
}

.articleList {
  left: 5%;
  border: 1px solid #ccc;
  width: 20%;
}

.imgLogo {
  object-fit: contain;
}

.authorList {
  right: 5%;
  width: 20%;
  border: 1px solid #ccc;
}

.authorItem {
  background-color: #fff;
  margin-bottom: 10px;
  padding: 10px;
  border-radius: 8px;
}

footer {
  background-color: #333;
  color: #fff;
  text-align: center;
  padding: 10px;
  position: fixed;
  bottom: 0;
  width: 100%;
}

.articleItemTitle {
  height: 8px;
}

.articleListTiTle {
  width: auto;
}
</style>
