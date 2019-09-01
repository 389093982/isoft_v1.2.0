<template>
  <IBeautifulCard :title="placement_label" v-if="elements.length > 0">
    <div slot="content">
      <ul>
        <li v-for="element in elements" class="li">
          <a :href="element.linked_refer" target="_blank">
            <img :src="element.img_path" width="30px" height="30px"/>
            <p style="font-size: 12px;">{{element.title}}</p>
          </a>
        </li>
      </ul>
      <div style="clear: both;"></div>
    </div>
  </IBeautifulCard>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import {FilterElementByPlacement} from "../../../api"
  import {checkEmpty} from "../../../tools"

  export default {
    name: "ToolBox",
    components:{IBeautifulCard},
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
  a{
    color: black;
  }
 .li {
   list-style: none;
   float: left;
   width: 70px;
   height: 70px;
   padding: 10px 10px;
   text-align: center;
   border: 1px solid transparent;
   cursor: pointer;
 }
 li:hover{
   background-color: #f4f4f4;
   border: 1px solid #d0cdd2;
 }
 li:hover a{
   color:red;
 }
</style>
