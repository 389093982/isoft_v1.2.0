<template>
  <div>
    <!--文章-->
    <div class="Article">
      <!--文章编辑-->
      <div class="Article-Edit">
        <ArticleEdit @toPublishArticle="toPublishArticle" :article="editArticle"></ArticleEdit>
      </div>

      <!--文章列表-->
      <div class="Article-List">
        <ArticleList @toEditArticle="toEditArticle"></ArticleList>
      </div>
    </div>

    <!--共享大厅-->
    <div class="SharingHall">
      <SharingHall :article="publishArticle"></SharingHall>
    </div>

  </div>
</template>

<script>
  import SharingHall from "../ShareArticle/SharingHall"
  import ArticleList from "../ShareArticle/ArticleList"
  import ArticleEdit from "../ShareArticle/ArticleEdit"
  export default {
    name: "ShareArticlePlace",
    components:{ArticleEdit,SharingHall,ArticleList},
    data(){
      return{
        editArticle:{title:"",content:""},
        publishArticle:{title:"",content:""},
      }
  },
    methods:{
      toEditArticle:function (article) {
        this.editArticle = article
      },
      toPublishArticle:function (article) {
        this.publishArticle=article
      }
    },
  }
</script>

<style scoped>
  .Article{ float:left;width:53%;border:1px }
  .Article-List{ float:left;width:30%;border:1px }
  .Article-Edit{ float:right;width:70%;border:1px }
  .SharingHall{ float:right;width:45%;border:1px;}
</style>




