<template>
  <span>
    <Modal
      v-model="showDialog"
      width="500"
      title="上传/更新视频"
      :mask-closable="false">
      <div>
        <p>课程名称：{{course.course_name}}</p>

        <div style="margin-top: 20px;margin-bottom: 20px;">
          <Button size="small" v-for="(cVideo, num) in cVideos"
                  type="success" @click="uploadVideoNum = num + 1" style="margin: 5px;">第{{num + 1}}集: {{cVideo.video_name}}</Button>
          <Button size="small" type="success" @click="uploadVideoNum = cVideos.length + 1">新一集{{cVideos.length + 1}}</Button>
        </div>

         <IFileUpload ref="fileUpload" btn-size="small" :auto-hide-modal="true" :extra-data="{'id':course.id, 'video_number':uploadVideoNum}"
                      @uploadComplete="uploadComplete" action="/api/iwork/httpservice/fileUpload" uploadLabel="上传视频"/>
        <span v-if="uploadVideoNum > 0" style="color: green;">*当前更新第{{uploadVideoNum}}集</span>
        <span v-else style="color: red;">请选择更新集数</span>
      </div>
    </Modal>

    <a href="javascript:;" style="color:#f55e13;font-family: Arial;font-weight: 700;" @click="uploadVideoFunc">上传/更新视频</a>
  </span>
</template>

<script>
  import IFileUpload from "../../Common/file/IFileUpload";
  import {UploadVideo,ShowCourseDetail} from "../../../api"

  export default {
    name: "UploadVideo",
    components: {IFileUpload},
    // 当前需要上传视频的课程
    props:["course"],
    data(){
      return {
        showDialog:false,
        // 当前更新视频集数
        uploadVideoNum:-1,
        cVideos:[],
      }
    },
    methods:{
      uploadComplete: async function (data) {
        if(data.status == "SUCCESS"){
          let uploadFilePath = data.fileServerPath;
          let courseId = data.extraData.id;
          let video_number = data.extraData.video_number;
          const result = await UploadVideo(courseId, video_number, '1111111111111',uploadFilePath);
          if(result.status == "SUCCESS"){
            this.showDialog = false;
            this.$emit('uploadComplete');
          }
        }
      },
      refreshCourseDetail:async function(course_id){
        const result = await ShowCourseDetail(course_id);
        if(result.status=="SUCCESS"){
          this.cVideos = result.cVideos;
        }
      },
      uploadVideoFunc:function () {
        this.showDialog = true;
        this.refreshCourseDetail(this.course.id);
      }
    },
  }
</script>

<style scoped>

</style>
