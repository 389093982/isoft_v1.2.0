<template>
  <Row :gutter="10">
    <Col span="6">
      <div style="background-color: #fff;border: 1px solid #e6e6e6;padding: 20px;min-height: 500px;">
        <Button size="small" @click="editBookCatalog">新建目录</Button>

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

        <div style="margin-top: 5px;min-height: 250px;">
          <div v-if="bookCatalogs && bookCatalogs.length > 0">
            <dl>
              <dt><span style="color: green;font-weight: bold;">{{$route.query.book_name}}</span></dt>
              <dd class="isoft_font isoft_inline_ellipsis" style="color: #333333;" v-for="bookCatalog in bookCatalogs">
                <Icon type="ios-paper-outline"/>
                <span class="isoft_hover_red" @click="editBookArticle(bookCatalog.id)">{{bookCatalog.catalog_name}}</span>
                <a class="catalogEditIcon" style="margin-left: 5px;" @click="editBookCatalog(bookCatalog.id)"><Icon type="md-create"/></a>
              </dd>
            </dl>
          </div>
          <div v-else>
            暂未创建目录,直接在右边创建奥
          </div>
        </div>
      </div>
    </Col>
    <Col span="18">
      <BookArticleEdit ref="bookArticleEdit" :success-emit="true" @successEmitFunc="refreshBookCatalogList"/>
    </Col>
  </Row>
</template>

<script>
  import BookArticleEdit from "./BookArticleEdit";
  import {BookArticleList,BookCatalogEdit,BookCatalogList,ShowBookCatalogDetail} from "../../api";
  import IBeautifulCard from "../Common/card/IBeautifulCard"
  import IBeautifulLink from "../Common/link/IBeautifulLink"
  import ISimpleConfirmModal from "../Common/modal/ISimpleConfirmModal";

  export default {
    name: "BookCatalogEdit",
    components: {ISimpleConfirmModal, IBeautifulCard,IBeautifulLink,BookArticleEdit},
    data(){
      return {
        bookCatalogs:[],
        formValidate: {
          id:-1,
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
            const result = await BookCatalogEdit(parseInt(this.$route.query.book_id), this.formValidate.id, this.formValidate.catalogName);
            if(result.status == "SUCCESS"){
              this.$Message.success("编辑成功!");
              this.$refs.bookCatalogEditModal.hideModal();
              this.refreshBookCatalogList();
              this.handleReset('formValidate');
            }
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields();
      },
      editBookCatalog: async function(bookCatalogId){
        if(bookCatalogId > 0){
          const result = await ShowBookCatalogDetail(bookCatalogId);
          if(result.status == "SUCCESS"){
            this.formValidate.id = result.bookCatalog.id;
            this.formValidate.catalogName = result.bookCatalog.catalog_name;
            this.$refs.bookCatalogEditModal.showModal();
          }
        }else{
          this.$refs.bookCatalogEditModal.showModal();
        }

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
    },
    computed:{
      _book_id:function () {
        return parseInt(this.$route.query.book_id);
      },
    },
    mounted(){
      this.refreshBookCatalogList();
    },
  }
</script>

<style scoped>


  dd .catalogEditIcon{
    display: none;
  }
  dd:hover .catalogEditIcon{
    display: inline;
  }
</style>
