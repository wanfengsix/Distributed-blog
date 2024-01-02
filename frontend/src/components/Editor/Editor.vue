<template>
  <div class="navigationBar">
    <input
      type="text"
      style="width: 800px"
      v-model="headertext"
      placeholder="请输入文章标题..."
    />
    <div class="drafts" @click="getDraft">草稿箱</div>
    <div class="save" @click="saveArticle">保存</div>
    <div class="release" @click="releaseArticle">发布</div>
    <div v-if="isLoggedIn" class="loggedIn">
      <div class="navgationbarItemAvator">
        <img class="imgAvator" :src="imageSrc" alt="Image from backend" />
      </div>
    </div>
    <div v-else class="nologgedIn">
      <div class="navgationbarItemRight">
        <div class="navgationbarItemLog"> 
          <a href="../login">登录</a>
        </div>
        <div class="navgationbarItemRegister">
          <a href="regist">注册</a>
        </div>
      </div>
    </div>
  </div>
  <div>
    <div class="editor" ></div>
  </div>
</template>

<script>
import { ref, onMounted, toRefs, reactive } from "vue";
import Quill from "quill";
import "quill/dist/quill.core.css";
import "quill/dist/quill.snow.css";
import "quill/dist/quill.bubble.css";
import axios from "axios";
export default {
  data() {
    return {
      isLoggedIn: localStorage.getItem("isLoggedIn"),
      imageSrc: "",
      username: localStorage.getItem("username"),
      headertext: "",
      uid: "",
      
      contentTemp: "",
    };
  },
  setup() {
    const uid = ref("");
    const headertext = ref("");
    const contentTemp = ref("");

    const state = reactive({
      isLoggedIn: localStorage.getItem("isLoggedIn"),
      imageSrc: "",
      username: localStorage.getItem("username"),
      headertext: "",
      uid: "",
      contentTemp: "",
      
    });
    const refs = toRefs(state);

    const quill = ref(null);

    onMounted(() => {
      quill.value = new Quill(".editor", {
        theme: "snow",
        modules: {
          toolbar: [
            ["bold", "italic", "underline", "strike"],
            ["blockquote", "code-block"],
            [{ header: 1 }, { header: 2 }],
            [{ list: "ordered" }, { list: "bullet" }],
            [{ script: "sub" }, { script: "super" }],
            [{ indent: "-1" }, { indent: "+1" }],
            [{ direction: "rtl" }],
            [{ size: ["small", false, "large", "huge"] }],
            [{ header: [1, 2, 3, 4, 5, 6, false] }],
            [{ color: [] }, { background: [] }],
            [{ font: [] }],
            [{ align: [] }],
            ["clean"],
          ],
        },
      });
      quill.value.clipboard.dangerouslyPasteHTML(state.contentTemp);
    });

   
    

    const saveArticle = () => {
      if (quill.value) {
        const content = quill.value.root.innerHTML; // 或者使用 getText() 获取纯文本内容
        console.log("保存文章:", content);

        const instance = axios.create({
          withCredentials: true,
        });

        const data = {
          data: content,

          head: headertext.value,
        };
        console.log("----------", typeof data.uid); // 访问 uid

        instance
          .post(`http://127.0.0.1:8088/draft/${uid.value}`, data)
          .then((response) => {
            // 处理成功的响应
            console.log(response.data);
          })
          .catch((error) => {
            // 处理错误
            console.error(error);
          });

        // 在这里添加保存文章的逻辑，可以将 content 发送到后端保存
      }
    };

    const releaseArticle = () => {
      if (quill.value) {
        const content = quill.value.root.innerHTML; // 或者使用 getText() 获取纯文本内容
        console.log("发布文章:", content);

        const instance = axios.create({
          withCredentials: true,
        });
        const data = {
          data: content,

          head: headertext.value,
        };
        console.log("----------", typeof data.uid); // 访问 uid

        instance
          .post(`http://127.0.0.1:8088/write/${uid.value}`, data)
          .then((response) => {
            // 处理成功的响应
            console.log(response.data);
          })
          .catch((error) => {
            // 处理错误
            console.error(error);
          });

        // 在这里添加保存文章的逻辑，可以将 content 发送到后端保存
      }
    };

    //   const getDraft = () => {

    //     const instance = axios.create({
    //       withCredentials: true,
    //     });

    //     instance
    //       .get(`http://127.0.0.1:8088/draft/${uid.value}`) // 使用get请求获取commentList内容
    //       .then(async (response) => {
    //         console.log(response.data);
    //         content = response.data.data; // 显示评论列表内容
    //       })
    //       .catch((error) => {
    //         console.error(error);
    //       });

    //       // 在这里添加保存文章的逻辑，可以将 content 发送到后端保存

    // };
    return {
      refs,
      uid,
      headertext,
      contentTemp,
      quill,
      saveArticle,
      releaseArticle,
      // getDraft,
    };
  },

  methods: {
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
    getuid() {
      const instance = axios.create({
        withCredentials: true,
      });

      instance
        .get(`http://127.0.0.1:8088/userinfo/u_name/${this.username}`)
        .then(async (response) => {
          console.log(response.data);
          this.uid = response.data.data.uid.String; // 获取uid
        })
        .catch((error) => {
          console.error(error);
        });
    },

    getDraft() {
      const instance = axios.create({
        withCredentials: true,
      });

      instance
        .get(`http://127.0.0.1:8088/draft/${this.uid}`) // 使用get请求获取commentList内容
        .then(async (response) => {
          console.log(response.data);
          this.contentTemp = response.data.message; // 显示评论列表内容
        })
        .catch((error) => {
          console.error(error);
        });

      // 在这里添加保存文章的逻辑，可以将 content 发送到后端保存
    },
  },
  created() {
    this.fetchAvatar();
    this.getuid();
  },
};
</script>

<style lang="less" scoped>
.editor {
  height: 500px;
  width: 100%;
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 10px;
}

.navigationBar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #ccc;
  .drafts {
    padding: 8px 16px;
    font-size: 30px;
    cursor: pointer;
    border: 1px solid rgb(168, 207, 243);
  }
  .save {
    padding: 8px 16px;
    font-size: 30px;
    cursor: pointer;
    border: 1px solid #ccc;
    background-color: rgb(64, 149, 229);
    border-radius: 10%;
  }
  .release {
    padding: 8px 16px;
    font-size: 30px;
    cursor: pointer;
    border: 1px solid #ccc;
    background-color: rgb(64, 149, 229);
    border-radius: 10%;
  }
  .navgationbarItemAvator {
    height: auto;
    .imgAvator {
      height: 60px;
    }
    margin-right: 30px;
  }
}
</style>
