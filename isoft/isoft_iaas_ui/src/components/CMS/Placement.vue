<template>
  <div>
    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <span slot="left">
        <Button type="success" @click="addPlacement">新增</Button>
        <ISimpleConfirmModal ref="placementEditModal" modal-title="新增/编辑占位符" :modal-width="600" :footer-hide="true">
          <IKeyValueForm ref="placementEditForm" form-key-label="placement_name" form-value-label="placement_desc"
                         form-key-placeholder="请输入 placement_name" form-value-placeholder="请输入 placement_desc"
                         @handleSubmit="editPlacement"/>
        </ISimpleConfirmModal>
      </span>

      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="placements" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import ISimpleConfirmModal from "../Common/modal/ISimpleConfirmModal"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import IKeyValueForm from "../Common/form/IKeyValueForm"
  import ISimpleSearch from "../Common/search/ISimpleSearch"
  import {AddPlacement} from "../../api"
  import {FilterPlacement} from "../../api"

  export default {
    name: "Placement",
    components:{ISimpleLeftRightRow,ISimpleConfirmModal,IKeyValueForm,ISimpleSearch},
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
        placements: [],
        columns1: [
          {
            title: 'id',
            key: 'id',
            width: 100,
          },
          {
            title: 'placement_name',
            key: 'placement_name',
            width: 400,
          },
          {
            title: 'placement_desc',
            key: 'placement_desc',
            width: 500,
          },
        ],
      }
    },
    methods:{
      addPlacement:function () {
        this.$refs.placementEditModal.showModal();
      },
      editPlacement:async function (placement_id, placement_name, placement_desc) {
        const result = await AddPlacement(placement_name, placement_desc);
        if(result.status == "SUCCESS"){
          this.$refs.placementEditForm.handleSubmitSuccess("提交成功!");
          this.$refs.placementEditModal.hideModal();
          this.refreshPlacementList();
        }else{
          this.$refs.placementEditForm.handleSubmitError("提交失败!");
        }
      },
      refreshPlacementList:async function () {
        const result = await FilterPlacement(this.offset,this.current_page,this.search);
        if(result.status=="SUCCESS"){
          this.placements = result.placements;
          this.total = result.paginator.totalcount;
        }
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshPlacementList();
      },
      handleChange(page){
        this.current_page = page;
        this.refreshPlacementList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshPlacementList();
      },
    },
    mounted(){
      this.refreshPlacementList();
    }
  }
</script>

<style scoped>

</style>
