<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement">
    <Row style="padding: 10px;">
      <Col span="2">{{placement_label}}</Col>
      <Col span="2" v-for="element in elements">
        <IBeautifulLink @onclick="openLink(element)">{{element.element_label}}</IBeautifulLink>
      </Col>
    </Row>
  </ElementsLoader>
</template>

<script>
  import ElementsLoader from "../Background/CMS/ElementsLoader"
  import IBeautifulLink from "../Common/link/IBeautifulLink"

  export default {
    name: "HorizontalLinks",
    components:{IBeautifulLink,ElementsLoader},
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
      openLink:function (element) {
        window.location.href=element.linked_refer;
      },
      onLoadElement:function (placement_label, elements) {
        this.placement_label = placement_label;
        this.elements = elements;
      }
    },
  }
</script>

<style scoped>

</style>
