<template>
  <div>
    <div style="text-align: right;margin-bottom: 10px;">
      <Button type="success" size="small" @click="renderWorkstepList">运行时数据</Button>
      <Button type="error" size="small" @click="highlightError = !highlightError">高亮显示错误</Button>
    </div>

    <Table :columns="columns1" :data="runLogDetails" :row-class-name="rowClassName" size="small"></Table>
  </div>
</template>

<script>
  import {GetLastRunLogDetail} from "../../../api"

  export default {
    name: "RunLogDetail",
    data(){
      return {
        highlightError:false,
        runLogDetails:[],
        runLogRecord: null,
        columns1: [
          {
            title: 'work_step_name',
            key: 'work_step_name',
            width:150,
          },
          {
            title: 'log_level',
            key: 'log_level',
            width:120,
            filters: [
              {
                label: 'ERROR',
                value: 'ERROR',
              },
              {
                label: 'ALL',
                value: 'ALL',
              }
            ],
            filterMultiple: false,
            filterMethod (value, row) {
              if (value == "ERROR") {
                return row.log_level == "ERROR";
              }
              return true;
            }
          },
          {
            title: 'detail',
            key: 'detail',
            render: (h,params)=>{
              return h('div',{
                domProps:{
                  innerHTML:params.row.detail,
                }
              })
            }
          },
        ],
      }
    },
    methods:{
      renderWorkstepList: function(){
        this.$router.push({ path: '/iwork/workstepList',
          query: { work_id: this.runLogRecord.work_id, work_name: this.runLogRecord.work_name, tracking_id:this.$route.query.tracking_id }});
      },
      refreshRunLogDetail:async function () {
        const result = await GetLastRunLogDetail(this.$route.query.tracking_id);
        if(result.status=="SUCCESS"){
          this.runLogRecord = result.runLogRecord;
          this.runLogDetails = result.runLogDetails;
        }
      },
      rowClassName (row, index) {
        if (row.log_level === "ERROR" && this.highlightError) {
          return 'demo-table-error-row';
        }
        return '';
      }
    },
    mounted(){
      this.refreshRunLogDetail();
    }
  }
</script>

<style>
  /*
    vue中慎用style的scoped属性
    scoped肯定是解决了样式私有化的问题,但同时也引入了新的问题,scoped设计的初衷就是让样式变得不可修改
  */
  .ivu-table .demo-table-error-row td{
    background-color: pink;
  }
</style>
