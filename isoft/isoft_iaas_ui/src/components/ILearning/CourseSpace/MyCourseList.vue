<template>
  <div>
    <div v-for="myCourse in myCourses" style="border-bottom: 1px solid #f4f4f4;padding: 10px;">
      <Row>
        <Col span="8">
          <h6>课程名称：{{myCourse.course_name}}</h6>
          <p>
            <img v-if="myCourse.small_image" :src="myCourse.small_image" height="120" width="200"/>
            <img v-else src="../../../assets/default.png" height="120" width="200"/>
          </p>
          <p><ChangeCourseImg :course="myCourse" @uploadComplete="uploadImgComplete"/></p>
        </Col>
        <Col span="16">
          <CourseMeta :course="myCourse"/>
          <p>课程更新状态：{{myCourse.course_status}}
            <span v-if="myCourse.course_status != '已完结'">
                <a href="javascript:;" style="color:#f55e13;font-family: Arial;font-weight: 700;"
                  @click="endUpdate(myCourse.id)">完结更新</a>
            </span>
          </p>
          <p><router-link :to="{path:'/ilearning/course_detail',query:{course_id:myCourse.id}}"
              style="color:green;font-family: Arial;font-weight: 700;">查看视频详情</router-link></p>
          <p><UploadVideo v-if="myCourse.course_status != '已完结'" :course="myCourse" @uploadComplete="uploadVideoComplete"/></p>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import {GetMyCourseList} from "../../../api"
  import {EndUpdate} from "../../../api"
  import ChangeCourseImg from "../Course/ChangeCourseImg"
  import UploadVideo from "../Course/UploadVideo"
  import {getCookie} from "../../../tools"
  import CourseMeta from "../Course/CourseMeta";

  export default {
    name: "MyCourseList",
    components:{CourseMeta, ChangeCourseImg,UploadVideo},
    data(){
      return {
        // 我的课程
        myCourses:[],
      }
    },
    methods:{
      refreshMyCourseList:async function () {
        var userName = getCookie("userName");
        const result = await GetMyCourseList(userName);
        if(result.status=="SUCCESS"){
          this.myCourses = result.courses;
        }
      },
      uploadImgComplete:function () {
        this.refreshMyCourseList();
      },
      uploadVideoComplete:function () {
        this.refreshMyCourseList();
      },
      endUpdate:async function (course_id) {
        const result = await EndUpdate(course_id);
        if(result.status=="SUCCESS"){
          this.refreshMyCourseList();
        }
      }
    },
    mounted:function () {
      this.refreshMyCourseList();
    }
  }
</script>

<style scoped>

</style>
