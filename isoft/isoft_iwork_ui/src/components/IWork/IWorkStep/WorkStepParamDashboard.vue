<template>
  <span v-if="paramInputSchema != null && paramOutputSchema != null">
    <Card style="width:350px;margin: 10px;">
      <p slot="title">
        <Icon type="ios-film-outline"></Icon>
        {{workstep.work_step_name}}
      </p>
      <a href="#" slot="extra" @click.prevent="handleChange(workstep)">
        <Icon type="ios-loop-strong"></Icon>
        Change
      </a>

      <Row>
        <Col span="12">
          <ul v-for="pis in paramInputSchema.ParamInputSchemaItems">
            <li style="list-style: none;">
              <Tooltip max-width="400" theme="light">
                {{pis.ParamName}}
                <div slot="content" style="word-break: break-all;">{{pis.ParamValue}}</div>
              </Tooltip>
            </li>
          </ul>
        </Col>
        <Col span="12">
          <ul v-for="pos in paramOutputSchema.ParamOutputSchemaItems">
            <li style="list-style: none;">
              <Tooltip max-width="400" theme="light">
                {{pos.ParamName}}
                <div slot="content" style="word-break: break-all;">{{pos.ParamValue}}</div>
              </Tooltip>
            </li>
          </ul>
        </Col>
      </Row>
    </Card>
  </span>
</template>

<script>
  import {LoadWorkStepInfo} from "../../../api/index"

  export default {
    name: "WorkStepParamDashboard",
    props:{
      workstep:{
        type:Object,
        default:null,
      }
    },
    data(){
      return {
        paramInputSchema:null,
        paramOutputSchema:null,
      }
    },
    methods:{
      handleChange:function (workstep) {
        alert(workstep.work_step_input.ParamInputSchemaItems);
      },
      loadWorkStepInfo:async function () {
        const result = await LoadWorkStepInfo(this.workstep.work_id,this.workstep.work_step_id);
        if(result.status == "SUCCESS") {
          // 入参渲染
          this.paramInputSchema = result.paramInputSchema;
          // 出参渲染
          this.paramOutputSchema = result.paramOutputSchema;
        }
      }
    },
    mounted:function () {
      this.loadWorkStepInfo();
    }
  }
</script>

<style scoped>

</style>
