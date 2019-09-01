<template>
  <span v-if="course">
    <p style="color: #d6241e;">
      浏览量：{{course.watch_number}}
      课程分数：<Rate disabled show-text allow-half :v-model="course.score"/>
      &nbsp; <span>如何提升得分？</span> &nbsp;
      <IBeautifulLink2>我的课程空间</IBeautifulLink2>
    </p>
    <p>
      课程名称：{{course.course_name}}
      <router-link to="/ilearning/index" style="margin-left: 30px;">搜索同类资源</router-link>
    </p>
    <p>作者：{{course.course_author}}</p>
    <p>课程类型：{{course.course_type}}</p>
    <p>课程子类型：{{course.course_sub_type}}</p>
    <p>课程简介：{{course.course_short_desc}}</p>
    <p>课程集数：{{course.course_number}}</p>
    <p>课程更新状态：{{course.course_status}}</p>
  </span>
</template>

<script>
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";
  export default {
    name: "CourseMeta",
    components: {IBeautifulLink2},
    props:{
      course:{
        type:Object,
        default:null,
      }
    }
  }
</script>

<style scoped>

</style>
