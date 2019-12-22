<template>
  <div>
    <Spin fix size="large" v-if="isLoading">
      <div class="isoft_loading"></div>
    </Spin>
    <Form ref="formValidate" :model="formValidate" :rules="ruleValidate">
      <FormItem prop="content">
        <mavon-editor ref="md" v-model="formValidate.content" @imgAdd="$imgAdd"
                      :toolbars="toolbars" :ishljs = "true" style="z-index: 1;min-height: 500px;"/>
      </FormItem>
      <FormItem>
        <Button type="success" @click="handleSubmit('formValidate')">提交</Button>
        <Button type="error" v-if="formValidate.article_id > 0"
                style="margin-left: 8px" @click="handleDelete('formValidate')">删除该条目</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  import {ShowBookArticleDetail,BookArticleEdit} from "../../api"

  export default {
    name: "BookArticleEdit",
    data(){
      return {
        isLoading:false,
        bookArticle:null,
        toolbars: {
          bold: true, // 粗体
          italic: true, // 斜体
          header: true, // 标题
          underline: true, // 下划线
          // mark: true, // 标记
          superscript: true, // 上角标
          quote: true, // 引用
          ol: true, // 有序列表
          link: true, // 链接
          imagelink: true, // 图片链接
          help: true, // 帮助
          code: true, // code
          subfield: true, // 是否需要分栏
          fullscreen: true, // 全屏编辑
          readmodel: true, // 沉浸式阅读
          undo: true, // 上一步
          trash: true, // 清空
          save: true, // 保存（触发events中的save事件）
          navigation: true // 导航目录
        },
        formValidate: {
          id:-1,
          book_catalog_id:-1,
          content:"",
        },
        ruleValidate: {
          content: [
            { required: true, message: '文章内容不能为空', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      // 绑定@imgAdd event
      // 参考 https://github.com/hinesboy/mavonEditor/blob/master/doc/cn/upload-images.md
      $imgAdd(pos, $file){
        // 第一步.将图片上传到服务器.
        var formdata = new FormData();
        formdata.append('file', $file);
        axios({
          url: '/api/iwork/httpservice/fileUpload',
          method: 'post',
          data: formdata,
          headers: { 'Content-Type': 'multipart/form-data' },
        }).then((result) => {
          // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
          this.$refs.md.$img2Url(pos, result.data.fileServerPath);
        })
      },
      refreshBookArticleDetail:async function (book_catalog_id) {
        this.isLoading = true;
        try {
          this.formValidate.book_catalog_id = book_catalog_id;
          const result = await ShowBookArticleDetail(book_catalog_id);
          if(result.status=="SUCCESS"){
            if(result.bookArticle != null){
              this.bookArticle = result.bookArticle;
              this.formValidate.id = result.bookArticle.id;
              this.formValidate.content = result.bookArticle.content;
            }else{
              this.bookArticle = null;
              this.formValidate.id = -1;
              this.formValidate.content = "";
            }
          }else{
            this.$Message.error("加载失败!");
          }
        } finally {
          this.isLoading = false;
        }
      },
      handleSubmit: function(name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await BookArticleEdit(_this.formValidate.id, _this.formValidate.book_catalog_id, _this.formValidate.content);
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              if(this.successEmit){
                this.$emit("successEmitFunc");
              }
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
          }
        })
      },
    }
  }
</script>

<style scoped>


</style>
