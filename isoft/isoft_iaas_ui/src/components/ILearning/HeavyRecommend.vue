<template>
  <div style="width: 300px;height: 450px;position: absolute;top: 90px;left: 65px;z-index: 10;">
    <a class="section01">
      <p class="title">重磅推荐</p>
      <p class="content">重磅推荐1111111111111111</p>
    </a>
    <a class="section01">
      <p class="title">重磅推荐</p>
      <p class="content">重磅推荐1111111111111111</p>
    </a>
    <a class="section01">
      <p class="title">重磅推荐</p>
      <p class="content">重磅推荐1111111111111111</p>
    </a>
    <a class="section01">
      <p class="title">重磅推荐</p>
      <p class="content">重磅推荐1111111111111111</p>
    </a>
    <a class="section01">
      <p class="title">重磅推荐</p>
      <p class="content">重磅推荐1111111111111111</p>
    </a>
  </div>
</template>

<script>
  export default {
    name: "HeavyRecommend"
  }
</script>

<style scoped>
  .section01{
    height: 90px;
    color: white;
    padding: 20px;
    font-size: 16px;
    background-color: rgba(0,0,0,0.2);
    display: block;
    box-sizing: border-box;
  }
  .section01:hover{
    background-color: red;
    background-color: rgba(0,0,0,0.6);
    border-left: 1px solid yellow;
  }
  .section01:hover .title{
    color: yellow;
    font-weight: bold;
  }
</style>
