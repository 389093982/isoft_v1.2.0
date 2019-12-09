<template>
  <div style="width: 265px;height: 450px;position: absolute;top: 104px;left: 110px;z-index: 10;">
    <a class="section01">
      <p class="title">编程语言</p>
      <p class="content">汇集各种流行编程语言精华</p>
    </a>
    <a class="section01">
      <p class="title">流行框架</p>
      <p class="content">深度剖析多种流行框架源码</p>
    </a>
    <a class="section01">
      <p class="title">名师荟萃</p>
      <p class="content">重磅推荐各领域名家大师进行授课</p>
    </a>
    <a class="section01">
      <p class="title">技术交流</p>
      <p class="content">全网用户互动交流的集结地</p>
    </a>
    <a class="section01">
      <p class="title">能力计划</p>
      <p class="content">各领域多种多样的途径提升推广自己</p>
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
    height: 20%;
    color: white;
    padding: 20px;
    background-color: rgba(0,0,0,0.2);
    display: block;
    box-sizing: border-box;
  }
  .title{
    font-size: 16px;
  }
  .content{
    font-size: 14px;
  }
  .section01:hover{
    background-color: rgba(0,0,0,0.6);
  }
  .section01:hover .title{
    color: yellow;
    font-weight: bold;
  }
  .section01::before{
    content: '';
    position: absolute;
    margin-left: -20px;
    margin-top: 10%;
    border-left: 3px solid yellow;
    transition: all 300ms ease-in-out;
    height: 0;
  }

  .section01:hover::before{
    margin-top: -20px;
    height: 20%;
  }
</style>
