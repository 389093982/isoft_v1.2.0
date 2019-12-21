<template>
  <div>
    <div class="isoft_bg_white" style="padding: 10px 10px 0 10px;">
      <Row>
        <Col span="12" style="padding: 10px 0 0 20px;">
          <span style="color: #c3cbd6;font-size: 18px;">物联网时代，让交流更直接</span>
          <IBeautifulLink style="margin-left: 20px;">加入我们</IBeautifulLink>
        </Col>
        <Col span="12">
          <ISearch @submitFunc="searchFunc"/>
        </Col>
      </Row>

      <Row style="margin: 5px 0 5px 0;padding: 5px 15px 0px 15px;font-size: 12px;">
        <Col span="12" style="text-align: left;">
          <a @click="$router.push({path:'/ilearning/index'})" class="hovered hvr-grow hoverLinkColor mr5">返回首页</a>
        </Col>
        <Col span="12" style="text-align: right;">
          <a @click="$router.push({path:'/ilearning/about'})" class="hovered hvr-grow hoverLinkColor mr5">关于ILearning</a>
          <a @click="$router.push({path:'/user/guide'})" class="hovered hvr-grow hoverLinkColor mr5">站点引导</a>
          <a @click="$router.push({path:'/advertisement/apply'})" class="hovered hvr-grow hoverLinkColor mr5">广告位招租</a>
          <a @click="$router.push({path:'/user/mine/detail',query:{username:'mine'}})" class="hovered hvr-grow hoverLinkColor mr5">个人中心</a>
          <a @click="$router.push({ path: '/ilearning/mine/course_space'})" class="hovered hvr-grow hoverLinkColor">我的课程空间</a>
        </Col>
      </Row>
    </div>

    <div class="isoft_bg_white isoft_pd10 isoft_bordertop_red" style="margin-top: 5px;">
      <IBeautifulCard title="课程天地">
        <div slot="content" style="padding: 5px;">
          <div>
            <div style="border-bottom: 2px solid #edf1f2;">
              <a href="javascript:;" @click="showCourseType=true" style="color: red;">热门课程推荐</a>
              <a href="javascript:;" @click="showCourseType=!showCourseType" style="color: red;float: right;">
                <IBeautifulLink style="font-size: 14px;"> 更多 </IBeautifulLink>
              </a>
            </div>
            <div>
              <HotCourseType :placement_name="GLOBAL.placement_host_course_type_carousel" v-show="showCourseType===true" @chooseCourseType="chooseCourseType"/>
              <TotalCourseType v-show="showCourseType===false" @chooseCourseType="chooseCourseType"/>
            </div>
          </div>
        </div>

        <div slot="header_right">
          <IBeautifulLink style="font-size: 14px;" @onclick="$router.push({ path: '/ilearning/mine/course_space'})"> 我的课程空间 </IBeautifulLink>
        </div>
      </IBeautifulCard>
    </div>

  </div>
</template>

<script>
  import HotCourseType from "./HotCourseType"
  import TotalCourseType from "./TotalCourseType"
  import ISearch from "../../Common/search/ISearch"
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import IBeautifulLink from "../../Common/link/IBeautifulLink"

  export default {
    name: "ChooseCourseType",
    components:{ISearch,HotCourseType,TotalCourseType,IBeautifulCard,IBeautifulLink},
    data(){
      return {
        showCourseType:true,
      }
    },
    methods: {
      searchFunc:function (data) {
        this.$router.push({ path: '/ilearning/course_search', query: { search: data }});
      },
      chooseCourseType:function (course_type, course_sub_type) {
        // params是路由的一部分
        // query是拼接在url后面的参数
        // 由于动态路由也是传递params的,所以在 this.$router.push() 方法中path不能和params一起使用,否则params将无效.需要用name来指定页面
        this.$router.push({ path: '/ilearning/course_search', query: { search: course_sub_type }});
      },
      toggle:function (data) {
        alert(data);
      }
    }
  }
</script>

<style scoped>
  @import "../../../assets/css/isoft_common.css";

</style>
