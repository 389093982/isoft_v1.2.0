<template>
  <div v-if="userName" style="border: 1px #dbdbdb solid;margin-left: 5px;margin-bottom: 5px;padding: 15px;">
    <Tag color="primary">{{userName}}</Tag>
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
              <IBeautifulLink2 @onclick="$router.push({path:'/ilearning/course_detail',query:{course_id:course.id}})">
                {{course.course_name | filterLimitFunc}}
              </IBeautifulLink2>
            </Col>
            <Col span="8">
              <IBeautifulLink2 @onclick="$router.push({ path:'/ilearning/course_search', query: { search: course.course_type }})">
                {{course.course_type | filterLimitFunc}}
              </IBeautifulLink2>
            </Col>
            <Col span="8">
              <IBeautifulLink2 @onclick="$router.push({ path: '/ilearning/course_search', query: { search: course.course_sub_type }})">
                {{course.course_sub_type | filterLimitFunc}}
              </IBeautifulLink2>
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
  import {GetMyCourseList} from "../../api"
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2"

  export default {
    name: "UserDetail",
    components: {IBeautifulLink2},
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
      }
    },
    methods:{
      refreshUserInfo:function () {
        this.refreshCourseList();
      },
      refreshCourseList:async function () {
        const result = await GetMyCourseList(this.userName);
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
    }
  }
</script>

<style scoped>

</style>
