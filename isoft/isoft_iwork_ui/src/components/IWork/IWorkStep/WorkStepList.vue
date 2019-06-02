<template>
  <div style="margin: 10px;">
    <h4 v-if="$route.query.work_name" style="text-align: center;margin-bottom: 10px;">当前流程为：{{$route.query.work_name}}</h4>

    <Row type="flex" justify="start" class="code-row-bg" style="margin-bottom: 20px;">
      <Col span="2"><Button type="error" size="small" @click="showComponet = !showComponet" style="margin-right: 5px;">显示组件</Button></Col>
      <Col span="2"><Button type="warning" size="small" @click="showRefactorModal">重构流程</Button></Col>
      <Col span="2"><Button type="info" size="small" @click="batchChangeIndent('left', null)">向左缩进</Button></Col>
      <Col span="2"><Button type="error" size="small" @click="batchChangeIndent('right', null)">向右缩进</Button></Col>
      <Col span="2"><Button type="info" size="small" @click="flushCache">刷新缓存</Button></Col>
      <Col span="2"><Button type="warning" size="small" @click="runWork">运行流程</Button></Col>
      <Col span="2"><Button type="info" size="small" @click="showRunLogList">运行日志</Button></Col>
      <Col span="2"><WorkValidate /></Col>
      <Col span="2"><Button type="error" size="small" @click="renderSourceXml">View XML</Button></Col>

      <ISimpleConfirmModal ref="refactor_modal" modal-title="重构为子流程" :modal-width="500" @handleSubmit="refactor">
        <Input v-model.trim="refactor_worksub_name" placeholder="请输入重构的子流程名称"></Input>
      </ISimpleConfirmModal>
    </Row>
    <BaseInfo ref="workStepBaseInfo" @reloadWorkStepBaseInfo="showWorkStepBaseInfo" @handleSuccess="refreshWorkStepList" :worksteps="worksteps"/>
    <ParamInfo ref="workStepParamInfo" @handleSuccess="refreshWorkStepList"/>

    <Row type="flex">
      <Col v-if="showComponet" span="6">
        <Scroll height="500">
          <span v-for="default_work_step_type in default_work_step_types" style="margin: 5px;float: left;"
                draggable="true" @dragstart="dragstart($event, default_work_step_type.name)">
           <Tag>{{default_work_step_type.name}}</Tag>
          </span>
        </Scroll>
      </Col>
      <Col :span="showComponet ? 18 : 24">

        <Table :height="500" border :columns="columns1" ref="selection" :data="worksteps" size="small"></Table>
      </Col>
    </Row>

    <!-- 相关流程清单 -->
    <RelativeWork id="relativeWork" ref="relativeWork"/>
  </div>
</template>

<script>
  import {WorkStepList} from "../../../api/index"
  import {DeleteWorkStepByWorkStepId} from "../../../api/index"
  import {ChangeWorkStepOrder} from "../../../api/index"
  import {RefactorWorkStepInfo} from "../../../api/index"
  import {BatchChangeIndent} from "../../../api/index"
  import {AddWorkStep} from "../../../api/index"
  import {RunWork} from "../../../api/index"
  import ParamInfo from "./ParamInfo/ParamInfo"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import BaseInfo from "./BaseInfo/BaseInfo"
  import RelativeWork from "./RelativeWork/RelativeWork"
  import {oneOf} from "../../../tools/index"
  import {checkContainsInString} from "../../../tools/index"
  import WorkValidate from "../IValidate/WorkValidate"
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import {getRepeatStr} from "../../../tools/index"
  import {GetRelativeWork} from "../../../api/index"
  import {FlushCache} from "../../../api/index"
  import {EditWorkStepBaseInfo} from "../../../api/index"

  export default {
    name: "WorkStepList",
    components:{ParamInfo,ISimpleLeftRightRow,BaseInfo,RelativeWork,WorkValidate,ISimpleConfirmModal},
    data(){
      return {
        // 默认不显示组件
        showComponet:false,
        showRelativeWorkFlag:false,
        refactor_worksub_name:'',
        default_work_step_types: this.GLOBAL.default_work_step_types,
        worksteps: [],
        columns1: [
          {
            type: 'selection',
            width: 60,
            align: 'center',
          },
          {
            title: '步骤编号',
            key: 'work_step_id',
            width: 300,
            render: (h,params)=>{
              return h('div', {
                  on:{
                    drop: () => {
                      const event = window.event||arguments[0];
                      // 取消冒泡
                      event.stopPropagation();
                      event.preventDefault();
                      var work_step_type = event.dataTransfer.getData("Text");
                      this.addWorkStep(params.row.work_step_id, work_step_type);
                    },
                    dragover: () => this.allowDrop(),
                  }
                }, [
                  h('span', params.row.work_step_id),
                  h('Icon', {
                    props: {
                      type: 'md-arrow-round-up',
                      size: 15,
                    },
                    style: {
                      marginLeft: '5px',
                      display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"]) ? undefined : 'none',
                    },
                    on: {
                      click: () => {
                        this.changeWorkStepOrder(this.worksteps[params.index]['work_step_id'], "up");
                      },
                    }
                  }),
                  h('Icon', {
                    props: {
                      type: 'md-arrow-round-down',
                      size: 15,
                    },
                    style: {
                      marginLeft: '5px',
                      display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"]) ? undefined : 'none',
                    },
                    on: {
                      click: () => {
                        this.changeWorkStepOrder(this.worksteps[params.index]['work_step_id'], "down");
                      }
                    }
                  }),
                  h('Icon', {
                    props: {
                      type: 'md-arrow-round-back',
                      size: 15,
                    },
                    style: {
                      marginLeft: '5px',
                      display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"]) ? undefined : 'none',
                    },
                    on: {
                      click: () => {
                        this.batchChangeIndent('left', [this.worksteps[params.index]['work_step_id']]);
                      }
                    }
                  }),
                  h('Icon', {
                    props: {
                      type: 'md-arrow-round-forward',
                      size: 15,
                    },
                    style: {
                      marginLeft: '5px',
                      display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"]) ? undefined : 'none',
                    },
                    on: {
                      click: () => {
                        this.batchChangeIndent('right', [this.worksteps[params.index]['work_step_id']]);
                      }
                    }
                  }),
                  h('Button', {
                    props: {
                      type: 'error',
                      size: 'small'
                    },
                    style: {
                      marginLeft: '5px',
                      display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"])  ? undefined : 'none',
                    },
                    on: {
                      click: () => {
                        this.showWorkStepBaseInfo(this.worksteps[params.index]['work_step_id']);
                      },
                    }
                  }, '编辑'),
                  h('Button', {
                    props: {
                      type: 'info',
                      size: 'small'
                    },
                    style: {
                      marginLeft: '5px',
                    },
                    on: {
                      click: () => {
                        if (this.worksteps[params.index]['work_step_type']){
                          this.$refs.workStepParamInfo.showWorkStepParamInfo(this.$route.query.work_id, this.worksteps[params.index]['work_step_id']);
                        }
                      }
                    }
                  }, '参数'),
                  h('Button', {
                    props: {
                      type: 'success',
                      size: 'small'
                    },
                    style: {
                      marginLeft: '5px',
                      display: !oneOf(this.worksteps[params.index]['work_step_type'], ["work_start","work_end"])  ? undefined : 'none'
                    },
                    on: {
                      click: () => {
                        this.deleteWorkStepByWorkStepId(this.worksteps[params.index]['work_id'], this.worksteps[params.index]['work_step_id']);
                      }
                    }
                  }, '删除'),
                  h('Button', {
                    props: {
                      type: 'warning',
                      size: 'small'
                    },
                    style: {
                      marginLeft: '5px',
                      display: oneOf(this.worksteps[params.index]['work_step_type'], ["work_sub"]) ? undefined : 'none'
                    },
                    on: {
                      click: () => {
                        this.showWorkSubDetail(this.worksteps[params.index]);
                      }
                    }
                  }, '详情'),
                  h('Button', {
                    props: {
                      type: 'warning',
                      size: 'small'
                    },
                    style: {
                      marginLeft: '5px',
                      display: oneOf(this.worksteps[params.index]['work_step_type'], ["work_start"]) ? undefined : 'none'
                    },
                    on: {
                      click: () => {
                        this.goAnchor('#relativeWork');
                      }
                    }
                  }, '关联流程'),
                ]
              )
            }
          },
          {
            title: 'work_step_name',
            key: 'work_step_name',
            width: 250,
            render: (h, params) => {
              var _this = this; // vue 实例
              // 可编辑模式
              if(params.row.$isEdit && !oneOf(params.row.work_step_name, ["start", "end"])){
                return h('input', {
                  domProps: {
                    value: params.row.work_step_name,
                  },
                  on: {
                    input: function (event) {
                      params.row.work_step_name = event.target.value
                    },
                    blur: async function (event) {
                      // 发生过修改
                      const result = await EditWorkStepBaseInfo(params.row.work_id, params.row.work_step_id,
                        params.row.work_step_name,params.row.work_step_desc, params.row.work_step_type, params.row.is_defer);
                      if(result.status == "SUCCESS"){
                        _this.$Message.success('修改成功!');
                      }else{
                        _this.$Message.error('修改失败！');
                      }
                      // 刷新组件
                      _this.refreshWorkStepList();
                    }
                  }
                });
              }else{
                // 非可编辑模式
                return h('div', [
                  h('span', {
                    style: {
                      // work_step_name 根据缩进级别进行缩进,不同级别使用不同颜色
                      color: ['green','blue','grey','red'][params.row.work_step_indent],
                    },
                    on: {
                      // click 变成可编辑模式
                      click:function (event) {
                        var workstep = _this.worksteps[params.index];
                        workstep.$isEdit = true;
                        _this.$set(_this.worksteps[params.index], workstep);  // 刷新界面
                      }
                    }
                  }, getRepeatStr('\xa0\xa0\xa0\xa0\xa0', params.row.work_step_indent) + this.worksteps[params.index]['work_step_name']),
                ]);
              }
            }
          },
          {
            title: 'work_step_type',
            key: 'work_step_type',
            width: 180,
            render: (h, params) => {
              return h('div', [
                h('Icon', {
                  props: {
                    type: this.renderWorkStepTypeIcon(this.worksteps[params.index]['work_step_type']),
                    size: 25,
                  },
                  style: {
                    marginRight: '5px',
                  },
                }),
                h('span', this.worksteps[params.index]['work_step_type']),
                h('Badge', {    // 延迟执行函数显示效果
                  props: {
                    status: "error",
                  },
                  style: {
                    marginLeft: '5px',
                    display: oneOf(this.worksteps[params.index]['is_defer'], ["true"])  ? undefined : 'none',
                  },
                }),
              ]);
            }
          },
          {
            title: 'work_step_desc',
            key: 'work_step_desc',
            width: 250,
          },
        ],
      }
    },
    methods:{
      refreshWorkStepList:async function () {
        const result = await WorkStepList(this.$route.query.work_id);
        if(result.status=="SUCCESS"){
          this.worksteps = result.worksteps;
          // 刷新关联流程信息
          this.$refs.relativeWork.refreshRelativeWork(this.$route.query.work_id);
        }
      },
      deleteWorkStepByWorkStepId:async function(work_id, work_step_id){
        const result = await DeleteWorkStepByWorkStepId(work_id, work_step_id);
        if(result.status=="SUCCESS"){
          this.refreshWorkStepList();
        }
      },
      changeWorkStepOrder:async function(work_step_id, type){
        const result = await ChangeWorkStepOrder(this.$route.query.work_id, work_step_id, type);
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
          this.$Message.success('换位成功!');
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      renderSourceXml:function () {
        alert(11111);
      },
      addWorkStep:async function (work_step_id, work_step_type) {
        const result = await AddWorkStep(this.$route.query.work_id, work_step_id, work_step_type);
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
          this.$Message.success('添加成功!');
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      renderWorkStepTypeIcon:function (workStepType) {
        for(var i=0; i<this.default_work_step_types.length; i++){
          let default_work_step_type = this.default_work_step_types[i];
          if(default_work_step_type.name == workStepType){
            return default_work_step_type.icon;
          }
        }
      },
      showRefactorModal:function (){
        this.$refs.refactor_modal.showModal();
      },
      getSelectionArr:function(){
        let selectionArr = [];
        let selections = this.$refs.selection.getSelection();
        for(var i=0; i<selections.length; i++){
          selectionArr.push(selections[i].work_step_id);
        }
        return selectionArr;
      },
      refactor: async function () {
        let selections = this.getSelectionArr();
        const result = await RefactorWorkStepInfo(this.$route.query.work_id, this.refactor_worksub_name, JSON.stringify(selections));
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      batchChangeIndent:async function (mod, selections) {
        if (selections == null){
          selections = this.getSelectionArr();
        }
        const result = await BatchChangeIndent(this.$route.query.work_id, mod, JSON.stringify(selections));
        if(result.status == "SUCCESS"){
          this.refreshWorkStepList();
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      showWorkSubDetail:async function (currentworkStep) {
        const result = await GetRelativeWork(currentworkStep['work_id']);
        const work_subs = result.subworks.filter(subWork => subWork.id === currentworkStep['work_sub_id']);
        if(work_subs.length > 0){
          this.$router.push({ path: '/iwork/workstepList',
            query: { work_id: work_subs[0].id, work_name: work_subs[0].work_name }});
        }
      },
      showRunLogList: function(){
        this.$router.push({ path: '/iwork/runLogList', query: { work_id: this.$route.query.work_id }});
      },
      runWork: async function(){
        const result = await RunWork(this.$route.query.work_id);
        if(result.status == "SUCCESS"){
          this.$Message.success("运行任务已触发!");
        }
      },
      flushCache: async function(){
        const result = await FlushCache();
        if(result.status == "SUCCESS"){
          this.$Message.success("刷新成功！");
        }
      },
      // 前往锚点方法
      goAnchor: function(selector) {
        var anchor = this.$el.querySelector(selector);
        document.documentElement.scrollTop = anchor.offsetTop;
      },
      dragstart:function(event, transferData){
        event.dataTransfer.setData("Text", transferData);
      },
      allowDrop:function(){
        const event = window.event||arguments[0];
        event.preventDefault();
      },
      showWorkStepBaseInfo:function (work_step_id) {
        this.$refs.workStepBaseInfo.showWorkStepBaseInfo(this.$route.query.work_id, work_step_id);
      }
    },
    mounted: function () {
      this.refreshWorkStepList();
    },
    watch:{
      // 监听路由是否变化
      '$route' (to, from) {
        this.refreshWorkStepList();
      }
    }
  }
</script>

<style scoped>

</style>
