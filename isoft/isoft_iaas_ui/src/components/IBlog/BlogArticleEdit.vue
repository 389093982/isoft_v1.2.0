<template>
  <Row>
    <Col span="6">
      <CatalogList style="margin-top: 20px;"/>
    </Col>
    <Col span="18">
      <div style="padding: 30px;">
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
              <FormItem label="文章分类" prop="catalog_name">
                <Select v-model="formValidate.catalog_name" filterable>
                  <!-- 热门分类 -->
                  <Option v-for="(hotCatalogItem,index) in hotCatalogItems" :value="hotCatalogItem.title"
                          :key="'__hot__' + index + hotCatalogItem.title">热门分类： {{ hotCatalogItem.title }}</Option>
                  <!-- 我的分类 -->
                  <Option v-for="(mycatalog, index) in mycatalogs" :value="mycatalog.catalog_name"
                          :key="'__mine__' + index + mycatalog.catalog_name">我的分类：{{ mycatalog.catalog_name }}</Option>
                </Select>
              </FormItem>
            </Col>
          </Row>
          <FormItem label="文章内容" prop="content">
            <mavon-editor ref="md" v-model="formValidate.content" @imgAdd="$imgAdd"
                          :toolbars="toolbars" :ishljs = "true" style="z-index: 1;"/>
          </FormItem>
          <FormItem label="分享链接" prop="link_href">
            <Input v-model="formValidate.link_href" placeholder="请输入分享链接"></Input>
          </FormItem>
          <FormItem>
            <Button type="success" @click="handleSubmit('formValidate')">提交</Button>
            <Button type="error" v-if="formValidate.article_id > 0"
                    style="margin-left: 8px" @click="handleDelete('formValidate')">删除该条目</Button>
          </FormItem>
        </Form>
      </div>
    </Col>
  </Row>
</template>

<script>
  import {GetMyCatalogs,BlogArticleEdit,ArticleDelete,ShowBlogArticleDetail,FilterElementByPlacement} from "../../api"
  import axios from 'axios'
  import CatalogList from "./CatalogList";

  export default {
    name: "BlogArticleEdit",
    components: {CatalogList},
    props:{
      successEmit:{
        type:Boolean,
        default:false,
      },
      bookId:{
        type:Number,
        default: -1,
      }
    },
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
        hotCatalogItems:[],
        formValidate: {
          article_id:-1,
          blog_title: '',
          short_desc: '',
          key_words: '',
          catalog_name: '',
          content:"",
          link_href:"",
        },
        ruleValidate: {
          blog_title: [
            { required: true, message: '文章标题不能为空', trigger: 'blur' }
          ],
          // short_desc: [
          //   { required: true, message: '简短描述不能为空', trigger: 'blur' }
          // ],
          // key_words: [
          //   { required: true, message: '检索词条不能为空', trigger: 'blur' }
          // ],
          // catalog_name: [
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
          url: '/api/iwork/httpservice/fileUpload',
          method: 'post',
          data: formdata,
          headers: { 'Content-Type': 'multipart/form-data' },
        }).then((result) => {
          // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
          this.$refs.md.$img2Url(pos, result.data.fileServerPath);
        })
      },
      // createEmptyArticle:async function(book_id){
      //   const result = await BlogEdit(-1, book_id, '新建文章', '', '', '', '','');
      //   if(result.status == "SUCCESS"){
      //       this.$emit("successEmitFunc");
      //   }
      // },
      handleDelete: async function(name){
        if(this.formValidate.article_id > 0){
          const result = await ArticleDelete(this.formValidate.article_id);
          if(result.status == "SUCCESS"){
            this.$refs[name].resetFields();
            if(this.successEmit){
              this.$emit("successEmitFunc");
            }
          }
        }
      },
      handleSubmit: function(name) {
        var _this = this;
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await BlogArticleEdit(_this.formValidate.article_id, this.bookId,
              _this.formValidate.blog_title,_this.formValidate.short_desc,
              _this.formValidate.key_words, _this.formValidate.catalog_name,
              _this.formValidate.content,_this.formValidate.link_href);
            if(result.status == "SUCCESS"){
              _this.$Message.success('提交成功!');
              if(this.successEmit){
                this.$emit("successEmitFunc");
              }else{
                this.$router.push({ path: '/iblog/blog_list'});
              }
            }else{
              _this.$Message.error('提交失败!');
            }
          } else {
            _this.$Message.error('验证失败!');
          }
        })
      },
      refreshArticleDetail:async function (article_id) {
        var articleId = article_id > 0 ? article_id : this.$route.query.id;
        this.formValidate.article_id = articleId;
        const result = await ShowBlogArticleDetail(articleId);
        if(result.status=="SUCCESS"){
          if(result.blog != null){
            this.blog = result.blog;
            this.formValidate.article_id = result.blog.id;
            this.formValidate.blog_title = result.blog.blog_title;
            this.formValidate.short_desc = result.blog.short_desc;
            this.formValidate.key_words = result.blog.key_words;
            this.formValidate.catalog_name = result.blog.catalog_name;
            this.formValidate.content = result.blog.content;
            this.formValidate.link_href = result.blog.link_href;
          }
        }
      },
      refreshHotCatalogItems: async function () {
        const result = await FilterElementByPlacement(this.GLOBAL.placement_host_recommend_blog_tpyes);
        if(result.status == "SUCCESS"){
          this.hotCatalogItems = result.elements;
        }
      },
    },
    mounted:async function () {
      // 加载热门分类
      this.refreshHotCatalogItems();
      // 数据回显
      if(this.$route.query.id != undefined && this.$route.query.id != null){
        this.refreshArticleDetail();
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
