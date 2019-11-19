<template>
  <div style="margin: 10px 0 10px 0;">
    <div class="container">
      <Row style="margin: 10px 0 10px 0;">
        <Col span="12" style="text-align: left;">
          <a @click="$router.push({path:'/ilearning/index'})" class="hovered hvr-grow hoverLinkColor mr5">返回首页</a>
        </Col>
        <Col span="12" style="text-align: right;">
          <a @click="$router.push({path:'/ilearning/about'})" class="hovered hvr-grow hoverLinkColor mr5">关于ILearning</a>
          <a @click="$router.push({path:'/user/guide'})" class="hovered hvr-grow hoverLinkColor mr5">站点引导</a>
          <a @click="" class="hovered hvr-grow hoverLinkColor mr5">广告位招租</a>
          <a @click="$router.push({path:'/user/mine/detail',query:{username:'mine'}})" class="hovered hvr-grow hoverLinkColor mr5">个人中心</a>
          <a @click="" class="hovered hvr-grow hoverLinkColor mr5">我的课程空间</a>
        </Col>
      </Row>
    </div>
    <Row style="margin: -60px 50px 0 50px;">
      <Col span="6">
        <div class="box">
          <span class="title">农林牧渔、卫生、科学研究</span>
          <a>农业</a>
          <a>食品</a>
          <a>医疗</a>
          <a>药业</a>
          <a>公共卫生</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <span class="title">农林牧渔、卫生、科学研究</span>
          <a>农业</a>
          <a>食品</a>
          <a>医疗</a>
          <a>药业</a>
          <a>公共卫生</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <span class="title">农林牧渔、卫生、科学研究</span>
          <a>农业</a>
          <a>食品</a>
          <a>医疗</a>
          <a>药业</a>
          <a>公共卫生</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
          <a>国土</a>
        </div>
      </Col>
      <Col span="12">
        <div class="box" style="margin: 0 10px 0 10px;">
          https://www.cnki.net/
        </div>
      </Col>
      <Col span="6">
        <div class="box">
          https://www.cnki.net/
        </div>
      </Col>
    </Row>
  </div>
</template>

<script>
  export default {
    name: "Nav"
  }
</script>

<style scoped>
  @import "../../../assets/css/isoft_common.css";

  .container{
    background: url("../../../assets/banner1.jpg");
    height: 350px;
  }
  .box{
    height: 500px;
    background-color: #fff;
    border: 1px solid #e4e8eb;
    border-top: 4px solid #7dacf0;
    border-radius: 4px;
    padding: 20px;
  }
  .box:hover{
    border-top: 4px solid #ff9c00;
  }
  .title{
    display: block;
    font-size: 15px;
    color: #000;
    font-weight: normal;
    line-height: 28px;
    margin: 15px 0 2px 0;
  }
  a{
    margin-right: 12px;
    font-size: 12px;
    color: #778193;
    line-height: 22px;
  }
</style>
