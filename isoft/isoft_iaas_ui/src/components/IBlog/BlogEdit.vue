<template>
  <div style="padding: 30px;">
    <Row>
      <Col span="4">
        <p>新建博客</p>
        <p>新建分享</p>
        <p>新建视频</p>
      </Col>
      <Col span="20">
        <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="80">
          <Row>
            <Col span="12">
              <FormItem label="文章标题" prop="blog_title">
                <Input v-model="formValidate.blog_title" placeholder="Enter blog title..."/>
              </FormItem>
            </Col>
            <Col span="12">
              <FormItem label="简短描述" prop="short_desc">
                <Input v-model="formValidate.short_desc" placeholder="Enter short_desc..."></Input>
              </FormItem>
            </Col>
          </Row>
          <Row>
            <Col span="12">
              <FormItem label="检索词条" prop="key_words">
                <Input v-model="formValidate.key_words" placeholder="Enter key_words..."></Input>
              </FormItem>
            </Col>
            <Col span="12">
              <FormItem label="文章分类" prop="catalog_id">
                <Select v-model="formValidate.catalog_id" filterable>
                  <Option v-for="mycatalog in mycatalogs" :value="mycatalog.id" :key="mycatalog.id">{{ mycatalog.catalog_name }}</Option>
                </Select>
              </FormItem>
            </Col>
          </Row>
          <FormItem label="文章内容" prop="content">
              <mavon-editor ref="md" v-model="formValidate.content" @imgAdd="$imgAdd"
                          :toolbars="toolbars" :ishljs = "true" style="z-index: 1;"/>
          </FormItem>
          <FormItem>
            <Button type="primary" @click="handleSubmit('formValidate')">Submit</Button>
            <Button style="margin-left: 8px" @click="handleReset('formValidate')">Cancel</Button>
          </FormItem>
        </Form>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {GetMyCatalogs,BlogEdit,ShowBlogDetail} from "../../api"
  import axios from 'axios'

  export default {
    name: "BlogEdit",
    data () {
      return {
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
        blog:null,
        // 我的所有文章分类
        mycatalogs:[],
        formValidate: {
          id:-1,
          blog_title: '',
          short_desc: '',
          key_words: '',
          catalog_id: -1,
          content:"",
        },
        ruleValidate: {
          blog_title: [
            { required: true, message: '文章标题不能为空', trigger: 'blur' }
          ],
          short_desc: [
            { required: true, message: '简短描述不能为空', trigger: 'blur' }
          ],
          key_words: [
            { required: true, message: '检索词条不能为空', trigger: 'blur' }
          ],
          // catalog_id: [
          //   { required: true, type: 'number', message: '文章分类不能为空', trigger: 'blur' }
          // ],
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
          url: '/api/iwork/fileUpload/fileUpload',
          method: 'post',
          data: formdata,
          headers: { 'Content-Type': 'multipart/form-data' },
        }).then((result) => {
          // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
          this.$refs.md.$img2Url(pos, result.data.fileServerPath);
        })
      },
      handleSubmit (name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await BlogEdit(_this.formValidate.id,_this.formValidate.blog_title,_this.formValidate.short_desc,
              _this.formValidate.key_words, _this.formValidate.catalog_id, _this.formValidate.content);
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              this.$router.push({ path: '/iblog/blog_list'});
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
      refreshBlogDetail:async function () {
        const result = await ShowBlogDetail(this.$route.query.blog_id);
        if(result.status=="SUCCESS"){
          this.blog = result.blog;
          this.formValidate.id = result.blog.id;
          this.formValidate.blog_title = result.blog.blog_title;
          this.formValidate.short_desc = result.blog.short_desc;
          this.formValidate.key_words = result.blog.key_words;
          this.formValidate.catalog_id = result.blog.catalog_id;
          this.formValidate.content = result.blog.content;
        }
      },
    },
    mounted:async function () {
      if(this.$route.query.blog_id != undefined && this.$route.query.blog_id != null){
        this.refreshBlogDetail();
      }

      const result = await GetMyCatalogs();
      if(result.status=="SUCCESS"){
        this.mycatalogs = result.catalogs;
      }
    }
  }
</script>

<style scoped>

</style>
