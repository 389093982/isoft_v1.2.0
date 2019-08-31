<template>
  <div>
    <div v-for="(comment,index) in comments" style="margin-bottom:5px;padding: 10px;border: 1px solid #e9e9e9;">
      <p>
        <Tag color="default" v-if="parent_id == 0"># 楼层{{(current_page - 1) * offset + index + 1}}</Tag>
        <router-link to="">
          <Avatar icon="ios-person" size="small" />
          {{comment.created_by}}
        </router-link>
      </p>
      <p>
        <span v-if="comment.comment_type == 'comment'">
          回复
        </span>
        <span v-else>
          提问
        </span>
        <router-link to="">
          <Avatar icon="ios-person" size="small" />
          {{comment.refer_user_name}}
        </router-link>
        :{{comment.content}}
        <span style="float: right;"><Time :time="comment.created_time" :interval="1"/></span>
      </p>
      <p>
        <Row>
          <span style="float: right;">
            <a @click="toggleShow(index,comment)" v-if="comment.sub_amount > 0">子评论数({{comment.sub_amount}})</a>
            <a v-if="comment.depth < 2" href="javascript:;" @click="replyComment(comment.id,comment.created_by)">回复他/她</a>&nbsp;
            <a href="javascript:;">点赞</a>
          </span>
        </Row>
      </p>
      <!-- 递归,子评论区域 -->
      <CommentArea v-if="comment.sub_amount > 0" :parent_id="comment.id" :parent_comment="comment"
                   :theme_pk="theme_pk" :theme_type="theme_type" :key="comment.id"/>
    </div>

    <!-- 顶级评论支持分页 -->
    <Page v-if="parent_id == 0 && total > 0" :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>

    <!-- 评论表单 -->
    <Modal
      v-model="showCommentForm"
      width="800"
      title="回复"
      :mask-closable="false">
      <CommentForm v-if="showCommentForm" :parent_id="_parent_id" :theme_pk="theme_pk" :theme_type="theme_type"
        :refer_user_name="_refer_user_name" @refreshComment="refreshComment"/>
    </Modal>
  </div>
</template>

<script>
  import {FilterComment} from "../../api/index"
  import CommentForm from "./CommentForm"

  export default {
    name: "CommentArea",
    // 评论清单
    props:{
      parent_comment:{
        type:Object,
        default:null,
      },
      parent_id:{
        type:Number,
        default: -1,
      },
      theme_pk:{
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
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        comments:[],
        showCommentForm:false,
        // 回复评论,两个参数分别是被评论id,被评论人
        _parent_id:0,
        _refer_user_name:"",
      }
    },
    methods:{
      handleChange(page){
        this.current_page = page;
        this.refreshComment();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshComment();
      },
      // 刷新当前父级评论对应的评论列表
      refreshComment:async function(comment_type){
        if(comment_type == undefined){
          comment_type = "all";
        }
        const result = await FilterComment(this.theme_pk, this.theme_type, this.parent_id, comment_type, this.offset, this.current_page);
        if(result.status=="SUCCESS"){
          this.showCommentForm = false;
          this.comments = result.comments;
          if(this.parent_id == 0 && result.paginator != null){
            this.total = result.paginator.totalcount;     // 顶级评论支持分页
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
      this.refreshComment();
    },
    watch:{
      "theme_pk": "refreshComment",      // 如果 theme_pk 有变化,会再次执行该方法
      "parent_comment": "refreshComment",       // 如果 comment 有变化,会再次执行该方法
    }
  }
</script>

<style scoped>
  a{
    color:red;
    margin-right: 10px;
  }
</style>
