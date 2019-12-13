<template>
  <div>
    <div v-if="myCourses && myCourses.length > 0">
      <Row v-for="myCourse in myCourses" style="border-bottom: 1px solid #f4f4f4;padding: 10px;">
        <Col span="8">
          <h4>课程名称：{{myCourse.course_name}}</h4>
          <p>
            <img v-if="myCourse.small_image" :src="myCourse.small_image" height="120" width="180"/>
            <img v-else src="../../../assets/default.png" height="120" width="180"/>
          </p>
          <p>
            <IFileUpload ref="fileUpload" :extra-data="myCourse.id" btn-size="small" :auto-hide-modal="true"
                         @uploadComplete="uploadComplete" action="/api/iwork/httpservice/fileUpload" uploadLabel="换张图片"/>
          </p>
        </Col>
        <Col span="16">
          <CourseMeta :course="myCourse"/>
          <p><router-link :to="{path:'/ilearning/course_detail',query:{course_id:myCourse.id}}"
                          style="color:green;font-family: Arial;font-weight: 700;">查看视频详情</router-link></p>
          <p><UploadVideo :course="myCourse" @uploadComplete="uploadVideoComplete"/></p>
        </Col>
      </Row>
    </div>
    <div v-if="myCourses && myCourses.length == 0" style="padding: 30px 10px;">
      您还没有任何课程奥，如果你想传播你的知识，
      <IBeautifulLink font-weight="bold" @onclick="$router.push({path:'/ilearning/mine/course_space/newCourse'})">请前去开课！</IBeautifulLink>
      <IBeautifulLink @onclick="">如何开课呢？</IBeautifulLink>
    </div>
    <Spin fix size="large" v-if="isLoading">
      <div class="isoft_loading"></div>
    </Spin>
  </div>
</template>

<script>
  import {GetCourseListByUserName,EndUpdate,UpdateCourseIcon} from "../../../api"
  import UploadVideo from "../Course/UploadVideo"
  import {getCookie, handleSpecial} from "../../../tools"
  import CourseMeta from "../Course/CourseMeta";
  import IFileUpload from "../../Common/file/IFileUpload";

  export default {
    name: "MyCourseList",
    components:{IFileUpload,CourseMeta,UploadVideo},
    data(){
      return {
        isLoading:true,
        // 我的课程
        myCourses:null,
      }
    },
    methods:{
      uploadComplete: async function (data) {
        if(data.status == "SUCCESS"){
          let uploadFilePath = data.fileServerPath;
          let courseId = data.extraData;
          const result = await UpdateCourseIcon(courseId, handleSpecial(uploadFilePath));
          if(result.status == "SUCCESS"){
            this.refreshMyCourseList();
          }
        }
      },
      refreshMyCourseList:async function () {
        this.isLoading = true;
        try{
          var userName = getCookie("userName");
          const result = await GetCourseListByUserName(userName);
          if(result.status=="SUCCESS"){
            this.myCourses = result.courses;
          }
        } finally {
          this.isLoading = false;
        }
      },
      uploadVideoComplete:function () {
        this.refreshMyCourseList();
      },
    },
    mounted:function () {
      this.refreshMyCourseList();
    }
  }
</script>

<style scoped>
  @import "../../../assets/css/isoft_common.css";

</style>
