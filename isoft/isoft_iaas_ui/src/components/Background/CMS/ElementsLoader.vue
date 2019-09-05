<template>
  <div>
    <slot></slot>
  </div>
</template>

<script>
  import {FilterElementByPlacement} from "../../../api"
  import {checkEmpty} from "../../../tools"

  export default {
    name: "ElementsLoader",
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
              this.$emit("onLoadElement", result.placement.placement_label, result.elements);
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
