<template>
  <div style="margin: 10px;">
    <ISimpleLeftRightRow>
      <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>
    </ISimpleLeftRightRow>

    <Table :columns="columns1" :data="loginRecords" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {formatDate} from "../../tools"
  import {LoginRecordList} from "../../api"
  import ISimpleLeftRightRow from "../Common/layout/ISimpleLeftRightRow"
  import ISimpleSearch from "../Common/search/ISimpleSearch"

  export default {
    name: "LoginRecord",
    components: {ISimpleLeftRightRow,ISimpleSearch},
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
        loginRecords: [],
        columns1: [
          {
            title: 'origin',
            key: 'origin',
            width:180
          },
          {
            title: 'referer',
            key: 'referer',
            width:250
          },
          {
            title: '登录ip',
            key: 'login_ip',
            width:100
          },
          {
            title: '登录用户',
            key: 'user_name',
            width:100
          },
          {
            title: '角色名称',
            key: 'role_name',
            width:100
          },
          {
            title: '登录状态',
            key: 'login_status',
            width:100
          },
          {
            title: '登录时间',
            key: 'created_time',
            render: (h,params)=>{
              return h('div',
                formatDate(new Date(params.row.created_time),'yyyy-MM-dd hh:mm')
              )
            }
          },
        ],
      }
    },
    methods:{
      refreshLoginRecordList: async function(){
        const result = await LoginRecordList(this.offset, this.current_page, this.search);
        if(result.status=="SUCCESS"){
          this.loginRecords = result.loginRecords;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshLoginRecordList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshLoginRecordList();
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshLoginRecordList();
      }
    },
    mounted: function () {
      this.refreshLoginRecordList();
    },
  }
</script>

<style scoped>

</style>
