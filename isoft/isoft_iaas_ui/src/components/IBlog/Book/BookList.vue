<template>
  <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;">
    <Row style="padding: 10px;">
      <Col span="2">
        <IBeautifulLink2 @onclick="$router.push({path:'/iblog/book_list'})">全部书单</IBeautifulLink2>
      </Col>
      <Col span="2">
        <IBeautifulLink2 @onclick="$router.push({path:'/iblog/book_list'})">热门书单</IBeautifulLink2>
      </Col>
      <Col span="2">
        <IBeautifulLink2 @onclick="$router.push({path:'/iblog/mine/book_list',query:{type:'mine'}})">我的书单</IBeautifulLink2>
      </Col>
    </Row>

    <div style="min-height: 450px;">
      <ul style="overflow:hidden">
        <li v-for="book in books" style="float: left;list-style: none;">
          <router-link :to="{path:'/iblog/book_detail',query:{book_id:book.id}}">
            <img src="../../../assets/default.png" height="90px" width="120px"/>
            <p>{{book.book_name | filterLimitFunc}}</p>
          </router-link>
          <p v-if="mine">
            <IBeautifulLink2 @onclick="showBookEditModal2(book)">修改</IBeautifulLink2>
            <IBeautifulLink2 @onclick="deleteBook">删除</IBeautifulLink2>
            <IBeautifulLink2 @onclick="$router.push({path:'/iblog/mine/book_edit',query:{book_id:book.id,book_name:book.book_name}})">编辑</IBeautifulLink2>
          </p>
        </li>
      </ul>
      <div style="text-align: right;margin: 20px 100px 50px 0;">
        <Button v-if="mine" @click="showBookEditModal">新增书单</Button>
      </div>
    </div>

    <ISimpleConfirmModal ref="bookEditModal" modal-title="新增/编辑 Book" :modal-width="600" :footer-hide="true">
      <IKeyValueForm ref="bookEditForm" form-key-label="book_name" form-value-label="book_desc"
                     form-key-placeholder="请输入书名" form-value-placeholder="请输入描述"
                     @handleSubmit="editBook">
      </IKeyValueForm>
    </ISimpleConfirmModal>
  </div>
</template>

<script>
  import {BookList,BookEdit} from "../../../api"
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import IKeyValueForm from "../../Common/form/IKeyValueForm";
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";

  export default {
    name: "BookList",
    components: {IBeautifulLink2, IKeyValueForm, IBeautifulCard,ISimpleConfirmModal},
    data(){
      return {
        books:[],
        mine:false,
      }
    },
    methods:{
      deleteBook:function(){
        alert(111111111);
      },
      showBookEditModal2:function (book) {
        this.$refs.bookEditModal.showModal();
        this.$refs.bookEditForm.initFormData(book.id, book.book_name, book.book_desc);
      },
      editBook:async function (book_id, book_name, book_desc) {
        const result = await BookEdit(book_id, book_name, book_desc);
        if(result.status == "SUCCESS"){
          this.$refs.bookEditModal.hideModal();
          this.$refs.bookEditForm.handleSubmitSuccess("提交成功!");
          this.refreshBookList();
        }else{
          this.$refs.bookEditForm.handleSubmitError("提交失败!");
        }
      },
      showBookEditModal:function () {
        this.$refs.bookEditForm.handleReset('formValidate');
        this.$refs.bookEditModal.showModal();
      },
      refreshBookList:async function () {
        this.mine = this.$route.query.type == 'mine';

        const result = await BookList();
        if(result.status == "SUCCESS"){
          this.books = result.books;
        }
      }
    },
    mounted(){
      this.refreshBookList();
    },
    watch:{
      // 监听路由是否变化
      '$route':'refreshBookList'
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
    width: 160px;
    min-height: 160px;
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
