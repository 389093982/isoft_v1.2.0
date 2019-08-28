<template>
  <div style="margin-left: 10px;margin-top: 25px; padding: 20px;background: #fff;border-bottom: 1px solid #f4f4f4;">
    <CatalogAdd v-if="showCatalogAdd" @handleSuccess="handleSuccess"/>
    <Row>
      <Col span="16">我的博客分类 <Icon type="md-add" @click="showCatalogAdd = !showCatalogAdd"/></Col>
      <Col span="8">创建时间</Col>
    </Row>
    <Row v-for="(catalog,index) in catalogs">
      <Col span="16">{{ catalog.catalog_name | filterLimitFunc }}</Col>
      <Col span="8" style="font-size: 12px;"><Time :time="catalog.created_time" type="date"/></Col>
    </Row>
  </div>
</template>

<script>
  import {GetMyCatalogs} from "../../api"
  import CatalogAdd from "./CatalogAdd"

  export default {
    name: "CatalogList",
    components:{CatalogAdd},
    data(){
      return {
        showCatalogAdd:false,
        catalogs:[],
      }
    },
    methods:{
      handleSuccess:function(){
        this.showCatalogAdd = false;
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
      this.refreshMyCatalogList();
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

</style>
