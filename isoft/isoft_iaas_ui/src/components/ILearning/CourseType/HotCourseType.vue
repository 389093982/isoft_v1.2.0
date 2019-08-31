<template>
  <div>
    <div>
      <span style="height: 32px;line-height: 32px;margin-bottom: 5px;color: #000;float: left !important;">课程大类：</span>
      <!-- 对父级CSS选择器加overflow:hidden样式,可以清除父级内使用float产生浮动.优点是可以很少CSS代码即可解决浮动产生 -->
      <ul style="overflow:hidden;">
        <li v-for="element in elements"
            style="height: 32px;line-height: 32px;margin: 0 4px 5px;text-align: center;color: #333;float: left;display: inline;">
          <IBeautifulLink v-if="element.navigation_level == 0" @onclick="currentElement=element">{{element.title}}</IBeautifulLink>
        </li>
      </ul>
    </div>
    <div>
      <span style="height: 32px;line-height: 32px;margin-bottom: 5px;color: #000;float: left !important;">详细分类：</span>
      <ul>
        <li v-for="element in elements" v-if="currentElement != null && element.navigation_parent_id == currentElement.id"
            style="height: 32px;line-height: 32px;margin: 0 4px 5px;text-align: center;color: #333;float: left;display: inline;">
          <IBeautifulLink @onclick="chooseCourseType(currentElement.title, element.title)">
            {{element.title}}
          </IBeautifulLink>
        </li>
      </ul>
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
