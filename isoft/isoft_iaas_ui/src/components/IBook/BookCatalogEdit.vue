<template>
  <div>
    <Row>
      <Col span="6" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <Row>
          <Col span="18"><span style="color: green;font-weight: bold;">{{$route.query.book_name}}</span></Col>
          <Col span="6"><Button style="right: 50px;" size="small" @click="createBookCatalog">新建目录</Button></Col>
        </Row>

        <ISimpleConfirmModal ref="bookCatalogEditModal" modal-title="新增/编辑 目录" :modal-width="600" :footer-hide="true">
          <!-- 表单信息 -->
          <Form ref="formValidate" :model="formValidate" :rules="ruleValidate" :label-width="150">
            <FormItem label="目录名称" prop="catalogName">
              <Input v-model.trim="formValidate.catalogName" placeholder="请输入目录名称"></Input>
            </FormItem>
            <FormItem>
              <Button type="success" @click="handleSubmit('formValidate')" style="margin-right: 6px">Submit</Button>
              <Button type="warning" @click="handleReset('formValidate')" style="margin-right: 6px">Reset</Button>
            </FormItem>
          </Form>
        </ISimpleConfirmModal>

        <Scroll height="430" style="margin-top: 5px;">
          <div v-if="bookCatalogs && bookCatalogs.length > 0">
            <p v-for="bookCatalog in bookCatalogs" style="margin-left: 15px;">
              <Icon type="ios-paper-outline"/>
              <IBeautifulLink2 @onclick="editBookArticle(bookCatalog.id)">{{bookCatalog.catalog_name | filterLimitFunc}}</IBeautifulLink2>
            </p>
          </div>
          <div v-else>
            暂未创建目录,直接在右边创建奥
          </div>
        </Scroll>
      </Col>
      <Col span="18" style="background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;padding: 20px;min-height: 500px;">
        <span>
          <BookArticleEdit ref="bookArticleEdit" :success-emit="true" @successEmitFunc="refreshBookCatalogList"/>
        </span>
      </Col>
    </Row>
  </div>
</template>

<script>
  import BookArticleEdit from "./BookArticleEdit";
  import {BookArticleList,BookCatalogEdit,BookCatalogList} from "../../api";
  import IBeautifulCard from "../Common/card/IBeautifulCard"
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2"
  import ISimpleConfirmModal from "../Common/modal/ISimpleConfirmModal";

  export default {
    name: "BookCatalogEdit",
    components: {ISimpleConfirmModal, IBeautifulCard,IBeautifulLink2,BookArticleEdit},
    data(){
      return {
        bookCatalogs:[],
        bookArticles:[],
        formValidate: {
          catalogName: '',
        },
        ruleValidate: {
          catalogName: [
            { required: true, message: 'catalogName 不能为空!', trigger: 'blur' }
          ],
        },
      }
    },
    methods:{
      handleSubmit (name) {
        this.$refs[name].validate(async (valid) => {
          if (valid) {
            const result = await BookCatalogEdit(parseInt(this.$route.query.book_id), this.formValidate.catalogName);
            if(result.status == "SUCCESS"){
              this.$Message.success("新建成功!");
              this.$refs.bookCatalogEditModal.hideModal();
              this.refreshBookCatalogList();
            }
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
      createBookCatalog: function(){
        this.$refs.bookCatalogEditModal.showModal();
      },
      editBookArticle:function (bookCatalogId){
        this.$refs.bookArticleEdit.refreshBookArticleDetail(bookCatalogId);
      },
      refreshBookCatalogList:async function(){
        const result = await BookCatalogList(this.$route.query.book_id);
        if(result.status == "SUCCESS"){
          this.bookCatalogs = result.bookCatalogs;
        }
      },
      refreshBookInfo:async function () {
        const result = await BookArticleList(this.$route.query.book_id);
        if(result.status == "SUCCESS"){
          this.bookArticles = result.books;
        }
      },
    },
    computed:{
      _book_id:function () {
        return parseInt(this.$route.query.book_id);
      },
    },
    mounted(){
      this.refreshBookCatalogList();
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
