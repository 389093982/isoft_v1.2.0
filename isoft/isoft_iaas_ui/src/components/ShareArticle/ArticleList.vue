<template>
    <div>
        <div>{{thisIsArticleList}}</div>
        <div> </div>
        <div v-for="item in items" class="SingleArticle" :key="item.title">{{item.title}}</div>
    </div>
</template>

<script>
  export default {
    name: "ArticleList",
    data(){
      return{
        thisIsArticleList:"文章列表",
        items:[
          {title:"我的大学"},
          {title:"勇敢的第一步"},
          {title:"学会学习"},
          {title:"朋友的建议"},
          {title:"不放弃"},
          {title:"遇见了你"},
          {title:"我的大学"},
          {title:"勇敢的第一步"},
          {title:"学会学习"},
          {title:"朋友的建议"},
          {title:"不放弃"},
          {title:"遇见了你"},
          {title:"我的大学"},
          {title:"勇敢的第一步"},
          {title:"学会学习"},
          {title:"朋友的建议"},
          {title:"不放弃"},
          {title:"遇见了你"},
        ]
      }
    },
  }
</script>

<style scoped>
  .SingleArticle{color: #ff6900}
</style>
