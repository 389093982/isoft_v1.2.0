<template>
  <div class="leftSuspension">
    右侧悬浮菜单
  </div>
</template>

<script>
  export default {
    name: "RightSuspensionMenu"
  }
</script>

<style scoped>
  .leftSuspension {
    position: fixed;
    width: 60px;
    min-height: 200px;
    background-color: rgba(234, 234, 234, 0.99);
    right: -40px;
    top: 40%;
    z-index: 999;
    transition: right ease-in 200ms;
  }
  .leftSuspension:hover {
    right: 0px;
  }
</style>
