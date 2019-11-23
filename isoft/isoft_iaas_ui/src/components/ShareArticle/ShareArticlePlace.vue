<template>
  <div>

    <!--共享大厅-->
    <div class="SharingHall">
        <SharingHall></SharingHall>
    </div>

    <!--文章-->
    <div class="Article">
      <!--文章编辑-->
      <div class="Article-Edit">
        <ArticleEdit>
          <!--<div slot="title">title01</div>-->
          <!--<div slot="content">content01</div>-->
        </ArticleEdit>
      </div>

      <!--文章列表-->
      <div class="Article-List">
        <ArticleList></ArticleList>
      </div>
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
  .SharingHall{ float:left;width:40%;border:1px;}
  .Article{ float:right;width:60%;border:1px }
  .Article-Edit{ float:left;width:60%;border:1px }
  .Article-List{ float:right;width:30%;border:1px }
</style>




