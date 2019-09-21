<template>
  <div>
    <Row>
      <Col span="6" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <Row>
          <Col span="18"><span style="color: green;font-weight: bold;">{{$route.query.book_name}}</span></Col>
          <Col span="6"><Button style="right: 50px;" size="small" @click="createEmptyArticle">新建文章</Button></Col>
        </Row>

        <div v-if="bookArticles && bookArticles.length > 0">
          <p v-for="bookArticle in bookArticles" style="margin-left: 15px;">
            <Icon type="ios-paper-outline"/>
            <IBeautifulLink2 @onclick="editBookArticle(bookArticle)">{{bookArticle.blog_title | filterLimitFunc}}</IBeautifulLink2>
          </p>
        </div>
        <div v-else>
          暂未创建目录,直接在右边创建奥
        </div>
      </Col>
      <Col span="18" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <span>
          <ArticleEdit ref="blogArticleEdit" :book-id="_book_id" :success-emit="true" @successEmitFunc="refreshBookInfo"/>
        </span>
      </Col>
    </Row>
  </div>
</template>

<script>
  import ArticleEdit from "../ArticleEdit";
  import {BookArticleList} from "../../../api";
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2"
  import {checkFastClick} from "../../../tools"

  export default {
    name: "BookCatalogEdit",
    components: {IBeautifulCard, ArticleEdit,IBeautifulLink2},
    data(){
      return {
        bookArticles:[],
      }
    },
    methods:{
      createEmptyArticle:function(){
        if (checkFastClick()){
          this.$Message.error("点击过快,请稍后重试!");
          return;
        }
        this.$refs.blogArticleEdit.createEmptyArticle(parseInt(this.$route.query.book_id));
      },
      editBookArticle:function (bookArticle){
        this.$refs.blogArticleEdit.refreshArticleDetail(bookArticle.id);
      },
      refreshBookInfo:async function () {
        const result = await BookArticleList(this.$route.query.book_id);
        if(result.status == "SUCCESS"){
          this.bookArticles = result.books;
        }
      },
    },
    computed:{
      _book_id:function () {
        return parseInt(this.$route.query.book_id);
      },
    },
    mounted(){
      this.refreshBookInfo();
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
