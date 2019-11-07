<template>
  <div>
    <div class="header">为您精选</div>
    <Row>
      <Col span="12">
        <video ref="video" width="585px" height="365px" controls autoplay="autoplay">
          <source type="video/mp4">
          <source type="video/ogg">
          您的浏览器不支持Video标签。
        </video>
      </Col>
      <Col span="12">
        <p style="height: 50px">AAAAAAAA BBBBBBBBBB CCCCCCCCC DDDDDDDDD</p>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg width="184px" height="98px" style="float: left;margin: 0 10px 10px 0;"/>
      </Col>
    </Row>

  </div>
</template>

<script>
  import HoverBigImg from "../Common/img/HoverBigImg";
  export default {
    name: "WarmHeartRecommend",
    components: {HoverBigImg}
  }
</script>

<style scoped>
  .header{
    margin: 0 10px 10px 10px;
    position: relative;
    height: 40px;
    color: #111;
    font-size: 20px;
    font-weight: 400;
    line-height: 40px;
    white-space: nowrap;
  }
</style>
