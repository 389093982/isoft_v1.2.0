<template>
  <IBeautifulCard title="工具盒">
    <div slot="content">
      <ul>
        <li v-for="element in elements" class="li">
          <a :href="element.img_path" target="_blank">
            <img :src="element.img_path" width="25px" height="25px"/>
            <p class="label">{{element.title}}</p>
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

  export default {
    name: "ToolBox",
    components:{IBeautifulCard},
    data(){
      return {
        elements:[],
        tools:[
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"博客中心",
            "href":"/iblog/blog_list"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"心声社区",
            "href":"https://www.baidu.com"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"招聘空间",
            "href":"https://www.baidu.com"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"后台管理",
            "href":"https://www.baidu.com"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"商标查询",
            "href":"https://www.baidu.com"
          },
          {
            "icon":"https://as.zbjimg.com/rake-component/fe-employer/fe-zbj-quickservice-res/0.0.8/images/qs-icons.png",
            "label":"商标查询",
            "href":"https://www.baidu.com"
          }
        ],
      }
    },
    methods:{
      refreshElement: async function () {
        const result = await FilterElementByPlacement(this.GLOBAL.element_host_toolbox_list);
        if(result.status == "SUCCESS"){
          this.elements = result.elements;
        }
      }
    },
    mounted(){
      this.refreshElement();
    }
  }
</script>

<style scoped>
 .li {
   list-style: none;
   float: left;
   width: 25%;
   height: 75px;
   padding: 10px 0;
   text-align: center;
   border: 1px solid transparent;
   cursor: pointer;
 }
 .icon {
   background-position: 0 0;
   display: inline-block;
   width: 30px;height: 30px;
   vertical-align: top;
 }
 .label {
   margin: 10px 0 0;
   color: #333;
   font-size: 12px;
   line-height: 1;
 }
</style>
