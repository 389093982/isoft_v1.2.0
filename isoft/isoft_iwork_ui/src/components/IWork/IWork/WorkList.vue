<template>
  <div style="margin: 10px;">
    <Row type="flex" justify="center">
      <Col span="3"><IWorkDL/></Col>
      <Col span="3"><WorkValidate/></Col>
    </Row>
    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <span slot="left">
        <Button type="success" @click="addWork">新增</Button>
        <ISimpleConfirmModal ref="workEditModal" modal-title="新增/编辑 Work" :modal-width="600" :footer-hide="true">
          <IKeyValueForm ref="workEditForm" form-key-label="work_name" form-value-label="work_desc"
                         form-key-placeholder="请输入 work_name" form-value-placeholder="请输入 work_desc"
                         @handleSubmit="editWork" :formkey-validator="workNameValidator">
            <span slot="extra">
              <FormItem label="work_type">
                <Select :transfer="true" v-model="current_work_type">
                  <Option value="filter" key="filter">filter</Option>
                  <Option value="work" key="work">work</Option>
                </Select>
              </FormItem>
              <FormItem label="module_name">
                <Input v-model.trim="current_module_name" style="width: 85%;" readonly="readonly"/>
                <Poptip v-model="visible" placement="left-start" width="420">
                    <a href="javascript:;">选择模块</a>
                    <div slot="content">
                      <span v-for="module in modules" style="margin: 5px;float: left;">
                        <Tag><a @click="closePoptip(module.module_name)">{{module.module_name}}</a></Tag>
                      </span>
                    </div>
                </Poptip>
              </FormItem>
            </span>
          </IKeyValueForm>
        </ISimpleConfirmModal>
      </span>
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    所有模块：
    <span v-for="module in modules">
      <Tag><a @click="filterModuleWork(module.module_name)">{{module.module_name}}</a></Tag>
    </span>

    <Table border :columns="columns1" :data="works" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {WorkList,DeleteWorkById,RunWork,EditWork,GetAllModules} from "../../../api"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import IWorkDL from "../IWorkDL"
  import WorkValidate from "../IValidate/WorkValidate"
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import IKeyValueForm from "../../Common/form/IKeyValueForm"
  import {checkEmpty, validateCommonPatternForString} from "../../../tools/index"

  export default {
    name: "WorkList",
    components:{ISimpleLeftRightRow,ISimpleSearch,IWorkDL,WorkValidate,ISimpleConfirmModal,IKeyValueForm},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        // 搜索条件
        search:"",
        works: [],
        runLogRecordCount:{},
        current_work_type: "work",
        current_module_name: "",
        choose_module_name: "",
        modules: [],
        visible:false,
        columns1: [
          {
            title: 'work_name',
            key: 'work_name',
            width: 250,
            render: (h, params) => {
              return h('div', [
                h('span', {
                  style:{
                    marginRight: '10px',
                    color: 'red',
                  }
                }, this.works[params.index].work_type),
                h('span', this.works[params.index].work_name),
              ]);
            }
          },
          {
            title: 'module_name',
            key: 'module_name',
            width: 120,
          },
          {
            title: 'error/total',
            key: 'error/total',
            width: 100,
            render: (h, params) => {
              return h('div', [
                h('span', {
                  style:{
                    color: 'red',
                  }
                }, this.getErrorOrTotalCount(this.works[params.index].id, 'error')),
                h('span', this.getErrorOrTotalCount(this.works[params.index].id, 'total')),
              ]);
            }
          },
          {
            title: 'work_desc',
            key: 'work_desc',
            width: 400,
          },
          {
            title: '操作',
            key: 'operate',
            width: 400,
            fixed: 'right',
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
                      this.$refs.workEditModal.showModal();
                      this.$refs.workEditForm.initFormData(this.works[params.index].id, this.works[params.index].work_name, this.works[params.index].work_desc);
                      this.current_work_type = this.works[params.index].work_type;
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
                      this.deleteWorkById(this.works[params.index]['id']);
                    }
                  }
                }, '删除'),
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.download(this.works[params.index]['id']);
                    }
                  }
                }, '下载'),
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
                      this.editWorkSteps(this.works[params.index]['id'], this.works[params.index]['work_name']);
                    }
                  }
                }, '编辑步骤'),
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
                      this.runWork(this.works[params.index]['id']);
                    }
                  }
                }, '运行流程'),
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
                      this.$router.push({ path: '/iwork/runLogList', query: { work_id: this.works[params.index]['id'] }});
                    }
                  }
                }, '运行日志'),
              ]);
            }
          },
        ],
      }
    },
    methods:{
      closePoptip (module_name) {
        this.current_module_name=module_name;
        this.visible = false;
      },
      filterModuleWork:function(module_name){
        this.search = module_name;
        this.refreshWorkList();
      },
      refreshWorkList:async function () {
        const result = await WorkList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.works = result.works;
          this.total = result.paginator.totalcount;
          this.runLogRecordCount = result.runLogRecordCount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshWorkList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshWorkList();
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshWorkList();
      },
      download:function(id){
        window.location.href = "/api/iwork/download/" + id;
      },
      deleteWorkById:async function(id){
        const result = await DeleteWorkById(id);
        if(result.status=="SUCCESS"){
          this.refreshWorkList();
        }
      },
      addWork:function(){
        this.$refs.workEditModal.showModal();
      },
      editWork:async function (work_id, work_name, work_desc) {
        if(checkEmpty(this.current_module_name)){
          this.$Message.error("请选择模块！");
          return;
        }
        const result = await EditWork(work_id, work_name, work_desc, this.current_work_type, this.current_module_name);
        if(result.status == "SUCCESS"){
          this.$refs.workEditModal.hideModal();
          this.$refs.workEditForm.handleSubmitSuccess("提交成功!");
          this.refreshWorkList();
        }else{
          this.$refs.workEditForm.handleSubmitError("提交失败!");
        }
      },
      runWork:async function (work_id) {
        const result = await RunWork(work_id);
        if(result.status == "SUCCESS"){
          this.$Message.success("运行任务已触发!");
        }
      },
      editWorkSteps:function (id, work_name) {
        this.$router.push({ path: '/iwork/workstepList', query: { work_id: id, work_name: work_name }});
      },
      workNameValidator (rule, value, callback){
        if (value === '') {
          callback(new Error('字段值不能为空!'));
        } else if (!validateCommonPatternForString(value)) {
          callback(new Error('存在非法字符，只能包含字母，数字，下划线!'));
        } else {
          callback();
        }
      },
      getErrorOrTotalCount:function (workId, flag) {
        var key = Object.keys(this.runLogRecordCount).filter(function (key) {
          return key == workId;
        })[0];
        return flag == "error" ? this.runLogRecordCount[key].errorCount : "/" + this.runLogRecordCount[key].allCount;
      },
      refreshAllModules:async function(){
        const result = await GetAllModules();
        if(result.status == "SUCCESS"){
          this.modules = result.moudles;
        }
      }
    },
    mounted: function () {
      this.refreshWorkList();
      this.refreshAllModules();
    },
  }
</script>

<style scoped>

</style>
