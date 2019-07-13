<template>
  <Modal
    v-model="showFormModal"
    width="1200"
    title="查看/编辑 workstep 参数"
    :footer-hide="true"
    :transfer="false"
    :mask-closable="false"
    :styles="{top: '20px'}">
    <Row>
      <Col span="6">
        <ParamInputEditDataSource ref="paramInputEditDataSource" :pre-pos-tree-node-arr="prePosTreeNodeArr"/>
      </Col>
      <Col span="3" style="text-align: center;">
        <i-circle :percent="editProgress" :size="85">
          编辑进度
          <span class="demo-Circle-inner" style="font-size:24px">{{editProgress}}%</span>
        </i-circle>

        <Button @click="appendData('parent')" style="margin-top: 10px;"><Icon type="ios-arrow-forward"></Icon>选择父节点</Button>
        <Button @click="appendData('children')" style="margin-top: 10px;"><Icon type="ios-arrow-forward"></Icon>选择子节点</Button>

        <Scroll height="250" style="margin: 20px 20px 0px 20px;">
          <ul>
            <li style="list-style: none;margin: 2px;" v-for="(item,index) in paramInputSchemaItems">
              <Button :type="item.ParamValue ? 'success' : 'default'" long size="small" @click="handleReload(index)">{{ item.ParamName }}</Button>
            </li>
          </ul>
        </Scroll>
      </Col>
      <Col span="15">
        <div class="operate_link">
          <TemplateChooser ref="templateModal" @chooseTemplate="chooseTemplate"/>
          <QuickFuncList ref="quickFuncList"/>
          <ul style="text-align: right;">
            <li>
              <h3 style="color: #1600ff;">参数({{paramIndex}}):{{inputLabel}}</h3>
            </li>
            <li>
              <Checkbox v-model="pureText">纯文本值</Checkbox>
            </li>
            <li>
              <Button type="info" size="small" @click="parseToMultiValue()">多值预览</Button>
            </li>
            <li>
              <Button type="success" size="small" @click="showQuickFunc()">快捷函数</Button>
            </li>
            <li>
              <router-link :to="{ path: '/iwork/quickSql' }" tag="a" target="_blank">
                <Button type="warning" size="small">辅助工具</Button>
              </router-link>
            </li>
            <li>
              <Button type="success" size="small" @click="$refs.templateModal.showModal()">模板文字</Button>
            </li>
          </ul>
        </div>
        <div v-if="showMultiVals" style="margin-top: 20px;">
          <Scroll height="350">
            <table style="width: 100%;">
              <tr v-for="(val,index) in multiVals">
                <td style="width: 10%;">参数 {{index}}</td>
                <td><Input type="textarea" :value="val" :readonly="true"/></td>
              </tr>
            </table>

          </Scroll>
        </div>
        <div v-else>
          <Input v-model="inputTextData" type="textarea" :rows="12" placeholder="Enter something..."
                 @drop.native="handleInputDrop" @dragover.native="handleDragover"/>
          <p style="margin-top: 5px;">变量占位符</p>
          <Scroll height="100">
            <Tag color="default" v-for="(variable,index) in variables" style="margin-right: 10px;"
                 @drop.native="handlePlaceholderDrop($event, index)" @dragover.native="handleDragover">{{variable}}</Tag>
          </Scroll>

          <Row style="text-align: right;margin-top: 10px;">
            <Button type="success" size="small" @click="handleSubmit">Submit</Button>
            <Button type="info" size="small" @click="closeModal">Close</Button>
          </Row>
        </div>
      </Col>
    </Row>
  </Modal>
</template>

<script>
  import {LoadPreNodeOutput} from "../../../../api"
  import ISimpleBtnTriggerModal from "../../../Common/modal/ISimpleBtnTriggerModal"
  import QuickFuncList from "../../IQuickFunc/QuickFuncList"
  import {ParseToMultiValue} from "../../../../api"
  import TemplateChooser from "./TemplateChooser"
  import {getMatchArrForString} from "../../../../tools"
  import ParamInputEditDataSource from "./ParamInputEditDataSource"
  import {checkEmpty} from "../../../../tools"

  export default {
    name: "ParamInputEditDialog",
    components:{ISimpleBtnTriggerModal,QuickFuncList,TemplateChooser,ParamInputEditDataSource},
    props:{
      paramInputSchemaItems:{
        type: Array,
        default: () => [],
      },
    },
    data(){
      return {
        showFormModal:false,
        inputLabel:'',
        oldPureText:false,
        pureText:false,
        oldInputTextData:'',
        inputTextData:'',
        showMultiVals:false,  // 默认非多值视图
        multiVals:[],         // 存储多值列表
        paramIndex:1,
        prePosTreeNodeArr:[],
        variables:[],
        variableConcats:[],
        editProgress: 0,
      }
    },
    methods:{
      handlePlaceholderDrop:function(event, index){
        // 取消冒泡
        event.stopPropagation();
        event.preventDefault();
        var transferText = event.dataTransfer.getData("Text");
        this.variables[index] =  transferText.substr(0, transferText.lastIndexOf(";\n"));     // 将值替换进 variables
        var _inputTextData = "";
        for(var i=0; i<this.variableConcats.length; i++){
          _inputTextData += this.variableConcats[i] + (i == this.variableConcats.length - 1 ? "" : this.variables[i]);
        }
        this.inputTextData = _inputTextData;
      },
      handleInputDrop:function(){
        const event = window.event||arguments[0];
        // 取消冒泡
        event.stopPropagation();
        event.preventDefault();
        var transferText = event.dataTransfer.getData("Text");
        // 将数据添加到右侧
        this.inputTextData = this.inputTextData + transferText;
      },
      handleDragover:function(){
        const event = window.event||arguments[0];
        event.preventDefault();
      },
      parseToMultiValue: async function(){
        const result = await ParseToMultiValue(this.pureText, this.inputTextData);
        if(result.status == "SUCCESS"){
          this.showMultiVals = !this.showMultiVals;
          this.multiVals = result.multiVals;
        }else{
          this.$Message.error('提交失败!' + result.errorMsg);
        }
      },
      handleReload: function(paramIndex){
        var _this = this;
        if(this.checkDrity()){
          this.$Modal.confirm({
            title:"确认",
            content:"是否需要保存上一步操作?",
            onOk: function () {
              _this.handleSubmit(function () {
                _this.$emit("handleReload", paramIndex);
              });
            },
          });
        }else{
          this.$emit("handleReload", paramIndex);
        }
      },
      refreshEditProgress:function(){
        var count = 0;
        for(var i=0; i<this.paramInputSchemaItems.length; i++){
          if(!checkEmpty(this.paramInputSchemaItems[i].ParamValue) || this.paramInputSchemaItems[i].ParamName.indexOf("?") > 0){
            count ++;
          }
        }
        this.editProgress = Math.floor(count / this.paramInputSchemaItems.length * 100);
      },
      refreshParamInput: function(index, item){
        this.showFormModal = true;
        this.paramIndex = index;
        this.inputLabel = item.ParamName;
        this.pureText = item.PureText;
        // 文本输入框设置历史值
        this.inputTextData = item.ParamValue;
        this.showMultiVals = false;
        this.clearDirty();
        this.refreshPreNodeOutput();
        this.refreshEditProgress();
      },
      closeModal: function(){
        this.showFormModal = false;
      },
      showQuickFunc: function(){
        this.$refs.quickFuncList.showModal();
      },
      clearDirty: function (){
        this.oldInputTextData = this.inputTextData;
        this.oldPureText = this.pureText;
      },
      checkDrity: function(){
        return this.oldInputTextData != this.inputTextData || this.oldPureText != this.pureText;
      },
      handleSubmit:function (callback) {
        this.$emit("handleSubmit", this.inputLabel, this.inputTextData, this.pureText);
        this.clearDirty();
        if(callback != null && callback != undefined){
          callback();     // 提交成功后的回调函数
        }
      },
      refreshPreNodeOutput:async function () {
        const result = await LoadPreNodeOutput(this.$store.state.current_work_id, this.$store.state.current_work_step_id);
        if(result.status == "SUCCESS"){
          this.prePosTreeNodeArr = result.prePosTreeNodeArr;
        }
      },
      appendData:function (chooseType) {
        let datas = this.$refs.paramInputEditDataSource.getChooseDatas(chooseType);
        if(Array.isArray(datas)){
          for(var i=0; i<datas.length; i++){
            this.inputTextData = this.inputTextData + datas[i];
          }
        }else{
          this.inputTextData = this.inputTextData + datas;
        }
      },
      chooseTemplate:function (template) {
        this.inputTextData = this.inputTextData + template.template_value;
        this.$refs.templateModal.hideModal();
      },
    },
    watch: {
      inputTextData(val) {
        // 所有占位符变量
        this.variables = getMatchArrForString(this.inputTextData, /\$[a-zA-Z0-9]+.[a-zA-Z0-9]+/g);
        this.variableConcats = this.inputTextData.split(/\$[a-zA-Z0-9]+.[a-zA-Z0-9]+/g);
      }
    }
  }
</script>

<style scoped>
  .operate_link ul li{
   display: inline-block;
   margin-left: 10px;
   margin-bottom: 5px;
 }
</style>
