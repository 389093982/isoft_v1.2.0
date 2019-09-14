<template>
  <span>
    <IBeautifulCard title="我的全部书单">
      <span slot="content">
        <Row v-for="book in books" style="padding:0 5px 0 5px;">
          <Col span="16">
            <IBeautifulLink2 @onclick="chooseBook(book)">{{book.book_name | filterLimitFunc}}</IBeautifulLink2>
          </Col>
          <Col span="8" style="text-align: right;">
            <IBeautifulLink2 @onclick="showBookEditModal2(book)">编辑</IBeautifulLink2>
            <IBeautifulLink2 @onclick="showBookEditModal">新增</IBeautifulLink2>
            <IBeautifulLink2 @onclick="deleteBook">删除</IBeautifulLink2>
          </Col>
        </Row>
      </span>
    </IBeautifulCard>

    <ISimpleConfirmModal ref="bookEditModal" modal-title="新增/编辑 Book" :modal-width="600" :footer-hide="true">
      <IKeyValueForm ref="bookEditForm" form-key-label="book_name" form-value-label="book_desc"
                           form-key-placeholder="请输入书名" form-value-placeholder="请输入描述"
                           @handleSubmit="editBook">
      </IKeyValueForm>
    </ISimpleConfirmModal>
  </span>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import IKeyValueForm from "../../Common/form/IKeyValueForm";
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import {BookList,BookEdit} from "../../../api"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";

  export default {
    name: "BookList",
    components: {IBeautifulLink2, IKeyValueForm, IBeautifulCard,ISimpleConfirmModal},
    data(){
      return {
        books:[],
      }
    },
    methods:{
      showBookEditModal:function () {
        this.$refs.bookEditForm.handleReset('formValidate');
        this.$refs.bookEditModal.showModal();
      },
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
      chooseBook:function(book){
        this.$emit('chooseBook', book);
      },
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
        if(value && value.length > 20) {
          value= value.substring(0,20) + '...';
        }
        return value;
      },
    }
  }
</script>

<style scoped>

</style>
