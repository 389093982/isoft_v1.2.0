<template>
  <IBeautifulCard :title="placement_label" v-if="elements.length > 0">
    <div slot="content">
      <IShowMarkdown v-if="elements[0].md_content" :content="elements[0].md_content"/>
    </div>
  </IBeautifulCard>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import {FilterElementByPlacement} from "../../../api"
  import {checkEmpty} from "../../../tools"
  import IShowMarkdown from "../../Common/markdown/IShowMarkdown"

  export default {
    name: "MdCourseDesc",
    components:{IBeautifulCard,IShowMarkdown},
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
