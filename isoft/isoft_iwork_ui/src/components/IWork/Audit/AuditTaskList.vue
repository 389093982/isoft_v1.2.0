<template>
  <div>
    <h2 style="text-align: center;" class="press">智能内容审核系统</h2>
    <Button type="success" size="small" style="margin: 10px 0;" @click="showAuditEdit = true">新增审核任务</Button>
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
  import {QueryPageAuditTask,DeleteAuditTask} from "../../../api"
  import AuditTaskEdit from "./AuditTaskEdit"

    export default {
      name: "AuditTaskList",
      components:{AuditTaskEdit},
      data(){
        return {
          showAuditEdit:false,
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
                  }, '详情编辑'),
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
                        this.$Modal.confirm({
                          title: '删除',
                          content: '确认删除该条数据吗？请谨慎操作！',
                          onOk: () => {
                            this.deleteAuditTask(this.tasks[params.index].task_name);
                          },
                          onCancel: () => {
                            this.$Message.info('取消操作');
                          }
                        });
                      }
                    }
                  }, '删除'),
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
        deleteAuditTask:async function(task_name){
          const result = await DeleteAuditTask(task_name);
          if(result.status == "SUCCESS"){
            this.$Message.success("删除成功!");
            this.refreshAllAuditTask();
          }else{
            this.$Message.error("删除失败!");
          }
        },
        handleAuditEdit:function (task_name, task_desc) {
          this.showAuditEdit = false;
          this.refreshAllAuditTask();
        }
      },
      mounted (){
        this.refreshAllAuditTask();
      }
    }
</script>

<style scoped>
  .press {
    color: transparent;
    background-color : black;
    text-shadow : rgba(255,255,255,0.5) 0 5px 6px, rgba(255,255,255,0.2) 1px 3px 3px;
    -webkit-background-clip : text;
  }
</style>
