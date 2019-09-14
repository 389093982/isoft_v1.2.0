<template>
  <div>
    <Row>
      <Col span="6">
        <BookList @chooseBook="chooseBook"/>

        <IBeautifulCard v-if="book" title="图书目录">
          <span slot="header_right">
            {{book.book_name}}
          </span>
          <span slot="content">
            <p v-for="bookBlog in bookBlogs" style="margin-left: 15px;">
              <IBeautifulLink2 @onclick="editBookBlog(bookBlog)">{{bookBlog.blog_title}}</IBeautifulLink2>
            </p>
          </span>
        </IBeautifulCard>
      </Col>
      <Col span="18">
        <span v-if="book">
          <BlogEdit ref="blogEdit" :book-id="book.id" :success-emit="true" @successEmitFunc="refreshBookInfo"/>
        </span>
      </Col>
    </Row>
  </div>
</template>

<script>
  import BookList from "./BookList";
  import BlogEdit from "../BlogEdit";
  import {BookBlogList} from "../../../api";
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2"
  export default {
    name: "BookEdit",
    components: {IBeautifulCard, BlogEdit, BookList,IBeautifulLink2},
    data(){
      return {
        book:null,
        bookBlogs:[],
      }
    },
    methods:{
      chooseBook:function (book) {
        this.book = book;
        this.refreshBookInfo();
      },
      editBookBlog:function (bookBlog){
        this.$refs.blogEdit.refreshBlogDetail(bookBlog.id);
      },
      refreshBookInfo:async function () {
        const result = await BookBlogList(this.book.id);
        if(result.status == "SUCCESS"){
          this.bookBlogs = result.books;
        }
      }
    }
  }
</script>

<style scoped>

</style>
