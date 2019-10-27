<template>
  <IBeautifulCard title="我的博客分类">
    <div slot="content" style="padding: 10px;">
      <span v-if="hasLogin">
        <BlogCatalogEdit v-if="showBlogCatalogEdit" @handleSuccess="handleSuccess"/>
        <Row>
          <Col span="16">我的博客分类 <Icon type="md-add" @click="showBlogCatalogEdit = !showBlogCatalogEdit"/></Col>
          <Col span="8">创建时间</Col>
        </Row>
        <Row v-for="(catalog,index) in catalogs">
          <Col span="16">{{ catalog.catalog_name | filterLimitFunc }}</Col>
          <Col span="8" style="font-size: 12px;"><Time :time="catalog.created_time" type="date"/></Col>
        </Row>
      </span>
      <span v-else>
        <ForwardLogin/>
      </span>
    </div>
  </IBeautifulCard>
</template>

<script>
  import {GetMyCatalogs} from "../../api"
  import BlogCatalogEdit from "./BlogCatalogEdit"
  import {CheckHasLogin} from "../../tools"
  import IBeautifulCard from "../../components/Common/card/IBeautifulCard"
  import ForwardLogin from "../SSO/ForwardLogin"

  export default {
    name: "CatalogList",
    components:{ForwardLogin, IBeautifulCard, BlogCatalogEdit},
    data(){
      return {
        showBlogCatalogEdit:false,
        catalogs:[],
      }
    },
    methods:{
      handleSuccess:function(){
        this.showBlogCatalogEdit = false;
        this.refreshMyCatalogList();
      },
      refreshMyCatalogList:async function(){
        const result = await GetMyCatalogs();
        if (result.status == "SUCCESS"){
          this.catalogs = result.catalogs;
        }
      },
    },
    mounted(){
      if(CheckHasLogin()){
        this.refreshMyCatalogList();
      }
    },
    filters:{
      // 内容超长则显示部分
      filterLimitFunc:function (value) {
        if(value && value.length > 15) {
          value= value.substring(0,15) + '...';
        }
        return value;
      },
    },
    computed:{
      hasLogin:function () {
        return CheckHasLogin();
      }
    }
  }
</script>

<style scoped>

</style>
