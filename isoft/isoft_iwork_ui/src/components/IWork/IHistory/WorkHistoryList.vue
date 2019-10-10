<template>
  <span>
    <ISimpleSearch slot="right" @handleSimpleSearch="handleSearch"/>

    <Table :columns="columns1" :data="workHistories" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
    <Modal
      v-model="modal1"
      title="流程详情"
      :footer-hide="true"
      :mask-closable="false"
      :styles="{top: '20px'}"
      width="1000">
      <Input type="textarea" :value="workHistory" :readonly="true" :rows="20"/>
    </Modal>
  </span>
</template>

<script>
  import {FilterPageWorkHistory,RestoreFromWorkHistory} from "../../../api/index"
  import {formatDate} from "../../../tools/index"
  import ISimpleSearch from "../../Common/search/ISimpleSearch";

  export default {
    name: "WorkHistoryList",
    components: {ISimpleSearch},
    data(){
      return {
        modal1: false,
        workHistory: "",
        // 当前页
        current_page:1,
        // 搜索条件
        search:"",
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        workHistories: [],
        columns1: [
          {
            title: 'id',
            key: 'id',
          },
          {
            title: 'hash',
            key: 'hash',
            width: 200,
          },
          {
            title: 'work_name',
            key: 'work_name',
            width: 300,
          },
          {
            title: 'last_updated_time',
            key: 'last_updated_time',
            render: (h,params)=>{
              return h('div',
                formatDate(new Date(params.row.last_updated_time),'yyyy-MM-dd hh:mm')
              )
            }
          },
          {
            title: '操作',
            key: 'operate',
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'success',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.modal1 = true;
                      this.workHistory = this.workHistories[params.index]['work_history'];
                    }
                  }
                }, '查看详情'),
                h('Button', {
                  props: {
                    type: 'warning',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.restoreFromWorkHistory(this.workHistories[params.index].id);
                    }
                  }
                }, '还原'),
              ]);
            }
          }
        ],
      }
    },
    methods:{
      restoreFromWorkHistory:async function(id){
        const result = await RestoreFromWorkHistory(id);
        if(result.status == "SUCCESS"){
          this.$Message.success("还原成功！");
        }else{
          this.$Message.error("还原失败！");
        }
      },
      handleSearch(data){
        this.offset = 10;
        this.current_page = 1;
        this.search = data;
        this.refreshWorkHistoryList();
      },
      refreshWorkHistoryList:async function(){
        const result = await FilterPageWorkHistory(this.offset,this.current_page, this.search);
        if(result.status=="SUCCESS"){
          this.workHistories = result.workHistories;
          this.total = result.paginator.totalcount;
        }
      },
      handleChange(page){
        this.current_page = page;
        this.refreshWorkHistoryList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshWorkHistoryList();
      },
    },
    mounted: function () {
      this.refreshWorkHistoryList();
    },
  }
</script>

<style scoped>

</style>
