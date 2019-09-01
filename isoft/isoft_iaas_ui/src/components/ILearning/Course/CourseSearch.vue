<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <Row>
      <Col v-if="courses.length > 0" span="16">
        <div style="border-bottom: 1px solid #ccc;padding: 15px;border-right: 1px solid #ccc;padding: 15px;" v-for="course in courses">
          <Row>
            <Col span="8">
              <router-link :to="{path:'/ilearning/course_detail',query:{course_id:course.id}}">
                <h4>课程名称：{{course.course_name}}</h4>
                <img v-if="course.small_image" :src="course.small_image" height="120" width="200"/>
                <img v-else src="../../../assets/default.png" height="90px" width="120px"/>
              </router-link>
            </Col>
            <Col span="16">
              <CourseMeta :course="course"/>
            </Col>
          </Row>
        </div>
      </Col>
      <Col v-else span="16">
        未找到相应的搜索结果
      </Col>
      <Col span="8">
        <SiteDashBoard style="border: 1px solid #ccc;padding: 15px;margin-left: 5px;"/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {SearchCourseList} from "../../../api"
  import SiteDashBoard from "../Site/SiteDashBoard"
  import CourseMeta from "./CourseMeta";

  export default {
    name: "CourseSearch",
    components:{CourseMeta, SiteDashBoard},
    data(){
      return {
        courses:[],
      }
    },
    methods:{
      refreshCourseSearch:async function (search) {
        const result = await SearchCourseList(search);
        if(result.status=="SUCCESS"){
          this.courses = result.courses;
        }
      }
    },
    mounted:function () {
      // $route为当前router跳转对象里面可以获取name、path、query、params等
      // $router为VueRouter实例,想要导航到不同URL,则使用$router.push方法
      // 返回上一个history也是使用$router.go方法
      const search = this.$route.query.search;
      this.refreshCourseSearch(search);
    }
  }
</script>

<style scoped>
  a{
    color: black;
  }
  a:hover{
    color: red;
  }
</style>
