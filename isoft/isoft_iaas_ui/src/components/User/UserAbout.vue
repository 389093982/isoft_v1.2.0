<template>
  <div v-if="userName" style="border: 1px #dbdbdb solid;margin:2px 0 5px 5px;padding: 15px;">
    <IBeautifulLink><Avatar :src="user_small_icon" icon="ios-person" size="default"/>&nbsp;{{userName}}</IBeautifulLink>&nbsp;&nbsp;
    <IBeautifulLink style="font-size: 12px;float: right;"
      @onclick="$router.push({path:'/user/mine/detail',query:{username:'mine'}})">个人中心</IBeautifulLink>

    <div style="margin-top: 5px;">
      <Tabs :animated="false">
        <TabPane label="作者课程">
          <Row>
            <Col span="8">课程名称</Col>
            <Col span="8">课程类型</Col>
            <Col span="8">课程子类型</Col>
          </Row>
          <Row v-for="course in courses" :gutter="10">
            <Col span="8">
              <IBeautifulLink @onclick="$router.push({path:'/ilearning/course_detail',query:{course_id:course.id}})">
                {{course.course_name | filterLimitFunc}}
              </IBeautifulLink>
            </Col>
            <Col span="8">
              <IBeautifulLink @onclick="$router.push({ path:'/ilearning/course_search', query: { search: course.course_type }})">
                {{course.course_type | filterLimitFunc}}
              </IBeautifulLink>
            </Col>
            <Col span="8">
              <IBeautifulLink @onclick="$router.push({ path: '/ilearning/course_search', query: { search: course.course_sub_type }})">
                {{course.course_sub_type | filterLimitFunc}}
              </IBeautifulLink>
            </Col>
          </Row>
        </TabPane>
        <TabPane label="作者博文">作者博文</TabPane>
        <TabPane label="作者博文">作者博文</TabPane>
      </Tabs>
    </div>
  </div>
</template>

<script>
  import {GetCourseListByUserName,GetUserDetail} from "../../api"
  import IBeautifulLink from "../Common/link/IBeautifulLink"

  export default {
    name: "UserAbout",
    components: {IBeautifulLink},
    props:{
      userName: {
        type: String,
        default: ''
      },
    },
    data(){
      return {
        // 当前 userName 的课程列表
        courses:[],
        // 当前 user 对应头像信息
        user_small_icon:'',
      }
    },
    methods:{
      refreshUserInfo:function () {
        this.refreshCourseList();
        this.refreshUserDetail();
      },
      refreshUserDetail:async function(){
        const result = await GetUserDetail(this.userName);
        if(result.status == "SUCCESS"){
          this.user_small_icon = result.user.small_icon;
        }
      },
      refreshCourseList:async function () {
        const result = await GetCourseListByUserName(this.userName);
        if(result.status=="SUCCESS"){
          this.courses = result.courses;
        }
      },
    },
    watch:{
      "userName": "refreshUserInfo"      // 如果 userName 有变化,会再次执行该方法
    },
    filters:{
      // 内容超长则显示部分
      filterLimitFunc:function (value) {
        if(value && value.length > 12) {
          value= value.substring(0,12) + '...';
        }
        return value;
      },
    },
    mounted(){
      this.refreshUserInfo();
    }
  }
</script>

<style scoped>

</style>
