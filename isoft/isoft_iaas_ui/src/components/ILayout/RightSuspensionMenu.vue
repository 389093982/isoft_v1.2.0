<template>
  <div class="rightSuspension">
    <div class="itemArea">问题反馈</div>
    <div class="itemArea">问题反馈</div>
    <div class="itemArea">问题反馈</div>
    <div class="itemArea">问题反馈</div>
    <div class="itemArea">问题反馈</div>
  </div>
</template>

<script>
  export default {
    name: "RightSuspensionMenu"
  }
</script>

<style scoped>
  .rightSuspension {
    position: fixed;
    width: 50px;
    min-height: 200px;
    right: -40px;
    top: 30%;
    z-index: 999;
    transition: right ease-in 200ms;
  }
  .rightSuspension:hover {
    right: 0px;
  }
  .itemArea{
    background-color: rgba(234, 234, 234, 0.99);
    height: 50px;
    margin: 10px 0;
    color: white;
  }
  .itemArea:hover{
    background-color: red;
  }
</style>
