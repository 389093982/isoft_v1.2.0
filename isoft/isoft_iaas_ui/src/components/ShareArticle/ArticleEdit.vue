<template>
  <div>
      <div>{{onView}}</div>
      <div v-if="article.title" class="SingleArticle-title">{{article.title}}</div>
      <div v-if="article.title" class="SingleArticle-content">
        <input_textarea :content="article.content" :rows="rows" :maxlength="maxlength"></input_textarea>
      </div>
      <slot name="title"></slot>
      <slot name="content"></slot>
  </div>
</template>

<script>
    import input_textarea from "../Utils/input_textarea"
    export default {
      name: "ArticleEdit",
      components:{input_textarea},
      data(){
          return{
            onView:"正在预览",
            article:{
              title:"勇敢的第一步",
              content:"这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容," +
              "这里是文章内容,这里是文章内容,这里是文章内容,"
            },
            rows:15,
            maxlength:2000,
          }
      }
    }
</script>

<style scoped>
  .SingleArticle-title{color: #ff6900}
  .SingleArticle-content{color: #1b1b44}
</style>
