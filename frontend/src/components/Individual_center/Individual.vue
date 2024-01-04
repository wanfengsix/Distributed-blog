<template>
  <head>
    <meta charset="utf-8" />
  </head>

  <body>
    <div class="navgationbar">
      <img class="imgLogo" src="../../../public/img/logo.jpg" />
      <div class="navgationbarItemHome">
        <img class="imgHome" src="../../../public/img/home.png" alt="" />
        <a href="/">首页</a>
      </div>
      <div class="navgationbarItemNotifications">
        <img
          class="imgNotifications"
          src="../../../public/img/notifications.png"
          alt=""
        />
        <a href="#notifications">通知</a>
      </div>
      <div class="navgationbarItemProfile">
        <img class="imgProfile" src="../../../public/img/profile.png" alt="" />
        <a :href="'individual/' + this.uid">个人中心</a>
      </div>
      <div class="navgationbarItemSearch">
        <img
          class="imgSearch"
          @click="getSearchList"
          src="../../../public/img/search.png"
        />
        <input
          type="text"
          style="width: 400px"
          v-model="searchQuery"
          @keyup.enter="getSearchList"
        />
      </div>
      <div class="navgationbarItemCreaterCenter">
        <a :href="'../editor/' + this.uid" style="color: white">创作者中心</a>
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

    <div class="centen">
      <div class="main">
        <div class="centen1">
          <div class="left">
            <span class="aw1">通知</span>
            
              <div class="noticecommand">  
                <!-- 通过循环生成通知列表 -->  
                <div  v-if="shouldShowNoticeList()"
                  class="notice"  
                  v-for="notice in noticeList"  
                  :key="notice.Notice_ID.String"  
                  :style="{ position: 'relative' }"  
                >  
                <a >{{ notice.U_name.String }}给{{ notice.Author_name.String }}的文章“{{ notice.Head.String }}”{{ notice.Type.String }} </a>  
                <button
                class="read-button"
                
                @click="readnotice(notice.Notice_ID.String)"
              >
                已读
              </button>
                  
                </div>  
              </div>  
              <div class="clear"></div>  
            </div>
          

          <div class="zhong">
            <div class="shang">
              <div class="zuo" v-if="avatarPreview">
                <img :src="avatarPreview" alt="Avatar Preview" />
              </div>
              <div class="zuo" v-else>
                <img :src="imageSrc" />
              </div>
              <div class="zhong">
                <h3>{{ u_name }}</h3>
                <h3>LV.{{ level }}</h3>
                <span>个性签名</span>
                <input
                  v-model="signature"
                  v-if="isSetting"
                  type="text"
                  placeholder="请输入个性签名"
                />
                <p v-else>{{ signature }}</p>
              </div>
              <div class="you">
                <div>
                  <input
                    type="file"
                    ref="fileInput"
                    @change="handleFileChange"
                    style="width: 100px; height: 50px; font-size: 15px"
                  />
                  <button
                    @click="uploadAvatar"
                    style="width: 100px; height: 30px; font-size: 15px"
                  >
                    上传头像
                  </button>
                </div>
                <button
                  :disabled="!isValidUser()"
                  v-if="!isSetting"
                  @click="isSetting = true"
                  style="width: 100px; height: 30px; font-size: 15px"
                >
                  设置个性签名
                </button>
                <button
                  :disabled="!isValidUser()"
                  v-else
                  @click="sendSignature()"
                >
                  完成
                </button>
              </div>

              <div class="clear"></div>
            </div>

            <div class="xia">
              <div class="articleRecommendedLis">
                <!-- 通过循环生成文章列表 -->
                <div
                  class="article"
                  v-for="(item, index) in articleList.article_list"
                  :key="index"
                  :style="{ position: 'relative' }"
                >
                  <a :href="getLink(item) " style="font-size: 25px;">{{ item.head }}</a>
                  <p style="font-size: 20px;">文章摘要或内容简介...</p>
                  <!-- 添加删除按钮 -->
                  <button
                    class="delete-button"
                    :disabled="!isValidUser()"
                    @click="deleteArticle(item.article_id)"
                  >
                    删除
                  </button>

                  <div class="isVisible">
                    <div class="isVisibleText">文章允许可见或不可见</div>
                    <button class="visible-button" :disabled="!isValidUser()" @click="setVisible(item.article_id)">可见</button>  
                    <button class="invisible-button" :disabled="!isValidUser()" @click="setunVisible(item.article_id)">不可见</button> 
                  </div>
                </div>
              </div>
              <div class="clear"></div>
            </div>
          </div>

          <div class="right">
            <div class="shang">
              <h3>关注了 &nbsp;&nbsp;丨&nbsp;&nbsp; 关注者<br /></h3>
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
import axios, { Axios } from "axios";
export default {
  name: "Host_Page",
  data() {
    return {
      isLoggedIn: localStorage.getItem("isLoggedIn"),
      username: localStorage.getItem("username"),
      imageSrc: "",
      imageSrca: "",
      signature: "",
      isSetting: false,
      articleId: "0",
      noticeList: "",
      articleList: "",
      level: 0,
      following: 0,
      followed: 0,
      followersList: [], // 关注列表数据
      fansList: [], // 粉丝列表数据
      uid: this.$route.params.uid,
      u_name: "",
      isVisible: true,
      avatarFile: null,
      avatarPreview: null,
      avatarBinary: null,
      avatarBase64: null,
      vif:this.u_name === this.username || this.username === 'admin',
    };
  },
  created() {
    if (localStorage.getItem("isLoggedIn") != false) {
      this.fetchAvatar(); // 在页面加载时调用fetchAvatar方法

      //this.fetchArticleList();
      //this.getfollowing();
      //this.getfollowed();
      // this.getfollowers_list();
      //this.getfans_list();
      this.getu_name();
      this.NoticeList();

    }
    if (this.u_name === this.username || this.username === 'admin') {
      this.fetchArticleListb();
    }
    else{
      this.fetchArticleList();
    }
  },
  methods: {
    getBase64() {
      const reader = new FileReader();
      reader.onload = (e) => {
        this.avatarBase64 = e.target.result.split(",")[1];
      };
      reader.readAsDataURL(this.avatarFile);
    },
    convertBinaryToString(binaryData) {
      const decoder = new TextDecoder("utf-8"); // 使用 UTF-8 编码解析二进制数据
      return decoder.decode(binaryData);
    },
    getBinary() {
      const reader = new FileReader();
      reader.onload = (e) => {
        this.avatarBinary = new Uint8Array(e.target.result);
      };
      reader.readAsArrayBuffer(this.avatarFile);
    },
    handleFileChange() {
      const fileInput = this.$refs.fileInput;
      if (fileInput.files.length > 0) {
        this.avatarFile = fileInput.files[0];
        this.previewAvatar();
      }
    },
    previewAvatar() {
      const reader = new FileReader();
      reader.onload = (e) => {
        this.avatarPreview = e.target.result;
        this.getBinary();
        this.getBase64();
      };
      reader.readAsDataURL(this.avatarFile);
    },
    uploadAvatar() {
      // 在这里实现上传头像的逻辑，使用 this.avatarFile
      // 你可以使用 FormData 将文件上传到服务器
      // 例如，可以使用 axios 或 Fetch API 发送 POST 请求
      // 示例：axios.post('/upload/avatar', formData)
      const data = {
        name: this.u_name,
        resource_type: "avatar",
        post_data: this.avatarBase64,
      };
      const instance = axios.create({
        withCredentials: true,
      });
      console.log("post头像", this.avatarBase64);
      console.log("post头像", typeof this.avatarBase64);

      //将signature发送到后端
      instance
        .post(`http://127.0.0.1:8088/user/avatar/${this.u_name}`, data)
        .then((response) => {
          // 处理成功的响应
          console.log("post成功base", response.data);
          this.$forceUpdate(); // 强制Vue组件重新渲染
        })
        .catch((error) => {
          // 处理错误
          console.error(error);
        });
    },
    updateImage() {
      // 在这里更新 imageSrc 的值
      this.imageStyle = {
        width: "10px",
        borderRadius: "75%",
      };
    },
    isValidUser() {
      return this.u_name === this.username || this.username === "admin";
    },
    getLink(item) {
      if (this.u_name === "username" || this.username === "admin") {
        return `../editor/${item.article_id}`; // 跳转到editor路由
      } else {
        return `../article/${item.article_id}`; // 跳转到article路由
      }
    },
    shouldShowNoticeList() {
      // 判断是否显示通知列表
      return this.u_name === this.username || this.username === "admin";
    },
    sendSignature() {
      const data = {
        name: this.u_name,
        resource_type: "signature",
        post_data: this.signature,
      };
      const instance = axios.create({
        withCredentials: true,
      });
      //将signature发送到后端
      instance
        .post(`http://127.0.0.1:8088/user/signature/${this.u_name}`, data)
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
      this.isLoggedIn = false;
      localStorage.setItem("isLoggedIn", false);
      // 在这里编写退出登录的逻辑，比如清除本地存储的用户信息等。
      // 这里只是一个示例，具体的实现取决于您的应用需求。
    },
    fetchAvatar() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/avatar/${this.username}`) // 使用get请求获取头像图片文件
        .then(async (response) => {
          this.imageSrc = "data:img/png;base64," + response.data.data; // 更新imageSrc以显示头像
        })
        .catch((error) => {
          console.error(error);
        });
    },
    fetchAvatara() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/avatar/${this.username}`) // 使用get请求获取头像图片文件
        .then(async (response) => {
          this.imageSrc = "data:img/png;base64," + response.data.data; // 更新imageSrc以显示头像
        })
        .catch((error) => {
          console.error(error);
        });
    },

    getSignature() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/signature/${this.u_name}`) //
        .then(async (response) => {
          this.signature = response.data.data; // 更新
        })
        .catch((error) => {
          console.error(error);
        });
    },

    // getfollowing(){
    // const instance = axios.create({
    //       withCredentials: true,
    //     });

    //     instance.get(`http://127.0.0.1:8088/userinfo/u_name/${this.username}`)
    //       .then(async response => {
    //         console.log(response.data)
    //         this.following = response.data.data.following; // 更新关注者数量
    // 	  this.followed=response.data.data.followed
    //       })
    //       .catch(error => {
    //         console.error(error);
    //       });
    // },
    getfollowed() {
      const instance = axios.create({
        withCredentials: true,
      });
      console.log(this.u_name);
      instance
        .get(`http://127.0.0.1:8088/userinfo/u_name/${this.u_name}`)
        .then(async (response) => {
          console.log(response.data);
          this.followed = response.data.data.followed; // 更新粉丝数量
        })
        .catch((error) => {
          console.error(error);
        });
    },

    handleIsvisible() {
      this.isVisible = !this.isVisible;
    },

    fetchArticleList() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/article-list-individual/${this.uid}`) // 使用get请求获取文章标题
        .then(async (response) => {
          this.articleList = response.data; // 显示文章内容
          console.log(response.data);

          // this.articleList = make([]interface{},len(response.data.articleList))
        })
        .catch((error) => {
          console.error(error);
        });
    },
    fetchArticleListb() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/article-list-self/${this.uid}`) // 使用get请求获取文章标题
        .then(async (response) => {
          this.articleList = response.data; // 显示文章内容
          console.log(response.data);

          // this.articleList = make([]interface{},len(response.data.articleList))
        })
        .catch((error) => {
          console.error(error);
        });
    },
    NoticeList() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/notice-list/${this.uid}`) 
        .then(async (response) => {
          this.noticeList = response.data.noticelistdata; // 显示通知列表内容
          console.log(response.data);

          // this.articleList = make([]interface{},len(response.data.articleList))
        })
        .catch((error) => {
          console.error(error);
        });
    },
    readnotice(Notice_ID){
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/read-notice/${Notice_ID}`) // 使用get请求
        .then(async (response) => {
        
        })
        .catch((error) => {
          console.error(error);
        });

    },
    deleteArticle(article_id){
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/delete-article/${article_id}`) // 使用get请求
        .then(async (response) => {
          this.fetchArticleList();
        })
        .catch((error) => {
          console.error(error);
        });
    },
    setVisible(article_id){
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/article-visible/${article_id}`) // 使用get请求
        .then(async (response) => {
        
        })
        .catch((error) => {
          console.error(error);
        });

    },
    setunVisible(article_id){
      const instance = axios.create({
        withCredentials: true,
      });
      instance
        .get(`http://127.0.0.1:8088/user/article-unvisible/${article_id}`) // 使用get请求
        .then(async (response) => {
       
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
        .get(`http://127.0.0.1:8088/userinfo/following-list/${this.u_name}`) // 使用get请求获取关注列表
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
        .get(`http://127.0.0.1:8088/userinfo/followed-list/${this.u_name}`) // 使用get请求获取粉丝列表
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
          this.followed = response.data.data.followed;
          this.following = response.data.data.following;
          this.level = response.data.data.level;
          this.getSignature();
          this.fetchAvatara();
          this.getfollowers_list();
          this.getfans_list();
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>
<style lang="less">
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
      .navgationbarItemLog {
      }
    }
  }
}
/*通用类*/
* {
  margin: 0;
  padding: 0;
}

body {
  min-width: 1400px;
  margin: 0 auto;
  font-size: 14px;
  background-color: #f5e9d9;
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

.lists-container {
  display: flex; /* 使用Flex布局 */
  justify-content: space-between; /* 使两个列表水平排列，并在它们之间添加间距 */
}

.list {
  width: 45%; /* 控制每个列表的宽度 */
}
.centen {
  background-color: #f5e9d9;
  margin-top: 50px;
}

.centen1 .left {
  width: 20%;
  float: left;
}
.centen1 .left .notificationItems {
  background-color: #ffffff;
  padding: 20px 0;
  border: 1px solid #ccc;
  border-radius: 10px;
  margin-bottom: 20px;
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
  color: #a0404e;
  border: 1px solid #a0404e;
  display: block;
  text-align: center;
  background-color: #ffffff;
  border-radius: 15px;
}

.centen1 .left p {
  background-color: #ffffff;
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
  background-color: #ffffff;
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
  color: #4e93de;
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
  border: 1px solid #abccdd;
  color: #abccdd;
  text-align: center;
  font-size: 20px;
  font-weight: bold;
}

.centen1 .zhong .xia {
  background-color: #ffffff;
  margin-top: 30px;
  padding-bottom: 100px;
  .articleRecommendedLis {
    .delete-button {
      position: absolute;
      right: 0;
      top: 0;
      font-size: 30px;
      background-color: rgb(229, 229, 229);
      color: #fff;

      border-radius: 10%;
      outline: none;
      cursor: pointer;
    }
    .isVisibleText {  
      font-size: 13px;  
    }
    
    .visible-button, .invisible-button {  
      width: 50px;  
      height: 30px; 
      font-size: 15px;  
    }
  }
}
.centen1 .zhong .xia ul li {
  padding: 15px;
  box-sizing: border-box;
}
.centen1 .zhong .xia ul li .yi {
  width: 50%;
  float: left;
}
.centen1 .zhong .xia ul li .yi h3 {
  font-size: 14px;
  font-weight: bold;
}
.centen1 .zhong .xia ul li .yi p {
  font-size: 12px;
  color: #515151;
  margin: 10px 0;
}
.centen1 .zhong .xia ul li .yi span {
  display: inline-block;
  width: 45px;
}

.centen1 .zhong .xia ul li .er {
  width: 35%;
  float: left;
  margin: 0 15px;
}
.centen1 .zhong .xia ul li .er img {
  width: 100%;
}
.centen1 .zhong .xia ul li .san {
  width: 8%;
  float: left;
}
.centen1 .zhong .xia ul li .san a {
  display: block;
  margin-top: 20px;
}

.centen1 .right {
  width: 20%;
  float: right;
}

.centen1 .right .shang {
  background-color: #ffffff;
  padding: 20px 0;
}
.centen1 .right .shang h3 {
  text-align: center;
  font-size: 22px;
}
.centen1 .right .shang span {
  font-size: 22px;
  width: 48%;
  text-align: center;
  display: inline-block;
}

.centen1 .right .xia {
  text-align: center;
  font-size: 22px;
  margin: 30px 0;
}

.centen1 .right .xia .bg1 {
  background-color: #ffffff;
  padding: 15px;
  margin-top: 20px;
  height: 800px;
}
.centen1 .right .xia .bg1 ul li {
  width: 100%;
  margin-top: 15px;
}
.centen1 .right .xia .bg1 ul li a {
  display: block;
  background: url /*(../images/grzx.png)*/ no-repeat left;
  padding-left: 20px;
  background-size: 19px;
  text-align: left;
  padding-left: 30px;
}
</style>
