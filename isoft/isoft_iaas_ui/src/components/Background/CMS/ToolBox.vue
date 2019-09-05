<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement">
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
  </ElementsLoader>
</template>

<script>
  import ElementsLoader from "./ElementsLoader"

  import IBeautifulCard from "../../Common/card/IBeautifulCard"

  export default {
    name: "ToolBox",
    components:{IBeautifulCard, ElementsLoader},
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
