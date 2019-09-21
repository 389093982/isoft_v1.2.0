<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <ChooseCourseType/>

    <Row style="margin-bottom: 20px;margin: 10px auto;padding: 3px;">
      <Col span="18">
        <!-- 热门推荐区域 -->
        <HotRecommend/>
      </Col>
      <Col span="6">
        <ToolBox :placement_name="GLOBAL.element_host_toolbox_list2"/>
        <NewsNotice :placement_name="GLOBAL.Index_news_list"/>
        <NewsNotice :placement_name="GLOBAL.Index_news_list"/>
        <NewsNotice :placement_name="GLOBAL.Index_news_list"/>

      </Col>
    </Row>

    <ToolBox :placement_name="GLOBAL.element_host_toolbox_list3"/>
    <ToolBox :placement_name="GLOBAL.element_host_toolbox_list4"/>

    <!-- 热门开源项目 -->
    <IHotRecommand :placement_name="GLOBAL.hot_project_recommod"/>

    <IBeautifulTabLink/>

    <HorizontalLinks :placement_name="GLOBAL.element_want_to_find"/>

    <Fun/>

  </div>
</template>

<script>
  import ChooseCourseType from "./CourseType/ChooseCourseType"
  import HotRecommend from "./Course/HotRecommend"
  import NewsNotice from "./Course/NewsNotice"
  import HotIssue from "./Course/HotIssue"
  import ToolBox from "../Background/CMS/ToolBox"
  import IBeautifulTabLink from "../Common/link/IBeautifulTabLink"
  import Fun from "./Fun"
  import HorizontalLinks from "../Elementviewers/HorizontalLinks"
  import IHotRecommand from "../Common/recommend/IHotRecommand"

  export default {
    name: "Index",
    components:{
      IHotRecommand, HorizontalLinks, ChooseCourseType,HotRecommend,NewsNotice,HotIssue,ToolBox, IBeautifulTabLink,Fun},
  }
</script>

<style scoped>
  /* 引入公共样式库 */
  @import "../../../static/css/common.css";

</style>
