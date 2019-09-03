<template>
  <div style="margin: 10px;">
    <Row>
      <Col span="2">{{placement_label}}</Col>
      <Col span="2" v-for="element in elements"><IBeautifulLink2>{{element.title}}</IBeautifulLink2></Col>
    </Row>
  </div>
</template>

<script>
  import {FilterElementByPlacement} from "../../api"
  import {checkEmpty} from "../../tools"
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2"

  export default {
    name: "HorizontalLinks",
    components:{IBeautifulLink2},
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

</style>
