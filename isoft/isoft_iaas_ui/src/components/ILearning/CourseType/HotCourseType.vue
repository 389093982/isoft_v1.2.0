<template>
  <div>
    <div style="padding: 10px;">
      课程大类：
      <span v-for="element in elements">
        <IBeautifulLink style="margin-right: 10px;" v-if="element.navigation_level == 0"
                        @onclick="currentElement=element">{{element.title}}</IBeautifulLink>
      </span>
    </div>
    <div style="padding: 10px;">
      详细分类：
      <span v-for="element in elements" style="margin-right: 10px;"
            v-if="currentElement != null && element.navigation_parent_id == currentElement.id">
        <IBeautifulLink @onclick="chooseCourseType(currentElement.title, element.title)">
          {{element.title}}
        </IBeautifulLink>
      </span>
    </div>
  </div>
</template>

<script>
  import IBeautifulLink from "../../Common/link/IBeautifulLink"
  import {FilterElementByPlacement} from "../../../api"

  export default {
    name: "HotCourseType",
    components:{IBeautifulLink},
    data(){
      return {
        currentElement:null,
        elements:[],
      }
    },
    methods: {
      chooseCourseType:function (course_type, course_sub_type) {
        this.$emit("chooseCourseType", course_type, course_sub_type);
      },
      refreshElement: async function () {
        const result = await FilterElementByPlacement(this.GLOBAL.element_host_course_type_carousel);
        if(result.status == "SUCCESS"){
          this.elements = result.elements;
          this.currentElement = result.elements.filter(element => element.navigation_level == 0)[0];
        }
      }
    },
    mounted(){
      this.refreshElement();
    }
  }
</script>

<style scoped>
  a{
    color: #626262;
  }
  a:hover{
    color: red;
  }
</style>
