<template>
  <div id="header">
    <div class="wellogin">欢迎登陆</div>
    <div style="position: absolute;top: 40px;right: 100px;">
      <b id="question"></b>
      <span style="font-size: 14px;"><a href="#">加入我们</a> | <a @click="$router.push({path:'/ilearning/advise'})">我要吐槽 / 建议</a></span>
    </div>
  </div>
</template>

<script>
    export default {
      name: "LoginHeader"
    }
</script>

<style scoped>
  #header {
    background-color: rgb(255, 255, 255);
    text-align:center;
    height:70px;
    padding:5px;
  }
  .wellogin{
    display: block;
    width: 110px;
    height: 40px;
    position: absolute;
    top: 20px;
    left: 50px;
    font-size: 20px;
  }
  #question{
    display: inline-block;
    width: 18px;
    height: 14px;
    margin: 0 5px;
    background: url('../../../assets/sso/question_icon.png') no-repeat;
    overflow: hidden;
    vertical-align: middle;
  }
  a {
    color: #666;
    text-decoration: none;
  }
  a:hover {
    color: #E4393C;
    text-decoration: underline;
  }
</style>
