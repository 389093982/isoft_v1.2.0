<template>
  <div>
    <GraphicLink :placement_name="GLOBAL.placement_activity_list"/>
  </div>
</template>

<script>
  import GraphicLink from "../Background/CMS/GraphicLink";
  export default {
    name: "ActivityList",
    components: {GraphicLink}
  }
</script>

<style scoped>

</style>
