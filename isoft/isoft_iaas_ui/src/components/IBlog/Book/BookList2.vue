<template>
  <div>
    <Row>
      <Col span="20">
        <ul v-for="book in books">
          <li style="float: left;list-style: none;">
            <div style="padding: 20px;">
              <img src="../../../assets/default.png" height="90px" width="120px"/>
              <p>{{book.book_name | filterLimitFunc}}</p>
            </div>
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

</style>
