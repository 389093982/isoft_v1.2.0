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
        <CheckboxGroup>
          <Checkbox :label="moduleWork.work_name" v-for="moduleWork in works" v-if="moduleWork.module_name == module.module_name"></Checkbox>
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
    },
    mounted(){
      this.refreshFilterList();
    },
  }
</script>

<style scoped>

</style>
