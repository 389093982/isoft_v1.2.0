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
      <Col span="7">
        <Tabs type="card" :value="current_tab" @on-click="currentTabChanged">
          <TabPane label="前置节点输出参数" name="tab_output">
            <Scroll height="350">
              <Tree :data="data1" show-checkbox ref="tree1" :render="renderContent"></Tree>
            </Scroll>
          </TabPane>
          <TabPane label="快捷函数" name="tab_funcs">
            <Scroll height="350">
              <ul>
                <Tree :data="data2" show-checkbox ref="tree2" :render="renderContent"></Tree>
              </ul>
            </Scroll>
          </TabPane>
        </Tabs>
      </Col>
      <Col span="3" style="text-align: center;margin-top: 100px;">
        <Button @click="appendData('parent')" style="margin-top: 10px;"><Icon type="ios-arrow-forward"></Icon>选择父节点</Button>
        <Button @click="appendData('children')" style="margin-top: 10px;"><Icon type="ios-arrow-forward"></Icon>选择子节点</Button>
      </Col>
      <Col span="14">
        <div class="operate_link">
          <ul>
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
              <QuickFuncList ref="quickFuncList" @chooseFunc="chooseFunc"/>
              <Button type="success" size="small" @click="showQuickFunc()">快捷函数</Button>
            </li>
            <li>
              <router-link :to="{ path: '/iwork/quickSql' }" tag="a" target="_blank">
                <Button type="warning" size="small">辅助工具</Button>
              </router-link>
            </li>
            <li>
              <TemplateChooser ref="templateModal" @chooseTemplate="chooseTemplate"/>
              <Button type="success" size="small" @click="$refs.templateModal.showModal()">模板文字</Button>
            </li>
          </ul>
        </div>
        <div v-show="showMultiVals" style="margin-top: 20px;">
          <Scroll height="350">
            <table style="width: 100%;">
              <tr v-for="(val,index) in multiVals">
                <td style="width: 10%;">参数 {{index}}</td>
                <td><Input type="textarea" :value="val" :readonly="true"/></td>
              </tr>
            </table>

          </Scroll>
        </div>
        <Input v-show="showMultiVals == false" v-model="inputTextData" type="textarea" :rows="15" placeholder="Enter something..."
               @drop.native="handleDrop" @dragover.native="handleDragover"/>
        <div style="padding: 10px;">
          占位符：<Tag color="default" v-for="variable in variables" style="margin-right: 10px;">{{variable}}</Tag>
        </div>
      </Col>
    </Row>
    <Row style="text-align: right;margin-top: 10px;">
      <Button type="success" size="small" @click="handleSubmit">Submit</Button>
      <Button type="error" size="small" @click="showNext(-1)">Edit Last</Button>
      <Button type="warning" size="small" @click="showNext(1)">Edit Next</Button>
      <Button type="info" size="small" @click="closeModal">Close</Button>
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

  export default {
    name: "ParamInputEditDialog",
    components:{ISimpleBtnTriggerModal,QuickFuncList,TemplateChooser},
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
        funcs:this.GLOBAL.quick_funcs,
        current_tab:'tab_output',
        variables:[],
      }
    },
    methods:{
      renderContent (h, { root, node, data }) {
        return h('span', {
          style: {
            display: 'inline-block',
            width: '100%'
          },
          attrs: {
            draggable:'true'
          },
          on:{
            dragstart: () => this.handleDragStart(root, node, data),
          }
        }, [
          h('span', [
            h('Icon', {
              props: {
                type: 'ios-paper-outline'
              },
              style: {
                marginRight: '8px'
              }
            }),
            h('span', data.title)
          ]),
        ]);
      },
      handleDragStart(root, node, data){
        const event = window.event||arguments[0];
        // 获取父节点
        var getParent = function (_root, _node) {
          if(_root.find(el => el === _node)){
            const parentKey = _root.find(el => el === _node).parent;
            if(parentKey != null){
              parent = _root.find(el => el.nodeKey === parentKey);
              return parent;
            }
          }
          return null;
        };
        // 获取当前节点对应的 title
        var getCurrentTitle = function (_root, _node) {
          var parent = getParent(_root, _node);
          if(parent == null){
            return _node.node.title;
          }
          return getCurrentTitle(_root, parent) + "." + _node.node.title;
        };
        var data = getCurrentTitle(root, node)  + ";\n";
        // 传递数据
        event.dataTransfer.setData("Text", data);
      },
      handleDrop:function(){
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
        this.$emit("handleReload", paramIndex);
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
      },
      showNext: function(num){
        var _this = this;
        if(this.checkDrity()){
          this.$Modal.confirm({
            title:"确认",
            content:"是否需要保存上一步操作?",
            onOk: function () {
              _this.handleSubmit();
            },
          });
        }else{
          this.handleReload(this.paramIndex + num);
        }
      },
      closeModal: function(){
        this.showFormModal = false;
      },
      chooseFunc: function(funcDemo){
        // 将数据复制到右侧
        this.inputTextData = this.inputTextData + funcDemo + "\n";
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
      handleSubmit:function () {
        this.$emit("handleSubmit", this.inputLabel, this.inputTextData, this.pureText);
        this.clearDirty();
      },
      refreshPreNodeOutput:async function () {
        const result = await LoadPreNodeOutput(this.$store.state.current_work_id, this.$store.state.current_work_step_id);
        if(result.status == "SUCCESS"){
          this.prePosTreeNodeArr = result.prePosTreeNodeArr;
        }
      },
      appendDataWithPrefix:function(prefix, item, chooseType){
        // 没有子节点
        if(item.children == null){
          if(item.indeterminate == false){
            if(item.checked){
              // 将数据添加到右侧
              this.inputTextData = this.inputTextData + prefix + ";\n";
            }
          }
        }else{
          // 有子节点
          let items = item.children;
          if(chooseType == 'parent' && item.indeterminate == false){
            if(item.checked){
              // 将数据添加到右侧
              this.inputTextData = this.inputTextData + prefix + ";\n";
            }
          }else{
            for(var i=0; i<items.length; i++){
              let item = items[i];
              this.appendDataWithPrefix(prefix + "." + item.title, item, chooseType);
            }
          }
        }
      },
      appendData:function (chooseType) {
        if(this.current_tab == "tab_output"){
          this.appendDataWithOutput(chooseType)
        }else{
          this.appendDataWithFuncs();
        }
      },
      appendDataWithFuncs: function(){
        let items = this.$refs.tree2.getCheckedAndIndeterminateNodes();
        if(items.length > 0){
          var item = items[0];
          this.inputTextData = this.inputTextData + item.title + ";\n";
          if(items.length > 1){
            this.$Message.warning("只有第一个有效！");
          }
        }
      },
      appendDataWithOutput:function (chooseType) {

        let items = this.$refs.tree1.getCheckedAndIndeterminateNodes();
        for(var i=0; i<items.length; i++){
          let item = items[i];
          // 只统计以 $ 开头的数据
          if(item.title.indexOf("$") != -1){
            this.appendDataWithPrefix(item.title,item, chooseType);
          }
        }
      },
      chooseTemplate:function (template) {
        this.inputTextData = this.inputTextData + template.template_value;
        this.$refs.templateModal.hideModal();
      },
      currentTabChanged:function (name) {
        this.current_tab = name;
      }
    },
    computed:{
      data1:function () {
        var appendChildrens = function (paramOutputSchemaTreeNode, node) {       // 父级节点对象、父级节点树元素
          if(paramOutputSchemaTreeNode.NodeChildrens != null && paramOutputSchemaTreeNode.NodeChildrens.length > 0){
            const arr = [];
            for(var i=0; i<paramOutputSchemaTreeNode.NodeChildrens.length; i++) {
              var childParamOutputSchemaTreeNode = paramOutputSchemaTreeNode.NodeChildrens[i];
              var childNode = {title: childParamOutputSchemaTreeNode.NodeName,expand: false};
              // 递归操作
              appendChildrens(childParamOutputSchemaTreeNode, childNode);
              arr.push(childNode);
            }
            node.children = arr;
          }
        };
        // tree 对应的 arr
        let treeArr = [];
        for(var i=0; i<this.prePosTreeNodeArr.length; i++){
          let preParamOutputSchemaTreeNode = this.prePosTreeNodeArr[i];
          const topTreeNode = {
            title: preParamOutputSchemaTreeNode.NodeName,
            expand: false,
          };
          appendChildrens(preParamOutputSchemaTreeNode,topTreeNode);
          treeArr.push(topTreeNode);
        }
        return treeArr;
      },
      data2:function () {
        // tree 对应的 arr
        let treeArr = [];
        for(var i=0; i<this.funcs.length; i++){
          treeArr.push({title:this.funcs[i].funcDemo});
        }
        return treeArr;
      },
    },
    watch: {
      inputTextData(val) {
        // 所有占位符变量
        this.variables = getMatchArrForString(this.inputTextData, /\$[a-zA-Z0-9]+/g);
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
