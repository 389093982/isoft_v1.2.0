<template>
  <div>
    <h3 style="text-align: center;">猜你喜欢</h3>
    <div v-show="showContent">
      <Row>
        <Col span="12">
          <ToolBox :placement_name="GLOBAL.element_host_app_recommand"/>
        </Col>
        <Col span="12">
          <ToolBox :placement_name="GLOBAL.element_host_app_recommand"/>
        </Col>
      </Row>
      <Row>
        <Col span="12">
          <ToolBox :placement_name="GLOBAL.element_host_app_recommand"/>
        </Col>
        <Col span="12">
          <ToolBox :placement_name="GLOBAL.element_host_app_recommand"/>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import ToolBox from "../Background/CMS/ToolBox";
  export default {
    name: "GussYouLike",
    components: {ToolBox},
    data(){
      return {

      }
    },
    props:{
      showContent:{
        type:Boolean,
        default:true,
      }
    }
  }
</script>

<style scoped>

</style>
