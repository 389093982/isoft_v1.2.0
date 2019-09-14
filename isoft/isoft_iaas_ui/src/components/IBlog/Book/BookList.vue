<template>
  <span>
    <IBeautifulCard title="我的全部书单">
      <span slot="header_right">
        <Button size="small" type="success" @click="showBookEditModal">新增</Button>
      </span>
      <span slot="content">
        <p v-for="book in books" style="margin-left: 15px;">
          <IBeautifulLink2 @onclick="chooseBook(book)">{{book.book_name}}</IBeautifulLink2>
          <Button style="float: right;margin-right: 15px;" size="small" type="success" @click="showBookEditModal2(book)">编辑</Button>
        </p>
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
        this.$refs.bookEditModal.showModal();
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
    }
  }
</script>

<style scoped>

</style>
