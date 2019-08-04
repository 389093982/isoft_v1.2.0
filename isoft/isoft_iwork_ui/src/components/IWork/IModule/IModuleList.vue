<template>
  <div>
    <Table border :columns="columns1" :data="modules" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {ModuleList} from "../../../api"

  export default {
    name: "IModuleList",
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        // 搜索条件
        search:"",
        modules:[],
        columns1:[
          {
            title: 'module_name',
            key: 'module_name',
          },
          {
            title: 'module_desc',
            key: 'module_desc',
          },
        ],
      }
    },
    methods:{
      refreshModuleList:async function () {
        const result = await ModuleList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.modules = result.modules;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshModuleList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshModuleList();
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshModuleList();
      },
    },
    mounted(){
      this.refreshModuleList();
    }
  }
</script>

<style scoped>

</style>
