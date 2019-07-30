<template>
  <div v-if="workstep" style="width: 500px;padding: 10px;">
    <Row>
      <Col span="8">输入参数</Col>
      <Col span="8">输出参数</Col>
      <Col span="8">引用关系</Col>
      <Col span="8" v-if="ParamInputSchemaItems">
        <ul>
          <li v-for="item in ParamInputSchemaItems">
            <Tag>{{item.ParamName}}</Tag>
          </li>
        </ul>
      </Col>
      <Col span="8" v-if="ParamInputSchemaItems">
        <ul>
          <li v-for="item in ParamOutputSchemaItems">
            <Tag>{{item.ParamName}}</Tag>
          </li>
        </ul>
      </Col>
      <Col span="8" v-if="usedMap">
        <span v-for="(usedWorkStepIds, currentWorkStepId) in usedMap">
          <span v-if="currentWorkStepId == workstep.work_step_id">
            <li style="list-style: none;" v-if="usedWorkStepIds" v-for="usedWorkStepId in usedWorkStepIds">
              <Tag>{{usedWorkStepId | renderWorkStepName }}</Tag>
            </li>
          </span>
        </span>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {checkEmpty} from "../../../tools"

  var global_this;
  export default {
    name: "WorkStepPoptip",
    props:{
      worksteps:{
        type: Array,
        default: [],
      },
      workstep:{
        type: Object,
        default: null,
      },
      usedMap:{
        type:Object,
        default: null,
      }
    },
    methods:{
      highlightUsedWorkStepIdFunc:function () {
        alert(JSON.stringify(this.usedMap));
      }
    },
    computed:{
      ParamInputSchemaItems () {
        return checkEmpty(this.workstep.work_step_input) ? null : JSON.parse(this.workstep.work_step_input).ParamInputSchemaItems;
      },
      ParamOutputSchemaItems () {
        return checkEmpty(this.workstep.work_step_output) ? null : JSON.parse(this.workstep.work_step_output).ParamOutputSchemaItems;
      }
    },
    beforeCreate: function () {
      global_this = this;         // 在 beforeCreate中将vue实例赋值给全局变量 global_this,然后filters中即可通过 global_this 获取data中数据
    },
    filters:{
      renderWorkStepName:function (workStepId) {
        return global_this.worksteps.filter(workStep => workStep.work_step_id == workStepId)[0].work_step_name;
      }
    }
  }
</script>

<style scoped>
  ul li{
    list-style: none;
  }
</style>
