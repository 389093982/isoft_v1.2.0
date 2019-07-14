<template>
  <span>
    <Table :columns="columns1" :data="runLogDetails" size="small"></Table>
  </span>
</template>

<script>
  import {GetLastRunLogDetail} from "../../../api"

  export default {
    name: "RunLogDetail",
    data(){
      return {
        runLogDetails:[],
        columns1: [
          {
            title: 'work_step_name',
            key: 'work_step_name',
            width:150,
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
      refreshRunLogDetail:async function () {
        const result = await GetLastRunLogDetail(this.$route.query.tracking_id);
        if(result.status=="SUCCESS"){
          this.runLogDetails = result.runLogDetails;
        }
      }
    },
    mounted(){
      this.refreshRunLogDetail();
    }
  }
</script>

<style scoped>

</style>
