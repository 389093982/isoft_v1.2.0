<template>
  <div>
    <table>
      <div v-for="(item,index) in paramInputSchemaItems">
        <tr>
          <td>
            <!--white-space: nowrap; //不换行-->
            <!--text-overflow: ellipsis; //超出部分用....代替-->
            <!--overflow: hidden; //超出隐藏-->
            <!--鼠标移动过去的时候显示全部文字,也很简单,给title赋值就可以了-->
            <div style="width: 150px;color: green;text-align: right;white-space: nowrap;text-overflow: ellipsis;overflow: hidden;"
              :title="item.ParamName">
              {{item.ParamName}}
            </div>
          </td>
          <td>
            <Icon type="ios-book-outline" size="18" style="margin-left: 3px;" @click="showParamDesc(item.ParamDesc)"/>
            <!-- transfer="true" 表示是否将弹层放置于 body 内,
              在 Tabs、带有 fixed 的 Table 列内使用时,建议添加此属性,它将不受父级样式影响,从而达到更好的效果-->
            <Select style="width: 350px;" v-if="item.ParamChoices" v-model="item.ParamValue" :transfer="true">
              <Option v-for="choice in item.ParamChoices" :value="choice" :key="choice">
                {{choice}}
              </Option>
            </Select>
            <Input style="width: 350px;" v-else size="small" v-model.trim="item.ParamValue" readonly type="text" placeholder="small size"/>
          </td>
          <td>
            <Button v-if="!item.ParamChoices" type="success" size="small" @click="handleReload(index, true)">查看/编辑</Button>
          </td>
        </tr>
        <div><span style="color: red;">{{getValidateErrors(item.ParamName)}}</span></div>
      </div>

      <ParamInputEditDialog ref="paramInputEditDialog" @handleSubmit="refreshParamInputSchemaItems"
                           :param-input-schema-items="paramInputSchemaItems" @handleReload="handleReload"/>
    </table>
  </div>
</template>

<script>
  import ParamInputEditDialog from "./ParamInputEditDialog"
  import {LoadValidateResult} from "../../../../api"

  export default {
    name: "ParamInputEdit",
    components:{ParamInputEditDialog},
    props:{
      workId: {
        type: Number,
        default: -1
      },
      workStepId: {
        type: Number,
        default: -1
      },
      paramInputSchemaItems:{
        type: Array,
        default: () => [],
      },
    },
    data(){
      return {
        validateErrors:[],
      }
    },
    methods:{
      // 根据 paramIndex 重新加载
      handleReload: function(paramIndex, refreshOutput){
        if(paramIndex >=0 && paramIndex <= this.paramInputSchemaItems.length -1){
          let item = this.paramInputSchemaItems[paramIndex];
          this.$refs["paramInputEditDialog"].refreshParamInput(paramIndex, item, refreshOutput);
        }
      },
      // 强制刷新组件
      refreshParamInputSchemaItems:function (label, text, pureText) {
        for(var i=0; i<this.paramInputSchemaItems.length; i++){
          var paramInputSchemaItem = this.paramInputSchemaItems[i];
          if(paramInputSchemaItem.ParamName == label){
            paramInputSchemaItem.ParamValue = text;
            paramInputSchemaItem.PureText = pureText;
            this.$set(this.paramInputSchemaItems, i, paramInputSchemaItem);
            this.$Message.success('临时参数保存成功!');
          }
        }
      },
      showParamDesc:function (paramDesc) {
        this.$Modal.info({
          title: "使用说明",
          content: paramDesc
        });
      },
      refreshWorkValidateDetail: async function(){
        const result = await LoadValidateResult(this.workId);
        if(result.status == "SUCCESS"){
          let validateDetails = result.details;
          this.validateErrors = validateDetails.filter(validateDetail => this.workId == validateDetail.work_id
            && this.workStepId == validateDetail.work_step_id);
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      getValidateErrors:function (paramName) {
        let _validateErrors = this.validateErrors.filter(validateDetail => validateDetail.param_name == paramName)
          .map(validateDetail => validateDetail.detail).join(",");
        return _validateErrors;
      }
    },
  }
</script>

<style scoped>

</style>
