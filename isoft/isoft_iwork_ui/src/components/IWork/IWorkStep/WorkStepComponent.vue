<template>
  <Drawer title="全部组件" placement="left" :closable="true" :mask="false" v-model="showComponentDrawer">
    <span v-for="default_work_step_type in default_work_step_types"
          draggable="true" @dragstart="dragstart($event, default_work_step_type.name)">
     <Tag v-if="showComponent(default_work_step_type.name)">{{default_work_step_type.name}}</Tag>
    </span>

    <span>
      AAAAAAAAAAAAAAAAAAAAAA
    </span>
  </Drawer>
</template>

<script>
  import {GetAllWorks} from "../../../api"
  import {oneOf} from "../../../tools"

  export default {
    name: "WorkStepComponent",
    data(){
      return {
        showComponentDrawer:false,
        default_work_step_types: this.GLOBAL.default_work_step_types,
      }
    },
    methods:{
      showComponent:function(name){
        return !oneOf(name, ['work_start',"work_end"]);   // 开始和结束节点不能添加和拖拽
      },
      dragstart:function(event, transferData){
        event.dataTransfer.setData("Text", transferData);
      },
      toggleShow:function () {
        this.showComponentDrawer = !this.showComponentDrawer;
      },
      refreshAllWorks:function () {

      }
    },
    mounted(){
      this.refreshAllWorks();
    }
  }
</script>

<style scoped>

</style>
