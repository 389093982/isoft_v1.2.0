<template>
  <div>
    <Row>
      <Col span="6" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <div>
          <Button @click="$router.push({path:'/iblog/mine/book_list',query:{type:'mine'}})">管理我的书单</Button>

          <p v-for="bookCatalog in bookCatalogs">
            <Icon type="ios-paper-outline"/>
            <IBeautifulLink2 @onclick="showDetail(bookCatalog.id)">{{bookCatalog.catalog_name | filterLimitFunc}}</IBeautifulLink2>
          </p>
        </div>
      </Col>
      <Col span="18" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <div style="text-align: right;">
          <IBeautifulLink2>上一篇</IBeautifulLink2>
          <IBeautifulLink2>下一篇</IBeautifulLink2>
        </div>

        <div v-if="bookArticle" style="min-height: 400px;">
          <IShowMarkdown v-if="bookArticle.content" :content="bookArticle.content"/>
        </div>

        <div style="text-align: right;">
          <IBeautifulLink2>上一篇</IBeautifulLink2>
          <IBeautifulLink2>下一篇</IBeautifulLink2>
        </div>
      </Col>
    </Row>

    <HorizontalLinks :placement_name="GLOBAL.element_want_to_find"/>
  </div>
</template>

<script>
  import {BookArticleList,BookCatalogList,ShowBookArticleDetail} from "../../api";
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2";
  import IShowMarkdown from "../Common/markdown/IShowMarkdown"
  import HorizontalLinks from "../Elementviewers/HorizontalLinks";

  export default {
    name: "BookArticleDetail",
    components: {HorizontalLinks, IBeautifulLink2,IShowMarkdown},
    data(){
      return {
        bookArticles:[],
        bookArticle:null,
        bookCatalogs:[],
      }
    },
    methods:{
      showDetail:async function (book_catalog_id) {
        const result = await ShowBookArticleDetail(book_catalog_id);
        if(result.status=="SUCCESS"){
          if(result.bookArticle != null){
            this.bookArticle = result.bookArticle;
          }
        }else{
          this.$Message.error("加载失败!");
        }
      },
      refreshBookInfo:async function (book_id) {
        const result = await BookArticleList(book_id);
        if(result.status == "SUCCESS"){
          this.bookArticles = result.books;
          if(this.bookArticles.length > 0){
            this.showDetail(this.bookArticles[0]);
          }
        }
      },
      refreshBookCatalogList:async function(book_id){
        const result = await BookCatalogList(book_id);
        if(result.status == "SUCCESS"){
          this.bookCatalogs = result.bookCatalogs;
        }
      },
    },
    mounted(){
      if(this.$route.query.book_id != undefined && this.$route.query.book_id != null){
        this.refreshBookCatalogList(this.$route.query.book_id);
      }
    },
    filters:{
      // 内容超长则显示部分
      filterLimitFunc:function (value) {
        if(value && value.length > 30) {
          value= value.substring(0,30) + '...';
        }
        return value;
      },
    }
  }
</script>

<style scoped>

</style>
