<template>
  <div>
    <div v-for="(comment_reply,index) in comment_replys" style="margin-bottom:5px;padding: 10px;border: 1px solid #e9e9e9;">
      <p>
        <router-link to="">
          <Avatar icon="ios-person" size="small" />
          {{comment_reply.created_by}}
        </router-link>
      </p>
      <p>
        <span v-if="comment_reply.reply_comment_type == 'comment'">
          回复
        </span>
        <span v-else>
          提问
        </span>
        <router-link to="">
          <Avatar icon="ios-person" size="small" />
          {{comment_reply.refer_user_name}}
        </router-link>
        :{{comment_reply.reply_content}}
        <span style="float: right;"><Time :time="comment_reply.created_time" :interval="1"/></span>
      </p>
      <p>
        <Row>
          <span style="float: right;">
            <a @click="toggleShow(index,comment_reply)" v-if="comment_reply.sub_reply_amount > 0">子评论数({{comment_reply.sub_reply_amount}})</a>
            <a v-if="comment_reply.depth < 2" href="javascript:;" @click="replyComment(comment_reply.id,comment_reply.created_by)">回复他/她</a>&nbsp;
            <a href="javascript:;">点赞</a>
          </span>
        </Row>
      </p>
      <!-- 递归,子评论区域 -->
      <CommentArea v-if="comment_reply.sub_reply_amount > 0 && comment_reply.toggleSubCommentShow"
         :parent_id="comment_reply.id" :comment_id="comment_id" :theme_type="theme_type" :key="comment_reply.id"/>
    </div>

    <!-- 顶级评论支持分页 -->
    <Page v-if="parent_id == 0" :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>

    <!-- 评论表单 -->
    <Modal
      v-model="showCommentForm"
      width="800"
      title="回复"
      :mask-closable="false">
      <CommentForm v-if="showCommentForm" :parent_id="_parent_id" :comment_id="comment_id" :theme_type="theme_type"
        :refer_user_name="_refer_user_name" @refreshCommentReply="refreshCommentReply"/>
    </Modal>
  </div>
</template>

<script>
  import {FilterCommentReply} from "../../../api/index"
  import CommentForm from "./CommentForm"

  export default {
    name: "CommentArea",
    // 评论清单
    props:{
      parent_id:{
        type:Number,
        default: -1,
      },
      comment_id:{
        type:Number,
        default:-1,
      },
      theme_type:{
        type:String,
        default:"",
      }
    },
    components:{CommentForm},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        comment_replys:[],
        showCommentForm:false,
        // 回复评论,两个参数分别是被评论id,被评论人
        _parent_id:0,
        _refer_user_name:"",
      }
    },
    methods:{
      toggleShow(index, comment_reply){
        comment_reply.toggleSubCommentShow = !(comment_reply.toggleSubCommentShow == null ? false : comment_reply.toggleSubCommentShow);
          this.$set(this.comment_replys, index, comment_reply);
      },
      handleChange(page){
        this.current_page = page;
        this.refreshCommentReply();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshCommentReply();
      },
      // 刷新当前父级评论对应的评论列表
      refreshCommentReply:async function(reply_comment_type){
        if(reply_comment_type == undefined){
          reply_comment_type = "all";
        }
        const result = await FilterCommentReply(this.comment_id, this.theme_type, this.parent_id, reply_comment_type, this.offset, this.current_page);
        if(result.status=="SUCCESS"){
          this.showCommentForm = false;
          this.comment_replys = result.comment_replys;
          if(result.paginator != null){
            this.total = result.paginator.totalcount;
          }
        }
      },
      // 回复评论,两个参数分别是被评论id,被评论人
      replyComment:function(id,refer_user_name){
        this._parent_id = id;
        this._refer_user_name = refer_user_name;
        this.showCommentForm = true;
      },
    },
    mounted:function () {
      this.refreshCommentReply();
    },
    watch:{
      "comment_id": "refreshCommentReply"      // 如果 comment_id 有变化,会再次执行该方法
    }
  }
</script>

<style scoped>
  a{
    color:red;
    margin-right: 10px;
  }
</style>
