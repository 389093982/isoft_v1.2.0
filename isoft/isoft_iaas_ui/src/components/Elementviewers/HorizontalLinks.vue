<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement">
    <Row style="padding: 10px;">
      <Col span="2">{{placement_label}}</Col>
      <Col span="2" v-for="element in elements"><IBeautifulLink2>{{element.title}}</IBeautifulLink2></Col>
    </Row>
  </ElementsLoader>
</template>

<script>
  import ElementsLoader from "../Background/CMS/ElementsLoader"
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2"

  export default {
    name: "HorizontalLinks",
    components:{IBeautifulLink2,ElementsLoader},
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

</style>
