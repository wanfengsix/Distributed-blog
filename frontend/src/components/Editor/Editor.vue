<template>
  <div class="navigationBar">
    <input type="text" style="width: 800px"   v-model="headertext" placeholder="请输入文章标题..." />
    <div class="drafts">草稿箱</div>
    <div class="save" @click="saveArticle">保存</div>
    <div class="release">发布</div>
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
    <div class="editor"></div>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
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
    };
  },
  setup() {
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
            ["image","video"]//上传图片、上传视频
          ],
        },
      });
    });

    const saveArticle = () => {
      if (quill.value) {
        const content = quill.value.root.innerHTML; // 或者使用 getText() 获取纯文本内容
        console.log("保存文章:", content);
        const instance = axios.create({
        withCredentials: true,
      });
      instance
      .post("http://127.0.0.1:8088/likes", content)
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

    return { quill, saveArticle };
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
    fetchHead() {
      const instance = axios.create({
        withCredentials: true,
      });
      instance
      .post("http://127.0.0.1:8088/likes", this.headertext)
        .then((response) => {
          // 处理成功的响应
          console.log(response.data);
        })
        .catch((error) => {
          // 处理错误
          console.error(error);
        });

    },
  },
  created() {
    this.fetchAvatar();
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
