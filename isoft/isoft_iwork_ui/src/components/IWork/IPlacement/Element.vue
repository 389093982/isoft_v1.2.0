<template>
  <div>
    <ISimpleLeftRightRow>
      <span slot="left">
        <Button type="success" size="small" @click="$router.push({ path: '/iwork/elementEdit', query: {placement_name:$route.query.placement_name}})">新增页面元素</Button>
        <IFileUpload size="small" ref="fileUpload" @uploadComplete="uploadComplete" action="/api/iwork/import" uploadLabel="导入"/>
      </span>

      <!-- right 插槽部分 -->
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table border :columns="columns1" :data="elements" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterElements,UpdateElementStatus,CopyElement,ImportElement} from "../../../api"
  import ISimpleSearch from "../../Common/search/ISimpleSearch"
  import MultiClickButton from "../../Common/button/MultiClickButton"
  import ISimpleLeftRightRow from "../../Common/layout/ISimpleLeftRightRow"
  import IFileUpload from "../../Common/file/IFileUpload"

  export default {
    name: "Element",
    components:{ISimpleLeftRightRow,MultiClickButton,ISimpleSearch,IFileUpload},
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
        elements: [],
        columns1: [
          {
            title: 'id',
            key: 'id',
            width:100
          },
          {
            title: 'placement',
            key: 'placement',
            width:150
          },
          {
            title: 'element_name',
            key: 'element_name',
            width:150
          },
          {
            title: 'status',
            key: 'status',
            width:100,
            render: (h,params)=> {
              return h('div', {
                  style:{
                    color: this.elements[params.index].status == 1 ?  'green' : 'red',
                  }
                },
                this.elements[params.index].status == 1 ?  '启用' : '停用')
            },
            filters: [
              {
                label: '启用',
                value: 1,
              },
              {
                label: '停用',
                value: -1,
              },
            ],
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
            width:340,
            fixed: 'right',
            render: (h,params)=> {
              return h('div',[
                h(MultiClickButton,{
                  props:{
                    btnCounts: 5,
                    btnTypes: ['primary','info','error',"warning", 'success'],
                    btnShows: [true, true, true, true, true, true],
                    btnBindDatas: [1, -1, -2, 2, 3],
                    btnTexts: ['启用', '停用', '删除', '编辑', "复制"],
                  },
                  on:{
                    handleClick:async function (index, bindData) {
                      if (bindData == 2){   // 编辑模式
                        _this.$router.push({ path: '/iwork/elementEdit', query: {id:_this.elements[params.index].id, placement_name:_this.$route.query.placement_name}});
                      }else if (bindData == 3){   // 复制模式
                        _this.copyElement(_this.elements[params.index].id);
                      }else{
                        const result = await UpdateElementStatus(_this.elements[params.index].id, bindData);
                        if(result.status == "SUCCESS"){
                          _this.$Message.success("操作成功!");
                          _this.refreshElementList();
                        }else{
                          _this.$Message.error("操作失败!");
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
      copyElement:async function(id){
        const result = await CopyElement(id);
        if(result.status == "SUCCESS"){
          this.$Message.success("复制成功！");
          this.refreshElementList();
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
      clickPlacement:function(placement_name){
        this.search = placement_name;
        this.refreshElementList();
      },
      refreshElementList:async function () {
        const result = await FilterElements(this.offset, this.current_page, this.$route.query.placement_name, this.search);
        if(result.status=="SUCCESS"){
          this.elements = result.elements;
          this.total = result.paginator.totalcount;
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
      uploadComplete: function (result) {
        if(result.status == "SUCCESS"){
          this.refreshElementList();
          this.$Message.success("导入成功！");
        }else{
          this.$Message.error(result.errorMsg);
        }
      },
    },
    mounted(){
      this.refreshElementList();
    }
  }
</script>

<style scoped>

</style>
