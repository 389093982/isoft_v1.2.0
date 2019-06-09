<template>
  <div style="margin: 10px;">
    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <span slot="left">
        <Button type="success" @click="addTemplate">新增</Button>
        <ISimpleConfirmModal ref="templateEditModal" modal-title="新增/编辑 template" :modal-width="600" :footer-hide="true">
          <IThemeKeyValueForm ref="templateEditForm" form-theme-label="template_theme" form-theme-placeholder="请选择模板"
                         :formthemes="['template_string','template_var']"
                         form-key-label="template_name" form-value-label="template_value"
                         form-key-placeholder="请输入 template_name" form-value-placeholder="请输入 template_value"
                         @handleSubmit="editTemplate" :formkey-validator="templateNameValidator"/>
        </ISimpleConfirmModal>
      </span>

      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>
  </div>
</template>

<script>
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import IThemeKeyValueForm from "../../Common/form/IThemeKeyValueForm"
  import {validateCommonPatternForString} from "../../../tools/index"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"

  export default {
    name: "Template",
    components:{ISimpleLeftRightRow,ISimpleConfirmModal,IThemeKeyValueForm,ISimpleSearch},
    data(){
      return {

      }
    },
    methods:{
      addTemplate:function(){
        this.$refs.templateEditModal.showModal();
      },
      editTemplate:async function (id, template_theme, template_name, template_value) {
       alert(1234);
      },
      templateNameValidator (rule, value, callback){
        if (value === '') {
          callback(new Error('字段值不能为空!'));
        } else if (!validateCommonPatternForString(value)) {
          callback(new Error('存在非法字符，只能包含字母，数字，下划线!'));
        } else {
          callback();
        }
      },
      handleSearch(data){
      },
    }
  }
</script>

<style scoped>

</style>
