<template>
  <div>
    <Row>
      <Col span="6" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <div>
          <Button @click="$router.push({path:'/iblog/mine/book_list',query:{type:'mine'}})">管理我的书单</Button>

          <p v-for="bookBlog in bookBlogs">
            <Icon type="ios-paper-outline"/>
            <IBeautifulLink2 @onclick="showDetail(bookBlog)">{{bookBlog.blog_title | filterLimitFunc}}</IBeautifulLink2>
          </p>
        </div>
      </Col>
      <Col span="18" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <div style="text-align: right;">
          <IBeautifulLink2>上一篇</IBeautifulLink2>
          <IBeautifulLink2>下一篇</IBeautifulLink2>
        </div>

        <div v-if="bookBlog">
          <IShowMarkdown v-if="bookBlog.content" :content="bookBlog.content"/>
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
  import {BookBlogList} from "../../../api";
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";
  import IShowMarkdown from "../../Common/markdown/IShowMarkdown"
  import HorizontalLinks from "../../Elementviewers/HorizontalLinks";

  export default {
    name: "BookDetail",
    components: {HorizontalLinks, IBeautifulLink2,IShowMarkdown},
    data(){
      return {
        bookBlogs:[],
        bookBlog:null,
      }
    },
    methods:{
      showDetail:function(bookBlog){
        this.bookBlog = bookBlog;
      },
      refreshBookInfo:async function (book_id) {
        const result = await BookBlogList(book_id);
        if(result.status == "SUCCESS"){
          this.bookBlogs = result.books;
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
