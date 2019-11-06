<template>
  <div>
    <div>
      为您精选
    </div>
    <Row>
      <Col span="12">
        <video ref="video" width="585px" height="365px" controls autoplay="autoplay">
          <source type="video/mp4">
          <source type="video/ogg">
          您的浏览器不支持Video标签。
        </video>
      </Col>
      <Col span="12">
        <HoverBigImg width="154px" height="86px"/>
        <HoverBigImg width="160px" height="90px"/>
        <HoverBigImg width="154px" height="86px"/>
        <HoverBigImg width="154px" height="86px"/>
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

</style>
