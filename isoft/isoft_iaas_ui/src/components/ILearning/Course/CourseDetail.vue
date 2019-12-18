<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <Row :gutter="10">
      <!-- 左侧课程详情部分 -->
      <Col span="16">
        <!-- 头部 -->
        <Row class="header">
          <Col span="8">
            <h4>课程名称：{{course.course_name}}</h4>
            <div class="course_img">
              <img :src="course.small_image" width="180" height="120" @error="defImg()"/>
              <div class="course_name">{{course.course_name}}</div>
            </div>
          </Col>
          <Col span="16">
            <CourseMeta :course="course"/>
            <p>
              标签： <Tag>生动形象</Tag> <Tag>生动形象</Tag> <Tag>java</Tag> <Tag>java</Tag> <Tag>java</Tag>
            </p>
            <p>
              <a href="javascript:;" v-if="course_collect==true" @click="toggle_favorite(course.id,'course_collect', '取消收藏')">取消收藏</a>
              <a href="javascript:;" v-else @click="toggle_favorite(course.id,'course_collect', '收藏')">加入收藏</a>&nbsp;
              <a href="javascript:;" v-if="course_parise==true" @click="toggle_favorite(course.id,'course_praise', '取消点赞')">取消点赞</a>
              <a href="javascript:;" v-else @click="toggle_favorite(course.id,'course_praise', '点赞')">我要点赞</a>
            </p>
          </Col>
        </Row>
        <hr style="margin-top: 10px;">
        <!-- 视频链接 -->
        <Row style="margin: 10px 0;min-height: 200px;">
          <Col span="12" v-for="cVideo in cVideos" style="padding: 5px;">
            <Row style="background-color: #f8f8f8;padding: 3px;">
              <Col span="20">
                <IBeautifulLink>第{{cVideo.video_number}}集：{{cVideo.video_name}}</IBeautifulLink>
              </Col>
              <Col span="4">
                <router-link :to="{path:'/ilearning/video_play',query:{course_id:course.id,video_id:cVideo.id}}">
                  <Button size="small" type="success" class="hovered hvr-grow">立即播放</Button>
                </router-link>
              </Col>
            </Row>
          </Col>
          <Spin fix size="large" v-if="isLoading">
            <div class="isoft_loading"></div>
          </Spin>
        </Row>
        <hr>
        <!-- 评论模块 -->
        <IEasyComment :theme_pk="course.id" theme_type="course_theme_type" style="margin-top: 50px;"/>
      </Col>

      <Col span="8">
        <UserAbout :userName="course.course_author"/>
        <HotUser style="margin-left: 2px;"/>
        <HotRecommend showMode="list" style="margin-left: 2px;"/>
        <!-- 推荐系统 -->
        <Recommand />
        <CommunicationGroup :placement_name="GLOBAL.placement_communication_group"/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {ShowCourseDetail,ToggleFavorite} from "../../../api"
  import Recommand from "./Recommand"
  import IEasyComment from "../../Comment/IEasyComment"
  import CommunicationGroup from "../Site/CommunicationGroup"
  import HotRecommend from "./HotRecommend"
  import UserAbout from "../../User/UserAbout"
  import HotUser from "../../User/HotUser"
  import CourseMeta from "./CourseMeta";

  export default {
    name: "CourseDetail",
    components:{CourseMeta, IEasyComment,Recommand,CommunicationGroup,HotRecommend,UserAbout,HotUser},
    data(){
      return {
        isLoading:true,
        defaultImg: require('../../../assets/default.png'),
        // 当前课程
        course:{},
        // 视频清单
        cVideos:[],
        // 课程收藏
        course_collect:false,
        // 课程点赞
        course_parise:false,
      }
    },
    methods:{
      defImg(){
        let img = event.srcElement;
        img.src = this.defaultImg;
        img.onerror = null; //防止闪图
      },
      refreshCourseDetail:async function(){
        this.isLoading = true;
        try{
          const course_id = this.$route.query.course_id;
          const result = await ShowCourseDetail(course_id);
          if(result.status=="SUCCESS"){
            this.course = result.course;
            this.cVideos = result.cVideos;
            this.course_collect = result.course_collect;
            this.course_parise = result.course_parise;
          }
        } finally {
          this.isLoading = false;
        }
      },
      toggle_favorite:async function (favorite_id, favorite_type, message) {
        const result = await ToggleFavorite(favorite_id, favorite_type);
        if(result.status=="SUCCESS"){
          this.$Message.success(message + "成功!");
          this.refreshCourseDetail();
        }
      }
    },
    mounted:function () {
      this.refreshCourseDetail(this.$route.query.course_id);
    },
    watch:{
      "$route.params": "refreshCourseDetail"      // 如果 $route.params 有变化,会再次执行该方法
    }
  }
</script>

<style scoped>
  @import "../../../assets/css/isoft_common.css";

  .header a{
    color: red;
  }
  .course_img{
    width: 180px;
    height: 120px;
    cursor: pointer;
  }
  .course_img .course_name {
    display: none;
    padding: 3px 0 0 10px;
    background-color: rgba(0,0,0,0.6);
    color: white;
    height: 30px;
    position: relative;
    top: 0px;
    transition: all ease-in 1s;
  }
  .course_img:hover .course_name {
    display: block;
    top: -30px;
  }
</style>
