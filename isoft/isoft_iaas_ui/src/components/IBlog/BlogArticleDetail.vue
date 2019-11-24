<template>
  <div v-if="blog" style="background: #ffffff;margin: 10px;padding: 20px;min-height: 800px;">
    <h3>{{blog.blog_title}}</h3>
    <div style="border-bottom: 1px solid #f4f4f4;margin-top:20px;margin-bottom: 20px;">
      <Row style="margin: 10px;">
        <Col span="18">
          发布于:<Time :time="blog.created_time" style="color:red;"/>&nbsp;
          更新于:<Time :time="blog.last_updated_time" style="color:red;"/>&nbsp;
          作者:{{blog.author}}
        </Col>
        <Col span="3">阅读次数 {{blog.views}}</Col>
        <Col span="3">编辑次数 {{blog.edits}}
          <Button type="success" size="small" v-if="editable"
            @click="$router.push({ path: '/iblog/mine/blog_edit', query: { id: blog.id }})">编辑</Button>
        </Col>
      </Row>
    </div>
    <IShowMarkdown v-if="blog.content" :content="blog.content"/>
    <span v-if="blog.link_href">分享链接：<a :href="blog.link_href" target="_blank">{{blog.link_href}}</a></span>

    <hr>
    <!-- 评论模块 -->
    <IEasyComment :theme_pk="blog.id" theme_type="blog_theme_type" style="margin-top: 50px;"/>
  </div>
</template>

<script>
  import {ShowBlogArticleDetail} from "../../api"
  import IShowMarkdown from "../Common/markdown/IShowMarkdown"
  import IEasyComment from "../Comment/IEasyComment"
  import {CheckHasLogin,GetLoginUserName} from "../../tools"

  export default {
    name: "BlogArticleDetail",
    components:{IShowMarkdown,IEasyComment},
    data(){
      return {
        blog: null,
      }
    },
    methods:{
      refreshArticleDetail:async function () {
        const result = await ShowBlogArticleDetail(this.$route.query.blog_id);
        if(result.status=="SUCCESS"){
          this.blog = result.blog;
        }
      },
    },
    mounted:function () {
      this.refreshArticleDetail();
    },
    computed:{
      editable:function () {
        return CheckHasLogin() && this.blog != null && this.blog.author == GetLoginUserName();
      }
    }
  }
</script>

<style scoped>
</style>
