<template>
  <div>
    <Tabs :animated="false" name="tab_level_1">
      <TabPane label="配置任务源数据" tab="tab_level_1">
        <AuditDetailSourceEdit/>
      </TabPane>
      <TabPane label="配置任务审核信息" tab="tab_level_1">
        <AuditDetailTargetEdit/>
      </TabPane>
      <TabPane label="配置映射关系" tab="tab_level_1">

      </TabPane>

    </Tabs>
  </div>
</template>

<script>
  import AuditDetailSourceEdit from "./AuditDetailSourceEdit";
  import AuditDetailTargetEdit from "./AuditDetailTargetEdit";

  export default {
    name: "AuditDetail",
    components: {AuditDetailTargetEdit, AuditDetailSourceEdit},
  }
</script>

<style scoped>

</style>
