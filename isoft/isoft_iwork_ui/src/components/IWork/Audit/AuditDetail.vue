<template>
  <div>
    <Tabs :animated="false" name="tab_level_1">
      <TabPane label="配置数据来源" tab="tab_level_1">
        <AuditDetailSourceEdit/>
      </TabPane>
      <TabPane label="配置审核场景" tab="tab_level_1">
        <AuditDetailTargetEdit/>
      </TabPane>
      <TabPane label="数据审核" tab="tab_level_1">
        <AuditDetailHandle/>
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
  }
</script>

<style scoped>

</style>
