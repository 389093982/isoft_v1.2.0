<template>
  <div>
    <Tabs :animated="false">
      <TabPane label="数据来源">
        请选择数据源：
        <Select v-model="taskDetail.resource_name" style="width:300px;margin: 10px 0;">
          <Option v-for="(resource,index) in resources" :value="resource.resource_name" :key="resource.resource_name">
            {{ resource.resource_name }} ~ {{ resource.resource_dsn }}
          </Option>
        </Select>

        <br/>

        请输入查询 Sql
        <Input type="textarea" :rows="5" v-model="taskDetail.query_sql"/>

        <Button type="success" style="margin-top: 10px;" @click="editAuditTaskSource">提交</Button>
      </TabPane>
      <TabPane label="审核步骤">
        审核步骤
      </TabPane>
      <TabPane label="更新审核状态">
        更新审核状态
      </TabPane>
    </Tabs>
  </div>
</template>

<script>
  import {GetAllResource, EditAuditTaskSource,QueryTaskDetail} from "../../../api"

  export default {
    name: "AuditDetail",
    data(){
      return {
        resources:[],
        taskDetail:{
          resource_name:'',
          query_sql:'',
        },
      }
    },
    methods:{
      editAuditTaskSource:async function(){
        const result = await EditAuditTaskSource(this.$route.query.task_name, this.taskDetail.resource_name, this.taskDetail.query_sql);
        if(result.status == "SUCCESS"){
          this.$Message.success("保存成功！");
          this.refreshAuditDetail();
        }else{
          this.$Message.error("配置错误!");
        }
      },
      refreshAllResource:async function(){
        const result = await GetAllResource("db");
        if(result.status == "SUCCESS"){
          this.resources = result.resources;
        }
      },
      refreshAuditDetail:async function () {
        const result = await QueryTaskDetail(this.$route.query.task_name);
        if(result.status == "SUCCESS"){
          this.taskDetail = result.taskDetail;
        }
      }
    },
    mounted(){
      this.refreshAllResource();
      this.refreshAuditDetail();
    }
  }
</script>

<style scoped>

</style>
