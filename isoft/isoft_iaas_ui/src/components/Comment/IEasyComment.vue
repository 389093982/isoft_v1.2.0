<template>
<div>
  <div>
    <CommentForm :parent_id="parent_id" :theme_pk="theme_pk" :theme_type="theme_type"
       :refer_user_name="refer_user_name" @refreshComment="refreshComment"/>
  </div>

  <div style="margin-top: 30px;">
    <div class="comment_title">
      <span style="text-align: right;">
        <a href="javascript:;" @click="refreshComment('all')">全部</a>
        <Divider type="vertical" />
        <a href="javascript:;" @click="refreshComment('comment')">评论</a>
        <Divider type="vertical" />
        <a href="javascript:;" @click="refreshComment('question')">提问</a>
      </span>
      <span style="float: right;">
        <a href="javascript:;">更多评论话题</a>
      </span>
    </div>
    <!-- 评论列表 -->
    <CommentArea ref="commentArea" v-if="theme_pk > 0" :parent_id="0" :theme_pk="theme_pk" :theme_type="theme_type"></CommentArea>
  </div>
</div>
</template>

<script>
  import CommentArea from "./CommentArea"
  import CommentForm from "./CommentForm"

  export default {
    name: "CourseComment",
    components:{CommentForm,CommentArea},
    // 当前评论的课程
    props:{
      theme_type:{
        type:String,
        default:"",
      },
      theme_pk:{
        type:Number,
        default:-1,
      },
      referUserName:{
        type:String,
        default:'被评论人',
      }
    },
    data(){
      return {
        // 父评论 id
        parent_id:0,
        // 提交评论内容
        submit_comment:"",
        // 被评论人
        refer_user_name: this.referUserName,
      }
    },
    methods:{
      // 重新刷新评论列表
      refreshComment (comment_type) {
        // 调用子组件的刷新方法
        this.$refs.commentArea.refreshComment(comment_type);
      },
    },
  }
</script>

<style scoped lang="stylus" rel="stylesheet/stylus">
  .comment_title
    a
      color: #333
      text-decoration: none
      font-weight: 600
    a:hover
      color:red
</style>
