<template>
  <div class="leftSuspension">
    左侧悬浮菜单
  </div>
</template>

<script>
  export default {
    name: "LeftSuspensionMenu"
  }
</script>

<style scoped>
  .leftSuspension {
    position: fixed;
    width: 150px;
    min-height: 200px;
    background-color: red;
    left: 0px;
    top: 40%;
    z-index: 999;
  }
</style>
