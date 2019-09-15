<template>
  <div>
    <Row>
      <Col span="20">
        <ul v-for="book in books">
          <li style="float: left;list-style: none;">
            <router-link :to="{path:'/iblog/book_detail',query:{book_id:book.id}}">
              <img src="../../../assets/default.png" height="90px" width="120px"/>
              <p>{{book.book_name | filterLimitFunc}}</p>
            </router-link>
          </li>
        </ul>
      </Col>
      <Col span="4">
        <Button @click="$router.push({path:'/iblog/mine/book_edit'})">
          我的图书管理
        </Button>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {BookList} from "../../../api"

  export default {
    name: "BookList2",
    data(){
      return {
        books:[],
      }
    },
    methods:{
      refreshBookList:async function () {
        const result = await BookList();
        if(result.status == "SUCCESS"){
          this.books = result.books;
        }
      }
    },
    mounted(){
      this.refreshBookList();
    },
    filters:{
      // 内容超长则显示部分
      filterLimitFunc:function (value) {
        if(value && value.length > 10) {
          value= value.substring(0,10) + '...';
        }
        return value;
      },
    }
  }
</script>

<style scoped>
  /* 引入公共样式库 */
  @import "../../../../static/css/common.css";

  a{
    color: black;
  }
  li{
    float: left;
    padding: 10px 9px 0;
    width: 140px;
    height: 135px;
    overflow: hidden;
    text-align: center;
    position: relative;
  }
  li:hover{
    background-color: #f4f4f4;
    border: 1px solid #d0cdd2;
  }
  li:hover a{
    color:red;
  }
</style>
