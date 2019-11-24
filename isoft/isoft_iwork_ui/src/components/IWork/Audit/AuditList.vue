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
      <AuditEdit @handleSucesss="handleAuditEdit"/>
    </Modal>


    <Table stripe :columns="columns1" :data="tasks"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>

  </div>
</template>

<script>
  import {GetAllResource} from "../../../api"
  import AuditEdit from "./AuditEdit"

    export default {
      name: "AuditList",
      components:{AuditEdit},
      data(){
        return {
          showAuditEdit:false,
          select_resource_name:'',
          resources:[],
          total:0,    // 总数量
          offset:10,  // 每页数据量
          columns1: [
            {
              title: 'id',
              key: 'task_name'
            },
            {
              title: '条目名称',
              key: 'task_desc'
            },
            {
              title: '审核源',
              key: 'source'
            }
          ],
          tasks: [
            {
              task_name: 'John Brown',
              task_desc: 'New York No. 1 Lake Park',
            },
            {
              task_name: 'John Brown',
              task_desc: 'New York No. 1 Lake Park',
            },
            {
              task_name: 'John Brown',
              task_desc: 'New York No. 1 Lake Park',
            },
          ]
        }
      },
      methods:{
        handleChange:function(){

        },
        handlePageSizeChange:function(){

        },
        async refreshAllResource (){
          const result = await GetAllResource("db");
          if(result.status == "SUCCESS"){
            this.resources = result.resources;
          }
        },
        handleAuditEdit:function (task_name, task_desc) {
          this.showAuditEdit = false;
          this.tasks.push({'task_name': task_name, "task_desc":task_desc});
          this.total = this.tasks.length;
        }
      },
      mounted (){
        this.refreshAllResource();
      }
    }
</script>

<style scoped>

</style>
