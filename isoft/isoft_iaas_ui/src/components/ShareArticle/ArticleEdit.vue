<template>
  <div>
    <div>正在编辑</div>
    <!--标题-->
    <div class="SingleArticle-title">
      <Input v-model="article.title" placeholder="Enter title..." maxlength="20" style="width: 100%" />
    </div>
    <!--内容-->
    <div class="SingleArticle-content">
      <Input type="textarea" v-model="article.content" @keydown.tab.native="tab" :rows="rows" :maxlength="max" show-word-limit="true" placeholder="Enter content" style="width:100%"/>
      <span style="float: right;color: #9ea7b4">{{getAlreadyInput}}/{{max}}</span>
    </div>
    <!--按钮-->
    <Button type="success" @click="saveArticle">保存</Button>
    <Button type="primary" @click="publishArticle">发表</Button>

    <div> </div>
    <slot name="title"></slot>
    <slot name="content"></slot>
  </div>
</template>

<script>
    import {saveArticle,publishArticle} from "../../api"
    export default {
      props:{
        article:{
          title:String,
          content:String,
        },
      },
      name: "ArticleEdit",
      components:{},
      data(){
        return{
          rows:17,
          max:2000,
          alreadyInput:0,
        }
      },
      computed:{
        getAlreadyInput:function(){
          return this.article.content.length
        }
      },
      methods:{
        saveArticle:async function () {
          // let saveResult = await saveArticle(this.article.title,this.article.content);
          this.$Message.success("保存成功")
        },
        publishArticle:async function () {
          // let publishResult = await publishArticle(this.article.title,this.article.content);
          this.$Message.success("发布成功")
          this.$emit("toPublishArticle",this.article)
        },
        tab:function(){
          this.$Message.success("tab")
        },
      },
    }
</script>

<style scoped>
  .SingleArticle-title{color: #ff6900}
  .SingleArticle-content{color: #1b1b44}
</style>
