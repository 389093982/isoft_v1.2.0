<template>
  <div style="margin: 10px;">
    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <QuartzAdd slot="left" @handleSuccess="refreshQuartzList"/>
      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="quartzs" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>

    <h3>cron举例说明</h3>
    <p>每隔5秒执行一次：<span style="color: red;">*/5 * * * * ?</span></p>
    <p>每隔1分钟执行一次：<span style="color: red;">0 */1 * * * ?</span></p>
    <p>每天23点执行一次：<span style="color: red;">0 0 23 * * ?</span></p>
    <p>每天凌晨1点执行一次：<span style="color: red;">0 0 1 * * ?</span></p>
    <p>每月1号凌晨1点执行一次：<span style="color: red;">0 0 1 1 * ?</span></p>
    <p>在26分、29分、33分执行一次：<span style="color: red;">0 26,29,33 * * * ?</span></p>
    <p>每天的0点、13点、18点、21点都执行一次：<span style="color: red;">0 0 0,13,18,21 * * ?</span></p>
  </div>
</template>

<script>
  import {formatDate} from "../../../tools/index"
  import {QuartzList,QueryWorkDetail,EditQuartz} from "../../../api"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import QuartzAdd from "./QuartzAdd"

  export default {
    name: "QuartzList",
    components:{ISimpleLeftRightRow,ISimpleSearch,QuartzAdd},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        // 搜索条件
        search:"",
        quartzs: [],
        columns1: [
          {
            title: '任务名称',
            key: 'task_name',
            width: 200,
          },
          {
            title: '任务类型',
            key: 'task_type',
            width: 120,
          },
          {
            title: 'cron表达式',
            key: 'cron_str',
            width: 100,
          },
          {
            title: '启用状态',
            key: 'enable',
            width: 100,
          },
          {
            title: '最后修改人',
            key: 'last_updated_by',
            width: 180,
          },
          {
            title: '最后修改时间',
            key: 'last_updated_time',
            width: 180,
            render: (h,params)=>{
              return h('div',
                formatDate(new Date(params.row.last_updated_time),'yyyy-MM-dd hh:mm')
              )
            }
          },
          {
            title: '操作',
            key: 'operate',
            width: 280,
            fixed: 'right',
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.editQuartz(this.quartzs[params.index]['task_name'], "start");
                    }
                  }
                }, '启用'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.editQuartz(this.quartzs[params.index]['task_name'], "stop");
                    }
                  }
                }, '停用'),
                h('Button', {
                  props: {
                    type: 'warning',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.editQuartz(this.quartzs[params.index]['task_name'], "delete");
                    }
                  }
                }, '删除'),
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: async () => {
                      const result = await QueryWorkDetail(this.$route.query.work_id);
                      if(result.status == "SUCCESS"){
                        this.$router.push({ path: '/iwork/runLogList', query: { work_id: result.work.id }});
                      }

                    }
                  }
                }, '调度记录'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      refreshQuartzList:async function () {
        const result = await QuartzList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.quartzs = result.quartzs;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshQuartzList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshQuartzList();
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshQuartzList();
      },
      editQuartz:async function (task_name, operate){
        const result = await EditQuartz(task_name, operate);
        if(result.status == "SUCCESS"){
          this.refreshQuartzList();
        }else{
          this.$Message.error(result.errorMsg);
        }
      }
    },
    mounted: function () {
      this.refreshQuartzList();
    },
  }
</script>

<style scoped>

</style>
