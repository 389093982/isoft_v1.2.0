<template>
  <div style="margin: 0 15px;padding-left:20px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;">
    <Row>
      <Col span="18" style="padding: 0 0 20px;border-right: 1px solid #e6e6e6;">
        <!-- 内外边距：上右下左 -->
        <Row style="padding: 15px 10px 10px 25px;">
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
              <div class="bookImg">
                <router-link :to="{path:'/iblog/book_detail',query:{book_id:book.id}}">
                  <img v-if="book.book_img" :src="book.book_img" height="160px" width="120px"/>
                  <img v-else src="../../../assets/default.png" height="160px" width="120px"/>
                  <p style="font-size: 12px;">{{book.book_name | filterLimitFunc}}</p>
                </router-link>
              </div>
              <div class="bookOper" v-if="mine" style="margin-top: 10px;padding-left: 30px;">
                <Row :gutter="10">
                  <Col span="12">
                    <IFileUpload btn-size="small" :auto-hide-modal="true"
                                 :extra-data="book.id" @uploadComplete="uploadComplete" action="/api/iwork/fileUpload/default" uploadLabel="换张图片"/>
                  </Col>
                  <Col span="12">
                    <IBeautifulLink2 @onclick="deleteBook(book.id)">删除</IBeautifulLink2>
                  </Col>
                </Row>
                <Row :gutter="10">
                  <Col span="12">
                    <IBeautifulLink2 @onclick="showBookEditModal2(book)">修改信息</IBeautifulLink2>
                  </Col>
                  <Col span="12">
                    <IBeautifulLink2 @onclick="$router.push({path:'/iblog/mine/book_edit',query:{book_id:book.id,book_name:book.book_name}})">编辑</IBeautifulLink2>
                  </Col>
                </Row>
              </div>
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
      </Col>
      <Col span="6" style="padding: 0 0 20px;border-right: 1px solid #e6e6e6;">
        <HotUser/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {BookList,BookEdit,UpdateBookIcon,DeleteBookById} from "../../../api"
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import IKeyValueForm from "../../Common/form/IKeyValueForm";
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2"
  import IFileUpload from "../../Common/file/IFileUpload"
  import HotUser from "../../User/HotUser";
  import IndexCarousel from "../../ILearning/IndexCarousel";

  export default {
    name: "BookList",
    components: {IndexCarousel, HotUser, IBeautifulLink2, IKeyValueForm, IBeautifulCard,ISimpleConfirmModal,IFileUpload},
    data(){
      return {
        books:[],
        mine:false,
      }
    },
    methods:{
      deleteBook:async function(book_id){
        const result = await DeleteBookById(book_id);
        if(result.status == "SUCCESS"){
          this.refreshBookList();
        }
      },
      uploadComplete: async function (data) {
        if(data.status == "SUCCESS"){
          if(data.status == "SUCCESS"){
            let uploadFilePath = data.fileServerPath;
            let bookId = data.extraData;
            const result = await UpdateBookIcon(bookId, uploadFilePath);
            if(result.status == "SUCCESS"){
              this.refreshBookList();
            }
          }
        }
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
        if(value && value.length > 15) {
          value= value.substring(0,15) + '...';
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
  }
  .bookImg{
    padding: 10px 9px 0;
    width: 160px;
    border: 1px solid #FFFFFF;
    overflow: hidden;
    text-align: center;
    position: relative;
  }
  .bookImg:hover{
    background-color: #f4f4f4;
    border: 1px solid #d0cdd2;
  }
  li a:hover {
    color:red;
  }
</style>
