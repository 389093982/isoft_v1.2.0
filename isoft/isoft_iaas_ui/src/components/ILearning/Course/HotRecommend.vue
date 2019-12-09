<template>
  <div>
    <!-- 列表形式显示 -->
    <div v-if="showMode=='list'">
      <IBeautifulCard title="热门课程推荐">
        <div slot="content" style="padding: 20px;">
          <Row v-for="course in courses">
            <IBeautifulLink @click="$router.push({path:'/ilearning/course_detail',query:{course_id:course.id}})">
              {{course.course_name}}
            </IBeautifulLink>
          </Row>
        </div>
      </IBeautifulCard>
    </div>
    <!-- 图标形式显示 -->
    <IBeautifulCard v-else title="热门课程推荐">
      <div slot="content" style="min-height:850px;padding: 10px;">
        <ul class="clear">
          <li v-for="course in courses">
            <router-link :to="{path:'/ilearning/course_detail',query:{course_id:course.id}}">
              <img v-if="course.small_image" :src="course.small_image" height="100" width="155"/>
              <img v-else src="../../../assets/default.png" height="100" width="155"/>
              <!-- 播放图标 -->
              <div v-if="showDisplayIcon" class="ico_play"></div>
              <p class="isoft_font12">{{course.course_name}}</p>
            </router-link>
          </li>
        </ul>

        <div style="text-align: right;margin: 20px 20px 0 0;">
          提供2019年至2020年最新热门编程视频，本视频每月定期更新，所有版本版权归xxx所有
        </div>
      </div>
    </IBeautifulCard>
  </div>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import {GetHotCourseRecommend} from "../../../api"
  import IBeautifulLink from "../../Common/link/IBeautifulLink";


  export default {
    name: "HotRecommend",
    components:{IBeautifulLink, IBeautifulCard},
    props:{
      // 显示方式,支持 detail 和 list
      showMode: {
        type: [String],
        default: 'detail'
      },
      showDisplayIcon:{
        type:Boolean,
        default:false,
      }
    },
    data(){
      return {
        courses:[],
      }
    },
    methods:{
      refreshHotRecommend:async function(){
        const result = await GetHotCourseRecommend();
        if(result.status == "SUCCESS"){
          this.courses = result.courses;
        }
      }
    },
    mounted:function () {
      this.refreshHotRecommend();
    }
  }
</script>

<style scoped>
  /* 引入公共样式库 */
  @import "../../../assets/css/isoft_common.css";

  a{
    color: black;
  }
  li{
    float: left;
    padding: 10px 9px 0;
    width: 175px;
    border: 1px solid #FFFFFF;
    overflow: hidden;
    text-align: center;
    position: relative;
  }
  li:hover{
    background-color: #f4f4f4;
    border: 1px solid #d0cdd2;
  }
  li:hover a{
    color:red;
  }
  .ico_play {
    background: url(../../../assets/ico_play.png) no-repeat;
    position: absolute;
    top: 35px;
    left: 55px;
    width: 60px;
    height: 60px;
  }

</style>
