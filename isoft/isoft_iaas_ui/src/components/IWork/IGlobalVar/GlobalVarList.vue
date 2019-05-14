<template>
  <div>

    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="globalVars" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {GlobalVarList} from "../../../api"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"

  export default {
    name: "GlobalVarList",
    components:{ISimpleLeftRightRow,ISimpleSearch},
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
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
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
