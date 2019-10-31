<template>
  <div style="background: #FFFFFF;padding: 10px;">
    <Row :gutter="10">
      <!-- 左侧课程详情部分 -->
      <Col span="16">
        <!-- 头部 -->
        <Row>
          <Col span="8">
            <h4>课程名称：{{course.course_name}}</h4>
            <p>
              <img v-if="course.small_image" :src="course.small_image" height="120" width="180"/>
              <img v-else src="../../../assets/default.png" height="120" width="180"/>
            </p>
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
        <Row style="margin: 10px 0 10px 0">
          <Col span="12" v-for="cVideo in cVideos" style="padding: 5px;">
            <Row>
              <Col span="2">{{cVideo.video_number}}</Col>
              <Col span="18">{{cVideo.video_name}}</Col>
              <Col span="4">
                <router-link :to="{path:'/ilearning/video_play',query:{course_id:course.id,video_id:cVideo.id}}">
                  <Button size="small" type="success" class="hovered hvr-grow">立即播放</Button>
                </router-link>
              </Col>
            </Row>
          </Col>
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
        <CommunicationGroup :placement_name="GLOBAL.communication_group"/>
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
      refreshCourseDetail:async function(){
        const course_id = this.$route.query.course_id;
        const result = await ShowCourseDetail(course_id);
        if(result.status=="SUCCESS"){
          this.course = result.course;
          this.cVideos = result.cVideos;
          this.course_collect = result.course_collect;
          this.course_parise = result.course_parise;
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
  a{
    color: red;
  }
</style>
