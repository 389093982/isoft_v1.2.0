<template>
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
</template>

<script>
  import {joinArray} from "../../../../tools"

  export default {
    name: "ParamInputEditDataSource",
    props:{
      prePosTreeNodeArr:{
        type: Array,
        default: [],
      }
    },
    data(){
      return {
        current_tab:'tab_output',
        funcs:this.GLOBAL.quick_funcs,
      }
    },
    methods:{
      getChooseDatas:function (chooseType){
        if(this.current_tab == "tab_output"){
          return this.getChooseDatasWithOutput(chooseType)
        }else{
          return this.getChooseDatasWithFuncs();
        }
      },
      getChooseDatasWithOutput:function (chooseType) {
        let result = [];
        let items = this.$refs.tree1.getCheckedAndIndeterminateNodes();
        for(var i=0; i<items.length; i++){
          let item = items[i];
          // 只统计以 $ 开头的数据
          if(item.title.indexOf("$") != -1){
            let _result = this.getChooseDataWithPrefix(item.title,item, chooseType);
            result = joinArray(result, _result);
          }
        }
        return result;
      },
      getChooseDataWithPrefix:function(prefix, item, chooseType){
        let result = [];
        // 没有子节点
        if(item.children == null){
          if(item.indeterminate == false){
            if(item.checked){
              result.push(prefix + ";\n");
            }
          }
        }else{
          // 有子节点
          let items = item.children;
          if(chooseType == 'parent' && item.indeterminate == false){
            if(item.checked){
              result.push(prefix + ";\n");
            }
          }else{
            for(var i=0; i<items.length; i++){
              let item = items[i];
              let _result = this.getChooseDataWithPrefix(prefix + "." + item.title, item, chooseType);
              result = joinArray(result, _result);
            }
          }
        }
        return result;
      },
      getChooseDatasWithFuncs: function(){
        let items = this.$refs.tree2.getCheckedAndIndeterminateNodes();
        if(items.length > 0){
          if(items.length > 1){
            this.$Message.warning("只有第一个有效！");
          }
          var item = items[0];
          return item.title + ";\n";
        }
        return null;
      },
      currentTabChanged:function (name) {
        this.current_tab = name;
      },
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
  }
</script>

<style scoped>

</style>
