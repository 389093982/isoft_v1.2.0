<template>
  <div>
    <isoft-lazy :time="1000">
      <div v-if="myCourses.length > 0">
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
      <div v-else>您还没有任何课程奥，如果你想传播你的知识，
        <IBeautifulLink2 font-weight="bold" @onclick="$router.push({path:'/ilearning/mine/course_space/newCourse'})">请前去开课！</IBeautifulLink2>
      </div>
    </isoft-lazy>
  </div>
</template>

<script>
  import {GetMyCourseList,EndUpdate,UpdateCourseIcon} from "../../../api"
  import UploadVideo from "../Course/UploadVideo"
  import {getCookie} from "../../../tools"
  import CourseMeta from "../Course/CourseMeta";
  import IFileUpload from "../../Common/file/IFileUpload";

  export default {
    name: "MyCourseList",
    components:{IFileUpload,CourseMeta,UploadVideo},
    data(){
      return {
        // 我的课程
        myCourses:[],
      }
    },
    methods:{
      uploadComplete: async function (data) {
        if(data.status == "SUCCESS"){
          let uploadFilePath = data.fileServerPath;
          let courseId = data.extraData;
          const result = await UpdateCourseIcon(courseId, uploadFilePath);
          if(result.status == "SUCCESS"){
            this.refreshMyCourseList();
          }
        }
      },
      refreshMyCourseList:async function () {
        var userName = getCookie("userName");
        const result = await GetMyCourseList(userName);
        if(result.status=="SUCCESS"){
          this.myCourses = result.courses;
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

</style>
