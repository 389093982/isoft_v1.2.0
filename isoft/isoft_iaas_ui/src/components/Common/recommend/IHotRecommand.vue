<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement">
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
  </ElementsLoader>
</template>

<script>
  import ElementsLoader from "../../Background/CMS/ElementsLoader"

  import IBeautifulButtonLink from "../../Common/link/IBeautifulButtonLink"

  export default {
    name: "IHotRecommand",
    components:{IBeautifulButtonLink,ElementsLoader},
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
      onLoadElement:function (placement_label, elements) {
        this.placement_label = placement_label;
        this.elements = elements;
      }
    },
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
