<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement">
    <div style="padding: 10px;">
      <Row>
        <Col span="2">
          <span style="font-weight: bold;">课程大类：</span>
        </Col>
        <Col span="22">
          <span v-for="element in elements">
            <a style="margin-right: 10px;" v-if="element.navigation_level == 0"
               @click="currentElement=element">{{element.title}}</a>
          </span>
        </Col>
      </Row>
      <Row>
        <Col span="2">
          <span style="font-weight: bold;">详细分类：</span>
        </Col>
        <Col span="22">
          <span v-for="element in elements" style="margin-right:10px;"
                v-if="currentElement != null && element.navigation_parent_id == currentElement.id">
            <a @click="chooseCourseType(currentElement.title, element.title)">
              {{element.title}}
            </a>
          </span>
        </Col>
      </Row>
    </div>
  </ElementsLoader>
</template>

<script>
  import ElementsLoader from "../../Background/CMS/ElementsLoader"

  export default {
    name: "HotCourseType",
    components:{ElementsLoader},
    props:{
      placement_name:{
        type:String,
        default: '',
      }
    },
    data(){
      return {
        currentElement:null,
        placement_label:'',
        elements:[],
      }
    },
    methods: {
      chooseCourseType:function (course_type, course_sub_type) {
        this.$emit("chooseCourseType", course_type, course_sub_type);
      },
      onLoadElement:function (placement_label, elements) {
        this.placement_label = placement_label;
        this.elements = elements;
        this.currentElement = elements.filter(element => element.navigation_level == 0)[0];
      }
    },
  }
</script>

<style scoped>
  a{
    color: #626262;
  }
  a:hover{
    color: red;
    border-bottom: 2px solid red;
  }
</style>
