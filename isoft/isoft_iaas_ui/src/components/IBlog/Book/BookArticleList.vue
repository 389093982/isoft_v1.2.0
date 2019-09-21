<template>
  <div>
    <Row>
      <Col span="6" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <div>
          <Button @click="$router.push({path:'/iblog/mine/book_list',query:{type:'mine'}})">管理我的书单</Button>

          <p v-for="bookArticle in bookArticles">
            <Icon type="ios-paper-outline"/>
            <IBeautifulLink2 @onclick="showDetail(bookArticle)">{{bookArticle.blog_title | filterLimitFunc}}</IBeautifulLink2>
          </p>
        </div>
      </Col>
      <Col span="18" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <div style="text-align: right;">
          <IBeautifulLink2>上一篇</IBeautifulLink2>
          <IBeautifulLink2>下一篇</IBeautifulLink2>
        </div>

        <div v-if="bookArticle">
          <IShowMarkdown v-if="bookArticle.content" :content="bookArticle.content"/>
        </div>

        <div style="text-align: right;">
          <IBeautifulLink2>上一篇</IBeautifulLink2>
          <IBeautifulLink2>下一篇</IBeautifulLink2>
        </div>
      </Col>
    </Row>

    <HorizontalLinks :placement_name="GLOBAL.want_to_find"/>
  </div>
</template>

<script>
  import {BookArticleList} from "../../../api";
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";
  import IShowMarkdown from "../../Common/markdown/IShowMarkdown"
  import HorizontalLinks from "../../Elementviewers/HorizontalLinks";

  export default {
    name: "BookArticleList",
    components: {HorizontalLinks, IBeautifulLink2,IShowMarkdown},
    data(){
      return {
        bookArticles:[],
        bookArticle:null,
      }
    },
    methods:{
      showDetail:function(bookArticle){
        this.bookArticle = bookArticle;
      },
      refreshBookInfo:async function (book_id) {
        const result = await BookArticleList(book_id);
        if(result.status == "SUCCESS"){
          this.bookArticles = result.books;
        }
      }
    },
    mounted(){
      if(this.$route.query.book_id != undefined && this.$route.query.book_id != null){
        this.refreshBookInfo(this.$route.query.book_id);
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
