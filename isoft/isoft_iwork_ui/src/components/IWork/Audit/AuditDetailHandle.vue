<template>
  <div>
    <div v-for="rowData in rowDatas" style="background-color: rgba(236,236,236,0.3);margin: 5px;padding: 10px;">
      <Row>
        <Col span="20">
          <span v-for="(colName, index) in colNames" style="margin-right: 10px;">
            <Tag color="orange">字段名：{{colName}}</Tag> {{rowData[colName]}}<br/>
          </span>
        </Col>
        <Col span="4">
          <span v-if="update_cases" v-for="update_case in update_cases">
            <Button v-if="update_case.case_name" type="success" size="small" style="margin-right: 5px;">{{update_case.case_name}}</Button>
          </span>
        </Col>
      </Row>
    </div>

    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {GetAuditHandleData,QueryTaskDetail} from "../../../api"
  export default {
    name: "AuditDetailHandle",
    data(){
      return {
        rowDatas:[],
        colNames:[],
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        search:"",
        update_cases:[],
      }
    },
    methods:{
      handleChange(page){
        this.current_page = page;
        this.refreshHandleData();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshHandleData();
      },
      refreshHandleData:async function () {
        const result = await GetAuditHandleData(this.$route.query.task_name, this.current_page, this.offset);
        if(result.status == "SUCCESS"){
          this.rowDatas = result.rowDatas;
          this.colNames = JSON.parse(result.colNames);
          this.total = result.totalcount;
        }
      },
      refreshAuditDetail:async function () {
        const result = await QueryTaskDetail(this.$route.query.task_name);
        if(result.status == "SUCCESS"){
          this.update_cases = result.taskDetail.update_cases;
        }
      }
    },
    mounted(){
      this.refreshHandleData();
      this.refreshAuditDetail();
    }
  }
</script>

<style scoped>

</style>
