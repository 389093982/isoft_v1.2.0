<template>
  <IBeautifulCard :title="placement_label">
    <ul slot="content" style="padding-left: 0em;">
      <li v-for="element in elements">
        <span style="float:right;color: #999;"><Time :time="element.created_time" type="date"/></span>
        <a target="_blank">
          {{element.title}} <img src="../../../assets/news.gif">
        </a>
      </li>
    </ul>
    <span slot="header_right">
      <IBeautifulLink2 style="font-size: 12px;">更多</IBeautifulLink2>
    </span>
  </IBeautifulCard>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import {FilterElementByPlacement} from "../../../api"
  import {checkEmpty} from "../../../tools"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";

  export default {
    name: "NewsNotice",
    components:{IBeautifulLink2, IBeautifulCard},
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
  a:hover{
    color: red;
  }
  li {
    padding: 3px 15px 0 15px;
    list-style: none;
    font-size: 14px;
  }
</style>
