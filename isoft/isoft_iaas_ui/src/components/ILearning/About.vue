<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <HeavyRecommend/>

    <div class="section01_bg">
      <WhyChooseUs style="margin-left: 350px;"/>
    </div>

    <Announce/>

    <Row :gutter="10">
      <Col span="6">
        <div class="hoverBorderShadow" style="height: 100px;padding: 20px;">
          <Row>
            <Col span="8">
              <img src="../../assets/icon_jy.png" style="width: 52px;height: 52px;"/>
            </Col>
            <Col span="16">
              <h4>好评如潮</h4>
              <p>好评如潮好评如潮</p>
            </Col>
          </Row>
        </div>
      </Col>
      <Col span="6">
        <div class="hoverBorderShadow" style="height: 100px;padding: 20px;">
          <Row>
            <Col span="8">
              <img src="../../assets/icon_jy.png" style="width: 52px;height: 52px;"/>
            </Col>
            <Col span="16">
              <h4>好评如潮</h4>
              <p>好评如潮好评如潮</p>
            </Col>
          </Row>
        </div>
      </Col>
      <Col span="6">
        <div class="hoverBorderShadow" style="height: 100px;padding: 20px;">
          <Row>
            <Col span="8">
              <img src="../../assets/icon_jy.png" style="width: 52px;height: 52px;"/>
            </Col>
            <Col span="16">
              <h4>好评如潮</h4>
              <p>好评如潮好评如潮</p>
            </Col>
          </Row>
        </div>
      </Col>
      <Col span="6">
        <div class="hoverBorderShadow" style="height: 100px;padding: 20px;">
          <Row>
            <Col span="8">
              <img src="../../assets/icon_jy.png" style="width: 52px;height: 52px;"/>
            </Col>
            <Col span="16">
              <h4>好评如潮</h4>
              <p>好评如潮好评如潮</p>
            </Col>
          </Row>
        </div>
      </Col>
    </Row>

    <Row :gutter="10">
      <Col span="16">
        <ShowUserExperience/>
      </Col>
      <Col span="8">
        <Plan/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import HeavyRecommend from "./HeavyRecommend";
  import WhyChooseUs from "./WhyChooseUs";
  import Announce from "./Announce";
  import ShowUserExperience from "./ShowUserExperience";
  import Plan from "./Plan";

  export default {
    name: "About",
    components:{HeavyRecommend,WhyChooseUs,Announce,ShowUserExperience,Plan},
  }
</script>

<style scoped>
  /* 引入公共样式库 */


  .section01_bg{
    height: 450px;
    background: url(../../assets/1.jpg) no-repeat;
    background-size: 100% 100%;
  }
</style>
