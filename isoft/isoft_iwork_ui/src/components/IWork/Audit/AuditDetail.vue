<template>
  <div>
    <Tabs :animated="false">
      <TabPane label="数据来源">
        请选择数据源：
        <Select v-model="select_resource_name" style="width:300px;margin: 10px 0;">
          <Option v-for="(resource,index) in resources" :value="resource.resource_name" :key="resource.resource_name">
            {{ resource.resource_name }} ~ {{ resource.resource_dsn }}
          </Option>
        </Select>

        <br/>

        请输入查询 Sql
        <Input type="textarea" :rows="5" v-model="querySql"/>

        <Button type="success" @click="editAuditTaskSource">提交</Button>
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
  import {GetAllResource, EditAuditTaskSource} from "../../../api"

  export default {
    name: "AuditDetail",
    data(){
      return {
        select_resource_name:'',
        resources:[],
        querySql:'',
      }
    },
    methods:{
      editAuditTaskSource:async function(){
        const result = await EditAuditTaskSource(this.$route.query.task_name, this.select_resource_name, this.querySql);
        if(result.status == "SUCCESS"){
          alert(JSON.stringify(result));
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
    },
    mounted(){
      this.refreshAllResource();
    }
  }
</script>

<style scoped>

</style>
