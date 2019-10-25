<template>
  <div>
    <Table border :columns="columns1" :data="advises" size="small"></Table>

    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {queryPageAdvise} from "../../api"

  export default {
    name: "AdviseList",
    data(){
      return {
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:20,
        advises:[],
        columns1: [
          {
            title: 'advise',
            key: 'advise',
            width: 700,
          },
          {
            title: 'created_by',
            key: 'created_by',
          },
          {
            title: 'created_time',
            key: 'created_time',
          },
        ],
      }
    },
    methods:{
      handleChange(page){
        this.current_page = page;
        this.refreshAdviseList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshAdviseList();
      },
      refreshAdviseList:async function () {
        const result = await queryPageAdvise(this.offset,this.current_page);
        if(result.status=="SUCCESS"){
          this.advises = result.advises;
          this.total = result.paginator.totalcount;
        }
      }
    },
    mounted(){
      this.refreshAdviseList();
    }
  }
</script>

<style scoped>

</style>
