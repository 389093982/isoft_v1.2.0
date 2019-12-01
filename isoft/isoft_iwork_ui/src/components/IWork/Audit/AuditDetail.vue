<template>
  <div>
    <Tabs :animated="false" name="tab_level_1" :value="active_index" @on-click="currentTabChanged">
      <TabPane label="配置数据来源" tab="tab_level_1" name="active_index_1">
        <AuditDetailSourceEdit v-if="active_index == 'active_index_1'"/>
      </TabPane>
      <TabPane label="配置审核场景" tab="tab_level_1" name="active_index_2">
        <AuditDetailTargetEdit v-if="active_index == 'active_index_2'"/>
      </TabPane>
      <TabPane label="数据审核" tab="tab_level_1" name="active_index_3">
        <AuditDetailHandle v-if="active_index == 'active_index_3'"/>
      </TabPane>

    </Tabs>
  </div>
</template>

<script>
  import AuditDetailSourceEdit from "./AuditDetailSourceEdit";
  import AuditDetailTargetEdit from "./AuditDetailTargetEdit";
  import AuditDetailHandle from "./AuditDetailHandle";

  export default {
    name: "AuditDetail",
    components: {AuditDetailHandle, AuditDetailTargetEdit, AuditDetailSourceEdit},
    data(){
      return {
        active_index: 'active_index_1',
      }
    },
    methods:{
      currentTabChanged:function (name) {
        this.active_index = name;
      },
    }
  }
</script>

<style scoped>

</style>
