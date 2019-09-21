<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement">
    <IBeautifulCard :title="placement_label" v-if="elements.length > 0">
      <div slot="content" style="padding: 10px;">
        <IShowMarkdown v-if="elements[0].md_content" :content="elements[0].md_content"/>
      </div>
    </IBeautifulCard>
  </ElementsLoader>
</template>

<script>
  import IBeautifulCard from "../../Common/card/IBeautifulCard"
  import ElementsLoader from "../../Background/CMS/ElementsLoader"
  import IShowMarkdown from "../../Common/markdown/IShowMarkdown"

  export default {
    name: "MarkDownElementRender",
    components:{IBeautifulCard,IShowMarkdown,ElementsLoader},
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
