<template>
  <div>
    <slot></slot>
  </div>
</template>

<script>
  import {FilterElementByPlacement} from "../../../api"
  import {expires} from "../../../api"
  import {checkEmpty} from "../../../tools"
  import {_store} from "../../../tools"

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
          var result = null;
          if(_store.getItem(this.placement_name)){
            result = JSON.parse(_store.getItem(this.placement_name));
          }else{
            if(!checkEmpty(this.placement_name)){
              result = await FilterElementByPlacement(this.placement_name);
              if(result.status == "SUCCESS"){
                _store.setItem({
                  name: this.placement_name,
                  value: JSON.stringify(result),
                  expires: expires,
                });
              }
            }
          }
          if(result != null && result.status == "SUCCESS"){
            this.$emit("onLoadElement", result.placement.placement_label, result.elements);
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
