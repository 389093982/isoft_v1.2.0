<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <ChooseCourseType/>

    <ToolBox :placement_name="GLOBAL.element_host_toolbox_list"/>

    <Row style="margin-bottom: 20px;margin: 10px auto;padding: 3px;">
      <Col span="18">
        <!-- 热门推荐区域 -->
        <HotRecommend/>
      </Col>
      <Col span="6">
        <ToolBox :placement_name="GLOBAL.element_host_toolbox_list2"/>
        <NewsNotice/>
        <HotIssue/>

        <HotArticleRecommend/>
      </Col>
    </Row>

    <ToolBox :placement_name="GLOBAL.element_host_toolbox_list"/>
    <ToolBox :placement_name="GLOBAL.element_host_toolbox_list"/>

    <!-- 热门编程语言 -->
    <HotLanguage/>
    <!-- 热门开源项目 -->
    <HotProject/>

    <IBeautifulTabLink/>
    <Find/>
    <Fun/>

  </div>
</template>

<script>
  import ChooseCourseType from "./CourseType/ChooseCourseType"
  import HotRecommend from "./Course/HotRecommend"
  import NewsNotice from "./Course/NewsNotice"
  import HotIssue from "./Course/HotIssue"
  import HotProject from "./Hot/HotProject"
  import ToolBox from "../Background/CMS/ToolBox"
  import HotLanguage from "./Hot/HotLanguage"
  import IBeautifulTabLink from "../Common/link/IBeautifulTabLink"
  import HotArticleRecommend from "./Course/HotArticleRecommend"
  import Fun from "./Fun"
  import Find from "./Find";

  export default {
    name: "Index",
    components:{
      Find, ChooseCourseType,HotRecommend,NewsNotice,HotIssue,HotProject,ToolBox,
      HotLanguage,IBeautifulTabLink,HotArticleRecommend,Fun},
  }
</script>

<style scoped>
  /* 引入公共样式库 */
  @import "../../../static/css/common.css";

</style>
