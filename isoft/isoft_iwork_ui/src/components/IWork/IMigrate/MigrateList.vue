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

    <Tabs :value="tabVal">
      <TabPane label="列表" name="lst">
        <Table border :columns="columns1" :data="migrates" size="small"></Table>
      </TabPane>
      <TabPane label="预览" name="views">
        <Scroll height="450">
          <div v-for="migrate in migrates" style="border-bottom: 1px solid green;padding: 10px;">
            {{migrate.id}} ~~ <span v-html="renderBr(migrate.migrate_sql)"></span>
          </div>
        </Scroll>
      </TabPane>
      <TabPane label="执行日志" name="log">
        <Scroll height="450">
          <p v-for="log in logs" :style="{color: log.status ? 'grey' : 'red'}">{{log.tracking_detail}}</p>
        </Scroll>
      </TabPane>
    </Tabs>

    <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
          @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
  </div>
</template>

<script>
  import {FilterPageSqlMigrate,ExecuteMigrate,GetLastMigrateLogs,ToggleSqlMigrateEffective} from "../../../api"

  export default {
    name: "MigrateList",
    data(){
      return {
        tabVal:'lst',
        logs:[],
        timer:'',
        currentResourceName:'',
        resources:[],
        migrates:[],
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        columns1: [
          {
            title: 'id',
            key: 'id',
            width: 50,
          },
          {
            title: 'migrate_name',
            key: 'migrate_name',
          },
          {
            title: 'effective',
            key: 'effective',
            width: 100,
            render: (h, params) => {
              return h('span', {
                style:{
                  color: this.migrates[params.index]['effective'] ? "blue" : "grey",
                }
              }, this.migrates[params.index]['effective'] ? "生效" : "失效");
            }
          },
          {
            title: 'migrate_hash',
            key: 'migrate_hash',
            width: 350,
          },
          {
            title: '操作',
            key: 'operate',
            width: 180,
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
                h('Button', {
                  props: {
                    type: 'info',
                    size: 'small'
                  },
                  style: {
                    marginRight: '5px',
                  },
                  on: {
                    click: () => {
                      this.toggleMigrateEffective(this.migrates[params.index]['id']);
                    }
                  }
                }, '生效/失效'),
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
        var routeData;
        if(id != undefined && id != null){
          routeData = this.$router.resolve({ path: '/iwork/editMigrate', query: {id: id}});
        }else{
          routeData = this.$router.resolve({ path: '/iwork/editMigrate'});
        }
        // 打开新窗口跳转
        window.open(routeData.href, '_blank');
      },
      refreshMigrateLogs:async function(trackingId){
        const result = await GetLastMigrateLogs(trackingId);
        if(result.status == "SUCCESS"){
          this.logs = result.logs;
          if(result.over == true){
            clearInterval(this.timer);
          }
          this.refreshMigrateList();
        }
      },
      executeMigrate: async function (forceClean) {
        this.tabVal = 'log';
        const result = await ExecuteMigrate(this.currentResourceName, forceClean);
        if(result.status == "SUCCESS"){
          var _this = this;
          _this.timer = setInterval(function () {
            _this.refreshMigrateLogs(result.trackingId);
          }, 1000);
        }
      },
      toggleMigrateEffective:async function (id) {
        const result = await ToggleSqlMigrateEffective(id);
        if(result.status == "SUCCESS"){
          this.$Message.success("操作成功!");
          this.refreshMigrateList();
        }else{
          this.$Message.error("操作失败!");
        }
      },
      renderBr:function (str) {
        return str.replace(/\n/g,"<br>");
      }
    },
    mounted(){
      this.refreshMigrateList();
    },
  }
</script>

<style scoped>

</style>
