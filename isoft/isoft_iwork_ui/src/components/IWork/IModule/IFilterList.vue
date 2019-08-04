<template>
  <div>
    <Select v-model="current_filter_id" style="width:400px">
      <Option v-for="filterWork in filterWorks" :value="filterWork.id" :key="filterWork.work_name">
        {{filterWork.work_name}}
      </Option>
    </Select>

    <ul>
      <li v-for="module in modules" style="list-style: none;margin-top: 20px;">
        <Tag><span>{{module.module_name}}</span></Tag>
        <Button type="success" size="small" @click="handleCheckAll(module.module_name)">全选</Button>
        <CheckboxGroup v-model="checkedWorks">
          <Checkbox :label="moduleWork.work_name" v-for="moduleWork in works" v-if="moduleWork.module_name == module.module_name"
            ></Checkbox>
        </CheckboxGroup>
      </li>
    </ul>

  </div>
</template>

<script>
  import {GetAllFilterWorks} from "../../../api"
  import {oneOf} from "../../../tools"

  export default {
    name: "IFilterList",
    data(){
      return {
        filterWorks:[],
        modules:[],
        works: [],
        moduleWorks:[],
        current_filter_id:-1,
        checkedWorks:[],
      }
    },
    methods:{
      refreshFilterList:async function () {
        const result = await GetAllFilterWorks();
        if(result.status == "SUCCESS"){
          this.filterWorks = result.filters;
          this.modules = result.modules;
          this.works = result.works;
        }
      },
      handleCheckAll:function (module_name) {
        let work_names = this.works.filter(work => work.module_name == module_name).map(work => work.work_name);
        var allInflag = true;
        for(var i=0; i<work_names.length; i++){
          let work_name = work_names[i];
          if(!oneOf(work_name, this.checkedWorks)){
            allInflag = false;
            break;
          }
        }
        if(allInflag){
          this.checkedWorks = this.checkedWorks.filter(checkWork => !oneOf(checkWork, work_names));
        }else{
          let addWorks = work_names.filter(work_name => !oneOf(work_name, this.checkedWorks));
          for(var i=0; i<addWorks.length; i++){
            let work_name = addWorks[i];
            this.checkedWorks.push(work_name);
          }
        }
      }
    },
    mounted(){
      this.refreshFilterList();
    },
  }
</script>

<style scoped>

</style>
