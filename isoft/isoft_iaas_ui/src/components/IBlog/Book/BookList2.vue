<template>
  <div>
    <ul v-for="book in books">
      <li style="float: left;list-style: none;">
        <div style="padding: 20px;">
          <img src="../../../assets/default.png" height="90px" width="120px"/>
          <p>{{book.book_name | filterLimitFunc}}</p>
        </div>
      </li>
    </ul>
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
