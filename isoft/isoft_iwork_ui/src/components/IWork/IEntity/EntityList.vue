<template>
  <div style="margin: 10px;">
    <ISimpleLeftRightRow style="margin-bottom: 10px;margin-right: 10px;">
      <!-- left 插槽部分 -->
      <span slot="left">
        <Button type="success">新增</Button>
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
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"

  export default {
    name: "EntityList",
    components:{ISimpleSearch,ISimpleLeftRightRow},
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
