<template>
  <Drawer title="全部组件" placement="left" :closable="true" :mask="false" v-model="showComponentDrawer">
    <Tabs value="name1">
      <TabPane label="组件" name="name1">
        <span v-for="default_work_step_type in default_work_step_types"
                draggable="true" @dragstart="dragstart($event, default_work_step_type.name)">
         <Tag v-if="showComponent(default_work_step_type.name)">{{default_work_step_type.name}}</Tag>
        </span>
      </TabPane>

      <TabPane label="流程" name="name2">
        <div v-for="work in works">
          <Tag>{{ work.module_name }} ~~~ {{ work.work_name }}</Tag>
        </div>
      </TabPane>
    </Tabs>
  </Drawer>
</template>

<script>
  import {GetAllFiltersAndWorks} from "../../../api"
  import {oneOf} from "../../../tools"

  export default {
    name: "WorkStepComponent",
    data(){
      return {
        showComponentDrawer:false,
        default_work_step_types: this.GLOBAL.default_work_step_types,
        works:[],
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
      refreshAllWorks:async function () {
        const result = await GetAllFiltersAndWorks();
        if(result.status == "SUCCESS"){
          this.works = result.works;
        }
      }
    },
    mounted(){
      this.refreshAllWorks();
    }
  }
</script>

<style scoped>

</style>
