<template>
  <div>
    <h4>内容审核系统</h4>
    请选择数据源：
    <Select v-model="select_resource_name" style="width:300px;margin: 10px 0;">
      <Option v-for="(resource,index) in resources" :value="resource.resource_name" :key="resource.resource_name">
        {{ resource.resource_name }} ~ {{ resource.resource_dsn }}
      </Option>
    </Select>

    <Button type="success" @click="showAuditEdit = true">新增审核任务</Button>
    <Modal
      v-model="showAuditEdit"
      title="编辑审核任务"
      :mask-closable="false"
      :footer-hide="true">
      <AuditTaskEdit @handleSucess="handleAuditEdit"/>
    </Modal>


    <Table stripe :columns="columns1" :data="tasks"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>

  </div>
</template>

<script>
  import {GetAllResource,QueryPageAuditTask} from "../../../api"
  import AuditTaskEdit from "./AuditTaskEdit"

    export default {
      name: "AuditTaskList",
      components:{AuditTaskEdit},
      data(){
        return {
          showAuditEdit:false,
          select_resource_name:'',
          resources:[],
          total:0,    // 总数量
          offset:10,  // 每页数据量
          current_page: 1,
          columns1: [
            {
              title: '任务名称',
              key: 'task_name'
            },
            {
              title: '任务描述',
              key: 'task_desc'
            },
            {
              title: '操作',
              key: 'operate',
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
                        this.$router.push({ path: '/iwork/audit_detail', query: { task_name: this.tasks[params.index].task_name }});
                      }
                    }
                  }, '编辑任务详情'),
                ]);
              }
            },
          ],
          tasks: []
        }
      },
      methods:{
        handleChange(page){
          this.current_page = page;
          this.refreshAllAuditTask();
        },
        handlePageSizeChange(pageSize){
          this.offset = pageSize;
          this.refreshAllAuditTask();
        },
        refreshAllAuditTask:async function (){
          const result = await QueryPageAuditTask(this.offset, this.current_page);
          if(result.status == "SUCCESS"){
            this.tasks = result.tasks;
            this.total = result.paginator.totalcount;
          }
        },
        refreshAllResource:async function(){
          const result = await GetAllResource("db");
          if(result.status == "SUCCESS"){
            this.resources = result.resources;
          }
        },
        handleAuditEdit:function (task_name, task_desc) {
          this.showAuditEdit = false;
          this.refreshAllAuditTask();
        }
      },
      mounted (){
        this.refreshAllResource();
        this.refreshAllAuditTask();
      }
    }
</script>

<style scoped>

</style>
