<template>
  <div>
    <Row>
      <Col span="6">
        <IBeautifulCard title="图书目录">
          <div slot="content" style="padding:10px;">
            <p>图书名称：{{$route.query.book_name}}</p>

            <div v-if="bookBlogs && bookBlogs.length > 0">
              <p v-for="bookBlog in bookBlogs" style="margin-left: 15px;">
                <IBeautifulLink2 @onclick="editBookBlog(bookBlog)">{{bookBlog.blog_title}}</IBeautifulLink2>
              </p>
            </div>
            <div v-else>
              暂未创建目录,直接在右边创建奥
            </div>
          </div>
        </IBeautifulCard>
      </Col>
      <Col span="18">
        <span v-if="$route.query.book_id > 0">
          <BlogEdit ref="blogEdit" :book-id="$route.query.book_id" :success-emit="true" @successEmitFunc="refreshBookInfo"/>
        </span>
      </Col>
    </Row>
  </div>
</template>

<script>
  import BlogEdit from "../BlogEdit";
  import {BookBlogList} from "../../../api";
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2"
  export default {
    name: "BookEdit",
    components: {IBeautifulCard, BlogEdit,IBeautifulLink2},
    data(){
      return {
        bookBlogs:[],
      }
    },
    methods:{
      editBookBlog:function (bookBlog){
        this.$refs.blogEdit.refreshBlogDetail(bookBlog.id);
      },
      refreshBookInfo:async function () {
        const result = await BookBlogList(this.$route.query.book_id);
        if(result.status == "SUCCESS"){
          this.bookBlogs = result.books;
        }
      },
    },
    mounted(){
      this.refreshBookInfo();
    }
  }
</script>

<style scoped>

</style>
