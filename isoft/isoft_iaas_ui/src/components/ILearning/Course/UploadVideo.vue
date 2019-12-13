<template>
  <span>
    <Modal
      v-model="showDialog"
      width="700"
      title="上传/更新视频"
      :mask-closable="false">
      <div>
        <p style="padding:10px;">课程名称：{{course.course_name}}</p>
        <p style="background-color: rgba(253,0,0,0.11);padding: 10px;">
          上传规则：1、上传视频暂不支持删除功能！2、可上传替换每一集视频 3、视频格式仅支持 mp4 格式！
        </p>

        <Scroll height="220" style="margin: 5px 0;">
          <Tag v-for="(cVideo, num) in cVideos" style="margin: 5px;">
            <span @click="uploadVideoNum = num + 1">第{{num + 1}}集: {{cVideo.video_name}}</span>
          </Tag>
          <Tag>
            <span @click="uploadVideoNum = cVideos.length + 1">新一集</span>
          </Tag>
          <Spin fix size="large" v-if="isLoading">
             <div class="isoft_loading"></div>
          </Spin>
        </Scroll>

         <IFileUpload ref="fileUpload" btn-size="small" :auto-hide-modal="true" :multiple="false"
                      :file-suffixs="['mp4']" :extra-data="{'id':course.id, 'video_number':uploadVideoNum}"
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
  import {handleSpecial} from "../../../tools";

  export default {
    name: "UploadVideo",
    components: {IFileUpload},
    // 当前需要上传视频的课程
    props:["course"],
    data(){
      return {
        isLoading:true,
        showDialog:false,
        // 当前更新视频集数
        uploadVideoNum:-1,
        cVideos:[],
      }
    },
    methods:{
      uploadComplete: async function (data) {
        if(data.status == "SUCCESS"){
          let uploadFilePath = data.fileServerPath;     // uploadFilePath 使用 hash 值时含有特殊字符 + 等需要转义
          let courseId = data.extraData.id;
          let video_number = data.extraData.video_number;
          let filename = data.file.name;      // 上传文件名称
          const result = await UploadVideo(courseId, video_number, filename, handleSpecial(uploadFilePath));
          if(result.status == "SUCCESS"){
            this.refreshCourseDetail(courseId);
            this.$emit('uploadComplete');
          }
        }
      },
      refreshCourseDetail:async function(course_id){
        this.isLoading = true;
        try {
          const result = await ShowCourseDetail(course_id);
          if(result.status=="SUCCESS"){
            this.cVideos = result.cVideos;
          }
        } finally {
          this.isLoading = false;
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
  @import "../../../assets/css/isoft_common.css";

</style>
