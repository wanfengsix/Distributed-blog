<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div id="app">
    <div class="navgationbar">
      <img class="imgLogo" src="img/logo.jpg" />
      <div class="navgationbarItemHome">
        <img class="imgHome" src="img/home.png" alt="" />
        <router-link to="/home">首页</router-link>
      </div>
      <div class="navgationbarItemNotifications">
        <img class="imgNotifications" src="img/notifications.png" alt="" />
        <router-link to="/notifications">通知</router-link>
      </div>
      <div class="navgationbarItemProfile">
        <img class="imgProfile" src="img/profile.png" alt="" />
        <router-link to="/profile">个人中心</router-link>
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
          <img class="imgAvator" src="img/temp.jpg" alt="Image from backend" />
        </div>
        <div class="navgationbarItemLog">
          <button
            @click="logout"
            style="width: 200px; height: 50px; font-size: 30px"
          >
            退出
          </button>
        </div>
      </div>
      <div v-else class="nologgedIn">
        <div class="navgationbarItemLog">
          <a href="login">登录</a>
        </div>
        <div class="navgationbarItemRegister">
          <a href="regist">注册</a>
        </div>
      </div>
    </div>

    <div class="Box">
      <div class="leftBox">
        <div class="like">
          <img class="imgLike" src="img/like.png" alt="" />
        </div>

        <div class="collect">
          <img class="imgCollect" src="img/collect.png" alt="" />
        </div>
      </div>

      <div class="middleBox">
        <div class="article">
          <article>
            <h1>文章标题</h1>
            <div v-html="articleContent"></div>
          </article>
        </div>

        <div class="commentArea">
          <div class="topComment">
            <textarea
              class="textInput"
              v-model="inputText"
              @input="handleInput"
              placeholder="在这里输入文本"
              style="width: 100px; height: 50px;display:inline-block;vertical-align: top;"
            ></textarea>
            <button
              style="width: 100px; height: 50px; font-size: 20px;display:inline-block;vertical-align: top;"
              class="sendComments"
              @click="showCommentsFangfa"
            >
              发送评论
            </button>
          </div>

          <div v-if="showComments">
            <div
              v-for="(text, index) in inputHistory"
              :key="index"
              class="displayText"
            >
              {{ text }}
            </div>
          </div>
        </div>
      </div>

      <div class="rightBox">
        <div class="personalInfo">
          <img class="imgIcon" src="img/personalIcon.jpg" alt="" />

          <div class="rightPersonalInfo">
            <p class="personalName">用户1772145090587</p>

            <img class="imgGrade" src="img/grade1.png" alt="" />
          </div>
        </div>
        <div class="articleInfo">
          <div class="numbers">
            <p>5</p>
            <p>文章</p>
          </div>
          <div class="reads">
            <p>5</p>
            <p>阅读</p>
          </div>
          <div class="fans">
            <p>5</p>
            <p>粉丝</p>
          </div>
        </div>
        <div class="follow">
          <button class="followButton" @click="toggleFollow">
            {{ buttonText }}
          </button>
        </div>
      </div>
    </div>

    <footer>
      <p>&copy; 2023 个人博客系统</p>
    </footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      articleContent: "",
      isFollowed: false,
      inputText: "",
      inputHistory: [],
      showComments: false,
      isLoggedIn: localStorage.getItem("isLoggedIn"),
    };
  },
  computed: {
    buttonText() {
      return this.isFollowed ? "已关注" : "关注";
    },
  },
  methods: {
    logout() {
      this.isLoggedIn = false;
    },
    showCommentsFangfa() {
      this.showComments = true;
      this.inputHistory.unshift(this.inputText);
      this.inputText = "";
    },
    handleInput() {
      // 将当前输入保存到历史记录中
    },
    handleFileChange(event) {
      const file = event.target.files[0];

      if (file) {
        const reader = new FileReader();

        reader.onload = () => {
          // 将文本内容绑定到Vue数据
          this.articleContent = reader.result;
          console.log(this.articleContent);
        };

        // 以文本格式读取文件
        reader.readAsText(file);
      }
    },
    toggleFollow() {
      this.isFollowed = !this.isFollowed;
    },
    loadDefaultFile() {
      // 读取默认文件，假设默认文件名为default.txt
      const defaultFilePath = "default.txt";

      // 使用XMLHttpRequest从同级目录下读取默认文件
      const xhr = new XMLHttpRequest();
      xhr.onload = () => {
        if (xhr.status === 200) {
          this.articleContent = xhr.responseText;
        } else {
          console.error("Failed to load the default file.");
        }
      };
      xhr.open("GET", defaultFilePath, true);
      xhr.send();
    },
  },
  created() {
    // 在实例创建时触发loadDefaultFile()函数
    this.loadDefaultFile();
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

  imgLogo {
    margin-right: 25px;
    margin-left: 10px;
  }

  .navgationbarItemHome {
    width: 5%;
    border: 1px solid rgb(128, 128, 128);
    float: left;
    text-align: center;
    line-height: 100%;

    imgHome {
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
    .imgSearch {
      width: auto;
      height: 70%;
      margin-right: 20px;
    }
    margin-right: 200px;

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

    imgProfile {
    }
  }

  .loggedIn {
    display: flex;
    justify-content: space-between;

    .navgationbarItemAvator {
      .imgAvator {
        width: 50px;
        height: 50px;
      }
    }
    .navgationbarItemLog {
      width: 50px;
      height: 50px;
    }
  }
  .nologgedIn {
    display: flex;
    justify-content: space-between;
  }
  // .navgationbarItemAvator {
  //   height: auto;
  //   .imgAvator {
  //     height: 100%;
  //   }
  //   margin-right: 30px;
  // }
  // .navgationbarItemLog {
  //   width: 5%;
  //   border: 1px solid rgb(128, 128, 128);
  //   float: left;
  //   text-align: center;
  //   line-height: 100%;
  //   margin-right: 20px;
  // }
  .navgationbarItemRegister {
    width: 5%;
    border: 1px solid rgb(128, 128, 128);
    float: left;
    text-align: center;
    line-height: 100%;
    margin-right: 70px;
  }
  .imgLogo {
    object-fit: contain;
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

footer {
  background-color: #333;
  color: #fff;
  text-align: center;
  padding: 10px;
  position: fixed;
  bottom: 0;
  width: 100%;
}

.Box {
  margin-top: 20px;
  // border-top: #333 1px solid;
  display: flex;
  justify-content: space-between;

  .leftBox {
    left: 0;
    display: flex;
    align-content: space-between;
    flex-flow: column;
    padding: 40px 40px;
    border: 1px solid gray;
    .imgLike {
      width: 50px;
      height: 50px;
    }
    .imgCollect {
      width: 50px;
      height: 50px;
    }
  }
  .middleBox {
    width: 700px;
    .comentArea {
      .topComment {
        display: flex;
        justify-content: space-between;
        align-items: space-between;
        .textInput {
          width: 300px;
          height: 200px;
          padding: 10px;
          margin-top: 20px;
        }
        .sendComments {
          width: 100px;
          height: 50px;
        }
      }

      .displayText {
        margin-top: 10px;
        border: 1px solid #ccc;
        padding: 10px;
      }
    }
  }
  .rightBox {
    width: 300px;
    border: 1px solid gray;
    height: 300px;
    .personalInfo {
      width: 250px;
      display: flex;
      .imgIcon {
        width: 80px;
        height: 80px;
        border-radius: 50%;
      }
      .rightPersonalInfo {
        display: flex;
        flex-flow: column;
        .imgGrade {
          width: 50px;
        }
      }
    }
    .articleInfo {
      display: flex;
      .numbers {
        width: 90px;
        height: 90px;
        border: 1px solid gray;
        margin: 10px;
        display: flex;
        flex-flow: column;
        justify-content: space-between;
        align-items: space-between;
      }
      .reads {
        width: 90px;
        height: 90px;
        border: 1px solid gray;
        margin: 10px;
        display: flex;
        flex-flow: column;
        justify-content: space-between;
        align-items: space-between;
      }
      .fans {
        width: 90px;
        height: 90px;
        border: 1px solid gray;
        margin: 10px;
        display: flex;
        flex-flow: column;
        justify-content: space-between;
        align-items: space-between;
      }
    }
    .follow {
      display: flex;
      justify-content: center;
      align-items: center;
      .followButton {
        background-color: blue;
        color: white;
        padding: 10px 20px;
        cursor: pointer;
        border: none;
        border-radius: 5px;
        font-size: 16px;
      }
      .followButton.followed {
        background-color: gray;
      }
    }
  }
}
</style>
