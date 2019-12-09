<template>
  <div>
    <Row>
      <!-- 评论表单 -->
      <Col span="14" style="padding-right: 10px;">
        <Input v-model.trim="content" type="textarea" :rows="8" placeholder="发表你的评论信息！" />
        <Button size="small" type="success" style="margin: 5px;float: right;" @click="submitComment('comment')">发表评论</Button>
        <Button size="small" type="error" style="margin: 5px;float: right;" @click="submitComment('question')">我要提问</Button>
      </Col>
      <Col span="10" style="border: 1px solid #e9e9e9;font-size:12px;padding: 10px;">
        <p>发表评论需知：</p>
        <p>1、请勿在评论中发表违法违规信息</p>
        <p>2、谢绝人身攻击、地域歧视、刷屏、广告等恶性言论</p>
        <p>3、所有评论均代表玩家本人意见，不代表官方立场</p>
        <p>4、用户发表的评论，经管理员审核后方可显示</p>
        <p>5、如果您有任何疑问，请在此以评论方式留言给我们</p>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {AddComment} from "../../api/index"
  import {checkFastClick} from "../../tools"

  export default {
    name: "CommentForm",
    // 父组件传递给子组件的字段
    props:["parent_id", "theme_pk", "theme_type", "refer_user_name"],
    data(){
      return {
        content:"",
      }
    },
    methods:{
      submitComment: async function (comment_type) {
        if (checkFastClick()){
          this.$Message.error("点击过快,请稍后重试!");
          return;
        }
        if(this.content == undefined || this.content.length < 10){
          this.$Notice.error({
            title: '温馨提示',
            desc: "评论信息过短,需要10个字符以上！"
          });
        }else{
          const result = await AddComment(this.parent_id, this.content, this.theme_pk,
            this.theme_type, comment_type, this.refer_user_name);
          if(result.status=="SUCCESS"){
            this.$Message.success("发表成功!");
            // 调用父组件的 refreshComment 方法
            this.$emit('refreshComment','all');
          }
        }
      }
    }
  }
</script>

<style scoped>

</style>
