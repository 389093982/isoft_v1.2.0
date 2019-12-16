<template>
  <div>
    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <Button type="success" size="small" slot="left" @click="$router.push({ path: '/iwork/placementEdit'})" v-if="!this.chooserMode">新增占位符</Button>

      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="placements" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import ISimpleConfirmModal from "../../Common/modal/ISimpleConfirmModal"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import IKeyValueForm from "../../Common/form/IKeyValueForm"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import {FilterPlacement,DeletePlacementById,CopyPlacement} from "../../../api"

  export default {
    name: "Placement",
    components:{ISimpleLeftRightRow,ISimpleConfirmModal,IKeyValueForm,ISimpleSearch},
    props:{
      chooserMode:{ // 选择模式
        type: Boolean,
        default: false,
      }
    },
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
            width: 300,
          },
          {
            title: 'placement_label',
            key: 'placement_label',
            width: 300,
          },
          {
            title: 'placement_desc',
            key: 'placement_desc',
            width: 400,
          },
          {
            title: '操作',
            key: 'operate',
            width: 250,
            fixed: 'right',
            render: (h,params)=> {
              return h('div',[
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: this.chooserMode ? 'undefined': 'none',      // 选择模式显示
                  },
                  on: {
                    click: () => {
                      this.$emit("choosePlacement", this.placements[params.index].placement_name);
                    }
                  }
                }, '选择'),
                h('Button', {
                  props: {
                    type: 'error',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: !this.chooserMode ? 'undefined': 'none',   // 非选择模式显示
                  },
                  on: {
                    click: () => {
                      this.deletePlacementById(this.placements[params.index].id);
                    }
                  }
                }, '删除'),
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: !this.chooserMode ? 'undefined': 'none',   // 非选择模式显示
                  },
                  on: {
                    click: () => {
                      this.$router.push({ path: '/iwork/placementEdit', query: { id: this.placements[params.index].id }});
                    }
                  }
                }, '编辑'),
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: !this.chooserMode ? 'undefined': 'none',   // 非选择模式显示
                  },
                  on: {
                    click: () => {
                      this.copyPlacement(this.placements[params.index].id);
                    }
                  }
                }, '复制'),
                h('Button', {
                  props: {
                    type: 'warning',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                    display: !this.chooserMode ? 'undefined': 'none',   // 非选择模式显示
                  },
                  on: {
                    click: () => {
                      this.$router.push({path:'/iwork/elementList',query:{ placement_name: this.placements[params.index].placement_name }});
                    }
                  }
                }, '元素管理'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      copyPlacement:async function(id){
        const result = await CopyPlacement(id);
        if(result.status == "SUCCESS"){
          this.$Message.success("复制成功！");
          this.refreshPlacementList();
        }else{
          this.$Message.error(result.errorMsg);
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
      deletePlacementById: function(id){
        var _this = this;
        _this.$Modal.confirm({
          title: '删除',
          content: '确认删除该条数据吗？请谨慎操作！',
          onOk: async () => {
            const result = await DeletePlacementById(id);
            if(result.status == "SUCCESS"){
              _this.$Message.success("删除成功！");
              _this.refreshPlacementList();
            }else{
              _this.$Message.error(result.errorMsg);
            }
          },
          onCancel: () => {
            _this.$Message.info('取消操作');
          }
        });
      }
    },
    mounted(){
      this.refreshPlacementList();
    }
  }
</script>

<style scoped>

</style>
