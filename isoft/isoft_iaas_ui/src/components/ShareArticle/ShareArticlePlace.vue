<template>
  <div>

    <div class="left-pane">
      共享大厅
    </div>

    <div class="center-pane">
      <ArticleEdit>
        <div slot="title">title01</div>
        <div slot="content">content01</div>
      </ArticleEdit>
    </div>
    <div class="right-pane">
      文章列表
    </div>
  </div>
</template>

<script>
  import ArticleEdit from "../ShareArticle/ArticleEdit"
  export default {
    name: "ShareArticlePlace",
    components:{ArticleEdit}
  }
</script>

<style scoped>
  .left-pane{ float:left;width:40%;border:1px }
  .center-pane{ float:left;width:40%;border:1px }
  .right-pane{ float:right;width:20%;border:1px }
</style>




