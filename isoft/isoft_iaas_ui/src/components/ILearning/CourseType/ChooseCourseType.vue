<template>
  <div>
    <div style="margin-bottom: 5px;">
      <ISearch @submitFunc="searchFunc"/>
      <div style="text-align: right;">
        <IBeautifulLink2 style="margin-top: 5px;" @onclick="$router.push({path:'/user/guide'})">站点引导</IBeautifulLink2>
        <IBeautifulLink2 style="margin-top: 5px;">广告位招租</IBeautifulLink2>

        <IBeautifulLink2 style="margin-top: 5px;"
                         @onclick="$router.push({path:'/user/mine/detail',query:{username:'mine'}})">个人中心</IBeautifulLink2>
        <IBeautifulLink2 style="margin-top: 5px;">我的课程空间</IBeautifulLink2>
      </div>
    </div>

    <IBeautifulCard title="课程天地">
      <div slot="content" style="padding: 20px;">
        <div>
          <div style="border-bottom: 2px solid #edf1f2;">
            <a href="javascript:;" @click="showCourseType=true" style="color: red;">热门课程推荐</a>
            <a href="javascript:;" @click="showCourseType=!showCourseType" style="color: red;float: right;">
              <IBeautifulLink2 style="font-size: 14px;"> 更多 </IBeautifulLink2>
            </a>
          </div>
          <div>
            <HotCourseType :placement_name="GLOBAL.element_host_course_type_carousel" v-show="showCourseType===true" @chooseCourseType="chooseCourseType"/>
            <TotalCourseType v-show="showCourseType===false" @chooseCourseType="chooseCourseType"/>
          </div>
        </div>
      </div>


      <div slot="header_right">
        <IBeautifulLink2 style="font-size: 14px;" @onclick="$router.push({ path: '/ilearning/mine/course_space'})"> 我的课程空间 </IBeautifulLink2>
      </div>
    </IBeautifulCard>
  </div>
</template>

<script>
  import HotCourseType from "./HotCourseType"
  import TotalCourseType from "./TotalCourseType"
  import ISearch from "../../Common/search/ISearch"
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2"

  export default {
    name: "ChooseCourseType",
    components:{ISearch,HotCourseType,TotalCourseType,IBeautifulCard,IBeautifulLink2},
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

</style>
