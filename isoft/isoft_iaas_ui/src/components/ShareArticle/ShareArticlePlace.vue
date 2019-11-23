<template>
  <div>

    <div class="left-pane">
        <!--共享大厅-->
        <SharingHall></SharingHall>
    </div>

    <div class="center-pane">
      <!--文章编辑-->
      <ArticleEdit>
        <div slot="title">title01</div>
        <div slot="content">content01</div>
      </ArticleEdit>
    </div>

    <div class="right-pane">
      <!--文章列表-->
      <ArticleList></ArticleList>
    </div>

  </div>
</template>

<script>
  import SharingHall from "../ShareArticle/SharingHall"
  import ArticleList from "../ShareArticle/ArticleList"
  import ArticleEdit from "../ShareArticle/ArticleEdit"
  export default {
    name: "ShareArticlePlace",
    components:{ArticleEdit,SharingHall,ArticleList}
  }
</script>

<style scoped>
  .left-pane{ float:left;width:40%;border:1px }
  .center-pane{ float:left;width:40%;border:1px }
  .right-pane{ float:right;width:20%;border:1px }
</style>




