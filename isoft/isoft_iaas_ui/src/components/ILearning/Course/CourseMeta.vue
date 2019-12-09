<template>
  <span v-if="course">
    <p style="color: #d6241e;">
      浏览量：{{course.watch_number}}
      课程分数：<Rate disabled show-text allow-half :v-model="course.score"/>
      &nbsp; <span>如何提升得分？</span>
    </p>
    <p>
      课程名称：{{course.course_name}}
      <span class="isoft_hover_red" @click="$router.push({path:'/ilearning/index'})"
            style="margin-left: 30px;">搜索同类资源</span>
      <span class="isoft_hover_red" @click="$router.push($router.push({ path: '/ilearning/mine/course_space'}))"
            style="margin-left: 30px;">我的课程空间</span>
      <span class="isoft_hover_red" @click=""
            style="margin-left: 30px;">开课流程</span>
    </p>
    <p>作者：{{course.course_author}}</p>
    <p>课程类型：{{course.course_type}}</p>
    <p>课程子类型：{{course.course_sub_type}}</p>
    <p>课程简介：{{course.course_short_desc}}</p>
    <p>课程集数：{{course.course_number}}</p>
  </span>
</template>

<script>
  import IBeautifulLink from "../../Common/link/IBeautifulLink";
  export default {
    name: "CourseMeta",
    components: {IBeautifulLink},
    props:{
      course:{
        type:Object,
        default:null,
      }
    }
  }
</script>

<style scoped>
  @import "../../../assets/css/isoft_common.css";

</style>
