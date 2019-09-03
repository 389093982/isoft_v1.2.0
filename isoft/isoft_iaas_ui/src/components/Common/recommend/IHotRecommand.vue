<template>
  <div>
    <span>
      <h6 class="title" :title="placement_label">{{placement_label}}</h6>
    </span>
    <hr style="border:1px solid #eee;height: 1px;"/>
    <Row :gutter="50">
      <Col span="8" style="margin-top: 12px;" v-for="element in elements">
        <span style="font-size: 14px;">{{element.title}}</span>
        <IBeautifulButtonLink msg="点击了解详情" floatstyle="right" :hrefaddr="element.linked_refer"/>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {FilterElementByPlacement} from "../../../api"
  import {checkEmpty} from "../../../tools"
  import IBeautifulButtonLink from "../../Common/link/IBeautifulButtonLink"

  export default {
    name: "IHotRecommand",
    components:{IBeautifulButtonLink},
    props:{
      placement_name:{
        type:String,
        default: '',
      }
    },
    data(){
      return {
        elements:[],
        placement_label:'',
      }
    },
    methods:{
      refreshElement: async function () {
        if(!checkEmpty(this.placement_name)){
          const result = await FilterElementByPlacement(this.placement_name);
          if(result.status == "SUCCESS"){
            this.placement_label = result.placement.placement_label;
            this.elements = result.elements;
          }
        }
      }
    },
    mounted(){
      this.refreshElement();
    }
  }
</script>

<style scoped>
  .title{
    width: 200px;
    height: 35px;
    font-size: 18px;
    line-height: 35px;
    text-align: center;
    background: #3b80db;
    color: #fff;
    position: relative;
    font-weight: normal;
    margin: 0;
    padding: 0;
    font-family: "微软雅黑";
  }
</style>
