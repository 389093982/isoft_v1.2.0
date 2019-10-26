<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <WhyChooseUs/>
    <ShowUserExperience/>
    <ChooseCourseType/>

    <Row>
      <Col span="18">
        <!-- 热门推荐区域 -->
        <HotRecommend/>
      </Col>
      <Col span="6">
        <ToolBox :placement_name="GLOBAL.element_host_app_recommand"/>
        <NewsNotice :placement_name="GLOBAL.Index_news_list"/>
      </Col>
    </Row>

    <Row>
      <Col span="12">
        <ToolBox :placement_name="GLOBAL.element_host_toolbox_list3"/>
      </Col>
      <Col span="12">
        <ToolBox :placement_name="GLOBAL.element_host_toolbox_list4"/>
      </Col>
    </Row>

    <!-- 热门开源项目 -->
    <IHotRecommand :placement_name="GLOBAL.hot_project_recommod"/>

    <IBeautifulTabLink/>

    <HorizontalLinks :placement_name="GLOBAL.element_want_to_find"/>

    <div class="newspaper">
      人民网北京2月24日电 (记者 刘阳)国家发展改革委近日发出通知，决定自2月25日零时起将汽、柴油价格每吨分别提高300元和290元，折算到90号汽油和0号柴油（全国平均）每升零售价格分别提高0.22元和0.25元。
      此次国内成品油价格调整幅度，是按照现行国内成品油价格形成机制，根据国际市场油价变化情况确定的。去年11月16日国内成品油价格调整以来，受市场预期欧美经济复苏前景向好以及中东局势持续动荡等因素影响，国际市场原油价格先抑后扬，2月上旬WTI和布伦特原油期货价格再次回升至每桶95美元和115美元以上。虽然近两日价格有所回落，但国内油价挂钩的国际市场三种原油连续22个工作日移动平均价格上涨幅度已超过4%，达到国内成品油价格调整的边界条件。
      通知指出，这次成品油调价后，国家将按照已建立的补贴机制，继续对种粮农民、渔业（含远洋渔业）、林业、城市公交、农村道路客运（含岛际和农村水路客运）等给予补贴。同时，为保证市场物价基本稳定，防止连锁涨价，对与居民生活密切相关的铁路客运、城市公交、农村道路客运（含岛际和农村水路客运）价格不作调整。
      通知要求，中石油、中石化、中海油三大公司要组织好成品油生产和调运，保持合理库存，加强综合协调和应急调度，保障成品油供应。各级价格主管部门要加大市场监督检查力度，依法查处不执行国家价格政策，以及囤积居奇、造谣惑众、合谋涨价、搭车涨价等违法行为，维护正常市场秩序。
    </div>

    <Fun/>

    <div class="newspaper">
      人民网北京2月24日电 (记者 刘阳)国家发展改革委近日发出通知，决定自2月25日零时起将汽、柴油价格每吨分别提高300元和290元，折算到90号汽油和0号柴油（全国平均）每升零售价格分别提高0.22元和0.25元。
      此次国内成品油价格调整幅度，是按照现行国内成品油价格形成机制，根据国际市场油价变化情况确定的。去年11月16日国内成品油价格调整以来，受市场预期欧美经济复苏前景向好以及中东局势持续动荡等因素影响，国际市场原油价格先抑后扬，2月上旬WTI和布伦特原油期货价格再次回升至每桶95美元和115美元以上。虽然近两日价格有所回落，但国内油价挂钩的国际市场三种原油连续22个工作日移动平均价格上涨幅度已超过4%，达到国内成品油价格调整的边界条件。
      通知指出，这次成品油调价后，国家将按照已建立的补贴机制，继续对种粮农民、渔业（含远洋渔业）、林业、城市公交、农村道路客运（含岛际和农村水路客运）等给予补贴。同时，为保证市场物价基本稳定，防止连锁涨价，对与居民生活密切相关的铁路客运、城市公交、农村道路客运（含岛际和农村水路客运）价格不作调整。
      通知要求，中石油、中石化、中海油三大公司要组织好成品油生产和调运，保持合理库存，加强综合协调和应急调度，保障成品油供应。各级价格主管部门要加大市场监督检查力度，依法查处不执行国家价格政策，以及囤积居奇、造谣惑众、合谋涨价、搭车涨价等违法行为，维护正常市场秩序。
    </div>

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
  import WhyChooseUs from "./WhyChooseUs";
  import ShowUserExperience from "./ShowUserExperience";

  export default {
    name: "Index",
    components:{
      ShowUserExperience,
      WhyChooseUs,
      IHotRecommand, HorizontalLinks, ChooseCourseType,HotRecommend,NewsNotice,HotIssue,ToolBox, IBeautifulTabLink,Fun},
  }
</script>

<style scoped>
  /* 引入公共样式库 */
  @import "../../../static/css/common.css";
  @import "../../assets/css/isoft_common.css";

</style>
