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

    <Table border :columns="columns1" :data="templates" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import IThemeKeyValueForm from "../../Common/form/IThemeKeyValueForm"
  import {validateCommonPatternForString} from "../../../tools/index"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import {EditTemplate} from "../../../api"
  import {TemplateList} from "../../../api"
  import {DeleteTemplateById} from "../../../api"

  export default {
    name: "Template",
    components:{ISimpleLeftRightRow,ISimpleConfirmModal,IThemeKeyValueForm,ISimpleSearch},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        // 搜索条件
        search:"",
        templates: [],
        columns1: [
          {
            title: 'id',
            key: 'id',
          },
          {
            title: 'template_theme',
            key: 'template_theme',
          },
          {
            title: 'template_name',
            key: 'template_name',
          },
          {
            title: 'template_value',
            key: 'template_value',
          },
          {
            title: '操作',
            key: 'operate',
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.$refs.templateEditModal.showModal();
                      this.$refs.templateEditForm.initFormData(this.templates[params.index].id, this.templates[params.index].template_theme, this.templates[params.index].template_name, this.templates[params.index].template_value);
                    }
                  }
                }, '编辑'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.deleteTemplate(this.templates[params.index].id);
                    }
                  }
                }, '删除'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      addTemplate:function(){
        this.$refs.templateEditModal.showModal();
      },
      editTemplate:async function (id, template_theme, template_name, template_value) {
        const result = await EditTemplate(id, template_theme, template_name, template_value);
        if(result.status == "SUCCESS"){
          this.$refs.templateEditForm.handleSubmitSuccess("提交成功!");
          this.$refs.templateEditModal.hideModal();
          this.refreshTemplateList();
        }else{
          this.$refs.globalVarForm.handleSubmitError("提交失败!");
        }

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
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshTemplateList();
      },
      handleChange(page){
        this.current_page = page;
        this.refreshTemplateList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshTemplateList();
      },
      refreshTemplateList: async function () {
        const result = await TemplateList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.templates = result.templates;
          this.total = result.paginator.totalcount;
        }
      },
      deleteTemplate:async function (id){
        const result = await DeleteTemplateById(id);
        if(result.status == "SUCCESS"){
          this.refreshTemplateList();
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
    },
    mounted:function () {
      this.refreshTemplateList();
    }
  }
</script>

<style scoped>

</style>
