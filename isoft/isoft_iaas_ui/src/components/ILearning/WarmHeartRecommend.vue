<template>
  <div class="isoft_bg_white isoft_pd10">
    <div class="header">为您精选</div>
    <Row>
      <Col span="12" style="height:365px;">
        <video ref="video" style="width:100%;height:100%;object-fit: fill" controls autoplay="autoplay" src="http://localhost:8086/api/files/75dd3df4-a3db-47ce-948e-4cacd6348d88">
          <source type="video/mp4">
          <source type="video/ogg">
          您的浏览器不支持Video标签。
        </video>
      </Col>
      <Col span="12" style="padding-left: 10px;">
        <div class="linearTransitionBg" style="height: 40px;padding:7px;margin-bottom:10px;background-color: rgba(228,228,228,0.4);">
          <a class="hovered hvr-grow hoverLinkColor mr5">热门</a>
          <a class="hovered hvr-grow hoverLinkColor mr5">最新</a>
          <a class="hovered hvr-grow hoverLinkColor mr5">高评</a>
          <a class="hovered hvr-grow hoverLinkColor mr5">特色</a>
        </div>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
        <HoverBigImg class="hoverBorderShadow" width="144px" height="98px" src-img="http://localhost:8086/api/files/510825.gif" style="float: left;margin: 0 10px 10px 0;"/>
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

  .linearTransitionBg{
    width:450px;
    transition: all 5s ease;
  }

  .linearTransitionBg:hover{
    width:250px;
    transition: all 5s ease;
  }
</style>
