<template>
  <div>
    <Row style="margin-bottom: 10px;">
      <Col span="10">
        <Button type="success" size="small" @click="editMigrate(null)" style="margin-bottom: 6px">新建迁移</Button>
      </Col>
      <Col span="14">
        <Select v-model="currentResourceName" style="width:300px">
          <Option v-for="resource in resources" :value="resource.resource_name">
            {{ resource.resource_name }} - {{ resource.resource_dsn }}
          </Option>
        </Select>
        <Button type="success" size="small" @click="executeMigrate(false)" style="margin-bottom: 6px">执行迁移</Button>
        <Button type="success" size="small" @click="executeMigrate(true)" style="margin-bottom: 6px">清理DB并执行迁移</Button>
      </Col>
    </Row>
    <p style="color: red;margin-bottom: 10px;margin-top: 10px;">{{errorMsg}}</p>

    <Table border :columns="columns1" :data="migrates" size="small"></Table>
    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterPageSqlMigrate} from "../../../api"
  import {ExecuteMigrate} from "../../../api"

  export default {
    name: "MigrateList",
    data(){
      return {
        errorMsg:'',
        currentResourceName:'',
        resources:[],
        migrates:[],
        // 当前页
        current_page:1,
        // 总页数
        total:1,
        // 每页记录数
        offset:10,
        columns1: [
          {
            title: 'id',
            key: 'id',
            width: 100,
          },
          {
            title: 'migrate_name',
            key: 'migrate_name',
          },
          {
            title: 'migrate_hash',
            key: 'migrate_hash',
            width: 350,
          },
          {
            title: '操作',
            key: 'operate',
            width: 220,
            fixed: 'right',
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
                      this.editMigrate(this.migrates[params.index]['id']);
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
      refreshMigrateList: async function(){
        const result = await FilterPageSqlMigrate(this.offset, this.current_page);
        this.migrates = result.migrates;
        this.resources = result.resources;
        this.total = result.paginator.totalcount;
      },
      handleChange(page){
        this.current_page = page;
        this.refreshMigrateList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshMigrateList();
      },
      editMigrate:function (id) {
        if(id != undefined && id != null){
          this.$router.push({ path: '/iwork/editMigrate', query: {id: id}});
        }else{
          this.$router.push({ path: '/iwork/editMigrate'});
        }
      },
      executeMigrate: async function (forceClean) {
        this.errorMsg = null;
        const result = await ExecuteMigrate(this.currentResourceName, forceClean);
        if(result.status == "SUCCESS"){
          this.$Message.success("SUCCESS");
          this.refreshMigrateList();
        }else{
          this.errorMsg = result.errorMsg;
        }
        this.refreshMigrateList();
      },
    },
    mounted(){
      this.refreshMigrateList();
    }
  }
</script>

<style scoped>

</style>
