<template>
  <div>
    <div v-for="rowData in rowDatas" style="background-color: rgba(236,236,236,0.4);margin: 5px;padding: 10px;">
      <span v-for="(value,key,index) in rowData" style="margin-right: 10px;"> <Tag color="orange">字段名：{{key}}</Tag> {{value}}</span>
    </div>

    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {GetAuditHandleData} from "../../../api"
  export default {
    name: "AuditDetailHandle",
    data(){
      return {
        rowDatas:[],
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        search:"",
      }
    },
    methods:{
      handleChange(page){
        this.current_page = page;
        this.refreshHandleData();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshHandleData();
      },
      refreshHandleData:async function () {
        const result = await GetAuditHandleData(this.$route.query.task_name, this.current_page, this.offset);
        if(result.status == "SUCCESS"){
          this.rowDatas = result.rowDatas;
          this.total = result.totalcount;
        }
      }
    },
    mounted(){
      this.refreshHandleData();
    }
  }
</script>

<style scoped>

</style>
