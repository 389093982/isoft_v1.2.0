<template>
  <div>

    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <IKeyValueForm ref="globalVarForm" slot="left" form-key-label="GlobalVarName" form-value-label="GlobalVarValue"
                     form-key-placeholder="请输入 GlobalVarName" form-value-placeholder="请输入 GlobalVarValue"
                     @handleSubmit="editGlobalVar"/>

      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="globalVars" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {GlobalVarList} from "../../../api"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import IKeyValueForm from "../../Common/form/IKeyValueForm"
  import {EditGlobalVar} from "../../../api"

  export default {
    name: "GlobalVarList",
    components:{ISimpleLeftRightRow,ISimpleSearch,IKeyValueForm},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        search:"",
        globalVars: [],
        columns1: [
          {
            title: 'name',
            key: 'name',
            width: 350,
          },
          {
            title: 'value',
            key: 'value',
          },
          {
            title: '操作',
            key: 'operate',
            render: (h, params) => {
              return h('div', [
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
                      this.$refs.globalVarForm.initFormData(this.globalVars[params.index].id, this.globalVars[params.index].name, this.globalVars[params.index].value);
                    }
                  }
                }, '编辑'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      editGlobalVar:async function(id, globalVarName, globalVarValue){
        const result = await EditGlobalVar(id, globalVarName, globalVarValue);
        if(result.status == "SUCCESS"){
          this.$refs.globalVarForm.handleSubmitSuccess("提交成功!");
          this.refreshGlobalVarList();
        }else{
          this.$refs.globalVarForm.handleSubmitError("提交失败!");
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshGlobalVarList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshGlobalVarList();
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshGlobalVarList();
      },
      refreshGlobalVarList:async function () {
        const result = await GlobalVarList(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.globalVars = result.globalVars;
          this.total = result.paginator.totalcount;
        }
      }
    },
    mounted: function () {
      this.refreshGlobalVarList();
    },
  }
</script>

<style scoped>

</style>
