<template>
  <div id="editor"></div>
  <!-- <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>文章编辑器</title>
  </head>

  <body>
    <quill-editor v-model="content" ref="Editor"></quill-editor>
    <div class="navigationBar">
      <input type="text" style="width: 800px" placeholder="请输入文章标题..." />
      <div class="drafts">草稿箱</div>
      <div class="save" onclick="saveArticle">保存</div>
      <div class="release">发布</div>
      <div v-if="isLoggedIn" class="loggedIn">
        <div class="navgationbarItemAvator">
          <img class="imgAvator" :src="imageSrc" alt="Image from backend" />
        </div>
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
    <div class="menuBar">
      <div class="head">
        <img
          src="../../../public/img/head.png"
          @click="makeBold"
          style="width: 20px; height: 20px"
          alt="标题"
        />
      </div>
      <div class="bold">
        <img
          src="../../../public/img/bold.png"
          style="width: 20px; height: 20px"
          alt="粗体"
        />
      </div>
      <div class="italics">
        <img
          src="../../../public/img/italics.png"
          style="width: 20px; height: 20px"
          alt="斜体"
        />
      </div>
      <div class="quote">
        <img
          src="../../../public/img/quote.png"
          style="width: 20px; height: 20px"
          alt="引用"
        />
      </div>
      <div class="link">
        <img
          src="../../../public/img/link.png"
          style="width: 20px; height: 20px"
          alt="链接"
        />
      </div>
      <div class="picture">
        <img
          src="../../../public/img/picture.png"
          style="width: 20px; height: 20px"
          alt="图片"
        />
      </div>
      <div class="code">
        <img
          src="../../../public/img/code.png"
          style="width: 20px; height: 20px"
          alt="代码段"
        />
      </div>
    </div>

    <div class="contentArea">
      <div class="editArea" contenteditable="true"></div>
      <div class="displayArea"></div>
    </div>
  </body> -->
</template>

<script>
import { quillEditor } from "vue-quill-editor";
import { ref, onMounted } from "vue";
import Quill from "quill";
import "quill/dist/quill.core.css";
import "quill/dist/quill.snow.css";
import "quill/dist/quill.bubble.css";

export default {
  components: {
    quillEditor,
  },
  // setup() {
  //   const content = ref("");
  //   const makeBold = () => {
  //     Editor.value.format("bold", true);
  //   };
  //   const makeItalic = () => {
  //     Editor.value.format("italic", true);
  //   };

  //   const insertQuote = () => {
  //     Editor.value.format("blockquote", true);
  //   };

  //   const insertLink = () => {
  //     const url = prompt("请输入链接地址:");
  //     Editor.value.format("link", url);
  //   };

  //   const insertImage = () => {
  //     const url = prompt("请输入图片链接地址:");
  //     Editor.value.insertEmbed(0, "image", url);
  //   };

  //   const insertCode = () => {
  //     Editor.value.insertText(0, "```代码片段```");
  //   };

  //   return {
  //     content,
  //     makeBold,
  //     makeItalic,
  //     insertQuote,
  //     insertLink,
  //     insertImage,
  //     insertCode,
  //     myQuillEditor,
  //   };
  // },
  setup() {
    const quill = ref(null);

    onMounted(() => {
      quill.value = new Quill("#editor", {
        theme: "snow",
      });
    });

    return { quill };
  },
  data() {
    return {
      content: "", // 文章内容
      isLoggedIn: localStorage.getItem("isLoggedIn"),
      editorOption: {
        modules: {
          toolbar: [
            // 加粗 斜体 下划线 删除线 -----['bold', 'italic', 'underline', 'strike']
            ["bold", "italic", "underline", "strike"], // 引用  代码块-----['blockquote', 'code-block']
            ["blockquote", "code-block"], // 1、2 级标题-----[{ header: 1 }, { header: 2 }]
            [{ header: 1 }, { header: 2 }], // 有序、无序列表-----[{ list: 'ordered' }, { list: 'bullet' }]
            [{ list: "ordered" }, { list: "bullet" }], // 上标/下标-----[{ script: 'sub' }, { script: 'super' }]
            [{ script: "sub" }, { script: "super" }], // 缩进-----[{ indent: '-1' }, { indent: '+1' }]
            [{ indent: "-1" }, { indent: "+1" }], // 文本方向-----[{'direction': 'rtl'}]
            [{ direction: "rtl" }], // 字体大小-----[{ size: ['small', false, 'large', 'huge'] }]
            [{ size: ["small", false, "large", "huge"] }], // 标题-----[{ header: [1, 2, 3, 4, 5, 6, false] }]
            [{ header: [1, 2, 3, 4, 5, 6, false] }], // 字体颜色、字体背景颜色-----[{ color: [] }, { background: [] }]
            [{ color: [] }, { background: [] }], // 字体种类-----[{ font: [] }]
            [{ font: [] }], // 对齐方式-----[{ align: [] }]
            [{ align: [] }], // 清除文本格式-----['clean']
            ["clean"], // 链接、图片、视频-----['link', 'image', 'video']
            ["image"],
          ],
        },
        theme: "snow",
      },
    };
  },
  methods: {
    saveArticle() {
      var content = document.getElementById("editor").innerHTML;
      console.log("保存文章:", content);
      // 在这里添加保存文章的逻辑，可以将 content 发送到后端保存
    },
  },
};
</script>

<style lang="less" scoped>
/* 在这里添加样式，根据需要自定义编辑器的外观 */
body {
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
  }
  .menuBar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    border-bottom: 1px solid #ccc;
    .head {
      width: 20px;
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
    .bold {
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
    .italics {
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
    .quote {
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
    .link {
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
    .picture {
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
    .code {
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
  }

  .contentArea {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    border-bottom: 1px solid #ccc;
    .editArea {
      width: 850px;
      height: 800px;
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
    .displayArea {
      width: 850px;
      height: 800px;
      padding: 8px 16px;
      font-size: 30px;
      cursor: pointer;
      border: 1px solid rgb(168, 207, 243);
    }
  }

  font-family: "Arial", sans-serif;
  margin: 20px;
  .ditor {
    border: 1px solid #ccc;
    height: 300px;
    padding: 10px;
    box-sizing: border-box;
  }

  .saveButton {
    margin-top: 10px;
    padding: 8px 16px;
    font-size: 16px;
    cursor: pointer;
  }
}
</style>
