<template>
  <div>
    <Table border :columns="columns1" :data="globalVars" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {GlobalVarList} from "../../../api"

  export default {
    name: "GlobalVarList",
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        search:"",
        globalVars: [],
        columns1: [
          {
            title: 'name',
            key: 'name',
            width: 350,
          },
          {
            title: 'value',
            key: 'value',
          },
        ],
      }
    },
    methods:{
      handleChange(page){
        this.current_page = page;
        this.refreshGlobalVarList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshGlobalVarList();
      },
      refreshGlobalVarList:async function () {
        const result = await GlobalVarList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.globalVars = result.globalVars;
          this.total = result.paginator.totalcount;
        }
      }
    },
    mounted: function () {
      this.refreshGlobalVarList();
    },
  }
</script>

<style scoped>

</style>
