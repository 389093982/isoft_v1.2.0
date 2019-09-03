<template>
  <div style="border: 1px #dbdbdb solid;margin-left: 5px;margin-top: 5px;padding: 15px;">
    <span style="font-size: 16px;">{{placement_label}}</span>
    <ul style="list-style: none;">
      <li v-for="element in elements">{{element.content}}</li>
    </ul>
  </div>
</template>

<script>
  import {FilterElementByPlacement} from "../../../api"
  import {checkEmpty} from "../../../tools"

  export default {
    name: "CommunicationGroup",
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
