<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement"
                   style="border: 1px #dbdbdb solid;margin-left: 5px;margin-top: 5px;padding: 15px;">
    <span style="font-size: 16px;">{{placement_label}}</span>
    <ul style="list-style: none;">
      <li v-for="element in elements">{{element.content}}</li>
    </ul>
  </ElementsLoader>
</template>

<script>
  import ElementsLoader from "../../Background/CMS/ElementsLoader"

  export default {
    name: "CommunicationGroup",
    components:{ElementsLoader},
    data(){
      return {
        elements:[],
        placement_label:'',
      }
    },
    props:{
      placement_name:{
        type:String,
        default: '',
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
