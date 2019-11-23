<template>
  <div>
      <div>{{onView}}</div>
      <div v-if="article.title" class="SingleArticle-title">{{article.title}}</div>
      <div v-if="article.title" class="SingleArticle-content">{{article.content}}</div>
      <slot name="title"></slot>
      <slot name="content"></slot>
  </div>
</template>

<script>
    export default {
        name: "ArticleEdit",
      data(){
          return{
            onView:"正在预览",
            article:{
              title:"勇敢的第一步",
              content:"这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容,这里是文章" +
              "内容这里是文章内容,这里是文章内容,这里是文章内容,这里是文章内容"
            }
          }
      }
    }
</script>

<style scoped>
  .SingleArticle-title{color: #ff6900}
  .SingleArticle-content{color: #1b1b44}
</style>
