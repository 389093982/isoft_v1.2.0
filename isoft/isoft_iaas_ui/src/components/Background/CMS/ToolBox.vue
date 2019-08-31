<template>
  <IBeautifulCard :title="title">
    <div slot="content">
      <ul>
        <li v-for="element in elements" class="li">
          <a :href="element.linked_refer" target="_blank">
            <img :src="element.img_path" width="25px" height="25px"/>
            <p>{{element.title}}</p>
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
      title:{
        type:String,
        default:'工具盒',
      },
      placement_name:{
        type:String,
        default: '',
      }
    },
    data(){
      return {
        elements:[],
      }
    },
    methods:{
      refreshElement: async function () {
        if(!checkEmpty(this.placement_name)){
          const result = await FilterElementByPlacement(this.placement_name);
          if(result.status == "SUCCESS"){
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
   width: 75px;
   height: 75px;
   padding: 10px 0;
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
