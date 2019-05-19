<template>
  <div style="margin: 10px;">
    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <span slot="left">
        <Button type="success" @click="addEntity">新增</Button>
        <ISimpleConfirmModal ref="entityEditModal" modal-title="新增/编辑 Entity" :modal-width="600" :footer-hide="true">
          <IKeyValueForm ref="entityEditForm" form-key-label="entity_name" form-value-label="entity_type"
                         form-key-placeholder="请输入 entity_name" form-value-placeholder="请输入 entity_type"
                         @handleSubmit="editEntity" :formkey-validator="entityNameValidator"/>
        </ISimpleConfirmModal>
      </span>

      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="entities" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterPageEntity} from "../../../api"
  import {EditEntity} from "../../../api"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import {validateCommonPatternForString} from "../../../tools/index"
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import IKeyValueForm from "../../Common/form/IKeyValueForm"

  export default {
    name: "EntityList",
    components:{ISimpleSearch,ISimpleLeftRightRow,ISimpleConfirmModal,IKeyValueForm},
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
        entities: [],
        columns1: [
          {
            title: 'entity_name',
            key: 'entity_name',
            width: 250,
          },
          {
            title: 'entity_type',
            key: 'entity_type',
            width: 250,
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
                      this.$refs.entityEditModal.showModal();
                      this.$refs.entityEditForm.initFormData(this.entities[params.index].id, this.entities[params.index].entity_name, this.entities[params.index].entity_type);

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
      entityNameValidator(rule, value, callback){
        if (value === '') {
          callback(new Error('字段值不能为空!'));
        } else if (!validateCommonPatternForString(value)) {
          callback(new Error('存在非法字符，只能包含字母，数字，下划线!'));
        } else {
          callback();
        }
      },
      addEntity:function(){
        this.$refs.entityEditModal.showModal();
      },
      editEntity:async function (entity_id, entity_name, entity_type) {
        const result = await EditEntity(entity_id, entity_name, entity_type);
        if(result.status == "SUCCESS"){
          this.$refs.entityEditForm.handleSubmitSuccess("提交成功!");
          this.$refs.entityEditModal.hideModal();
          this.refreshEntityList();
        }else{
          this.$refs.entityEditForm.handleSubmitError("提交失败!");
        }
      },
      refreshEntityList:async function () {
        const result = await FilterPageEntity(this.search, this.offset, this.current_page);
        if(result.status=="SUCCESS"){
          this.entities = result.entities;
          this.total = result.paginator.totalcount;
        }
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshEntityList();
      },
      handleChange(page){
        this.current_page = page;
        this.refreshEntityList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshEntityList();
      },
    },
    mounted: function () {
      this.refreshEntityList();
    },
  }
</script>

<style scoped>

</style>
