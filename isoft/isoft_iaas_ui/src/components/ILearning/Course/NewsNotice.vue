<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement">
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
  </ElementsLoader>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import ElementsLoader from "../../Background/CMS/ElementsLoader"
  import IBeautifulLink2 from "../../Common/link/IBeautifulLink2";

  export default {
    name: "NewsNotice",
    components:{IBeautifulLink2, IBeautifulCard,ElementsLoader},
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
