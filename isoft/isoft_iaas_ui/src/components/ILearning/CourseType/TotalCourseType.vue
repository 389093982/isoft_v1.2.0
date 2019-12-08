<template>
  <div>
    <div style="padding-top: 10px;">
      <Row>
        <Col span="2">
          <span style="font-weight: bold;">课程大类：</span>
        </Col>
        <Col span="22">
          <span v-for="course_type in course_types">
            <a class="isoft_font12" style="margin-right: 10px;" @click="loadSubCourseType(course_type)">{{course_type.course_type}}</a>
          </span>
        </Col>
      </Row>
      <Row>
        <Col span="2">
          <span style="font-weight: bold;">详细分类：</span>
        </Col>
        <Col span="22">
          <span v-for="sub_course_type in sub_course_types" style="margin-right:10px;">
             <a class="isoft_font12" @click="chooseCourseType(current_course_type, sub_course_type.course_sub_type)">{{sub_course_type.course_sub_type}}</a>
          </span>
        </Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import {GetAllCourseType} from "../../../api"
  import {GetAllCourseSubType} from "../../../api"

  export default {
    name: "TotalCourseType",
    components:{},
    data(){
      return {
        course_types:[],
        current_course_type:'',
        sub_course_types:[],
      }
    },
    methods:{
      refreshCourseType:async function () {
        const result = await GetAllCourseType();
        if(result.status=="SUCCESS"){
          this.course_types = result.course_types;
          this.loadSubCourseType(result.course_types[0]);
        }
      },
      loadSubCourseType:async function(course_type){
        const result = await GetAllCourseSubType(course_type.course_type);
        if(result.status=="SUCCESS"){
          this.current_course_type = course_type;
          this.sub_course_types = result.sub_course_types;
        }
      },
      chooseCourseType:function (course_type, course_sub_type) {
        this.$emit("chooseCourseType", course_type, course_sub_type);
      }
    },
    mounted:function () {
      this.refreshCourseType();
    }
  }
</script>

<style scoped>
  @import "../../../assets/css/isoft_common.css";

  a{
    color: #657180;
  }
  a:hover{
    color: red;
    border-bottom: 2px solid red;
  }
</style>
