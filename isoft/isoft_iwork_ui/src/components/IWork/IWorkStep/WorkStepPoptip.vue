<template>
  <div v-if="workstep" style="width: 450px;padding: 10px;">
    <div style="margin-bottom: 10px;">
      <Button type="success" size="small">步骤名称：{{workstep.work_step_name}}</Button>
      <Button type="success" size="small" style="float: right;">高亮显示引用关系</Button>
    </div>
    <Row>
      <Col span="12">输入参数</Col>
      <Col span="12">输出参数</Col>
      <Col span="12" v-if="ParamInputSchemaItems">
        <ul>
          <li v-for="item in ParamInputSchemaItems">
            <Tag>{{item.ParamName}}</Tag>
          </li>
        </ul>
      </Col>
      <Col span="12" v-if="ParamInputSchemaItems">
        <ul>
          <li v-for="item in ParamOutputSchemaItems">
            <Tag>{{item.ParamName}}</Tag>
          </li>
        </ul>
      </Col>
    </Row>
  </div>
</template>

<script>
  import {checkEmpty} from "../../../tools"

  export default {
    name: "WorkStepPoptip",
    props:{
      workstep:{
        type: Object,
        default: null,
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
  }
</script>

<style scoped>
  ul li{
    list-style: none;
  }
</style>
