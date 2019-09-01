<template>
  <div>
    <ISimpleLeftRightRow>
      <Button type="success" slot="left" @click="$router.push({ path: '/background/cms/element_edit'})">新增页面元素</Button>
      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Tag v-for="placement in placements" color="green">
      <span @click="clickPlacement(placement.placement_name)">{{placement.placement_name}}</span>
    </Tag>

    <Table border :columns="columns1" :data="elements" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterElements,UpdateElementStatus,GetAllPlacements} from "../../../api"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import MultiClickButton from "../../Common/button/MultiClickButton"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"

  export default {
    name: "Element",
    components:{ISimpleLeftRightRow,MultiClickButton,ISimpleSearch},
    data () {
      var _this = this;
      return {
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        // 搜索条件
        search:"",
        placements:[],
        elements: [],
        columns1: [
          {
            title: 'placement',
            key: 'placement',
            width:250
          },
          {
            title: '导航级别',
            key: 'navigation_level',
            width:100
          },
          {
            title: '父 id',
            key: 'navigation_parent_id',
            width:100
          },
          {
            title: 'title',
            key: 'title',
            width:150
          },
          {
            title: 'status',
            key: 'status',
            width:80,
            render: (h,params)=> {
              return h('div', {
                style:{
                  color: this.elements[params.index].status == 1 ?  'green' : (this.elements[params.index].status == 0 ? 'red' : 'grey'),
                }
              },
              this.elements[params.index].status == 1 ?  '启用' : (this.elements[params.index].status == 0 ? '停用' : '失效'))
            }
          },
          {
            title: 'img_path',
            key: 'img_path',
            width:200
          },
          {
            title: 'content',
            key: 'content',
            width:200
          },
          {
            title: 'linked_refer',
            key: 'linked_refer',
            width:200
          },
          {
            title: '操作',
            key: 'operate',
            width:280,
            fixed: 'right',
            render: (h,params)=> {
              return h('div',[
                h(MultiClickButton,{
                  props:{
                    btnCounts: 5,
                    btnTypes: ['primary','info','warning',"error", 'success'],
                    btnShows: [true, true, true, true, true],
                    btnBindDatas: [1, 0, -1, 2, 3],
                    btnTexts: ['启用', '停用', '失效', '删除', '编辑'],
                  },
                  on:{
                    handleClick:async function (index, bindData) {
                      if (bindData == 3){   // 编辑模式
                        _this.$router.push({ path: '/background/cms/element_edit', query: { id: _this.elements[params.index].id }});
                      }else{
                        const result = await UpdateElementStatus(_this.elements[params.index].id, bindData);
                        if(result.status == "SUCCESS"){
                          _this.refreshElementList();
                        }else{
                          _this.$Message.error("状态更新失败!");
                        }
                      }
                    }
                  }
                })
              ]);
            }
          },
        ],
      }
    },
    methods: {
      clickPlacement:function(placement_name){
        this.search = placement_name;
        this.refreshElementList();
      },
      refreshElementList:async function () {
        const result = await FilterElements(this.offset, this.current_page, this.search);
        if(result.status=="SUCCESS"){
          this.elements = result.elements;
          this.total = result.paginator.totalcount;
        }
      },
      refreshAllPlacements:async function () {
        const result = await GetAllPlacements();
        if(result.status == "SUCCESS"){
          this.placements = result.placements;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshElementList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshElementList();
      },
      handleSearch(data){
        this.search = data;
        this.refreshElementList();
      },
    },
    mounted(){
      if(this.$route.query.search != undefined){
        this.search = this.$route.query.search;
      }
      this.refreshElementList();
      this.refreshAllPlacements();
    }
  }
</script>

<style scoped>

</style>
