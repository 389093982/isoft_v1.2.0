<template>
  <div class="layout">
    <Layout>
      <Header>
        <Menu mode="horizontal" theme="dark" active-name="1">
          <div class="layout-logo"></div>
          <div class="layout-nav">
            <MenuItem name="1">
              <Icon type="ios-navigate"></Icon>
              使用说明
            </MenuItem>
            <MenuItem name="2">
              <Icon type="ios-keypad"></Icon>
              用户协议
            </MenuItem>
            <MenuItem name="3">
              <Icon type="ios-analytics"></Icon>
              License管理
            </MenuItem>
            <MenuItem name="4">
              <Icon type="ios-paper"></Icon>
              <span @click="saveProject">保存项目</span>
            </MenuItem>
            <MenuItem name="4">
              <Icon type="ios-paper"></Icon>
              <span @click="importProject">导入项目</span>
            </MenuItem>
          </div>
        </Menu>
      </Header>
      <Layout>
        <Sider hide-trigger :style="{background: '#fff'}">
          <Menu active-name="1-2" theme="light" width="auto" :open-names="['1']" accordion>
            <Submenu name="1">
              <template slot="title">
                <Icon type="ios-navigate"></Icon>
                流程管理
              </template>
              <MenuItem name="1-1"><router-link to="/iwork/moduleList">模块管理</router-link></MenuItem>
              <MenuItem name="1-2"><router-link to="/iwork/workList">流程列表</router-link></MenuItem>
              <MenuItem name="1-3"><router-link to="/iwork/filterList">过滤器管理</router-link></MenuItem>
              <MenuItem name="1-4"><router-link to="/iwork/globalVarList">全局变量管理</router-link></MenuItem>
              <MenuItem name="1-5">WorkDL管理</MenuItem>
            </Submenu>
            <Submenu name="2">
              <template slot="title">
                <Icon type="ios-keypad"></Icon>
                资源管理
              </template>
              <MenuItem name="2-1"><router-link to="/iwork/resourceList">资源列表</router-link></MenuItem>
              <MenuItem name="2-2"><router-link to="/iwork/migrateList">数据库迁移管理</router-link></MenuItem>
            </Submenu>
            <Submenu name="3">
              <template slot="title">
                <Icon type="ios-analytics"></Icon>
                定时任务
              </template>
              <MenuItem name="3-1"><router-link to="/iwork/quartzList">定时任务列表</router-link></MenuItem>
            </Submenu>
            <Submenu name="4">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                日志管理
              </template>
              <MenuItem name="4-1"><router-link to="/iwork/runLogList">日志列表</router-link></MenuItem>
              <MenuItem name="4-1"><router-link to="/iwork/workHistoryList">编辑历史</router-link></MenuItem>
              <MenuItem name="4-1"><router-link to="/iwork/dashboard">统计仪表盘</router-link></MenuItem>
            </Submenu>
            <Submenu name="5">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                帮助助手
              </template>
              <MenuItem name="5-1"><router-link to="/iwork/quickSql">快捷sql</router-link></MenuItem>
              <MenuItem name="5-2"><router-link to="/iwork/files">文件服务器管理</router-link></MenuItem>
              <MenuItem name="5-3"><router-link to="/iwork/audit">内容审核系统</router-link></MenuItem>
            </Submenu>
            <Submenu name="6">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                内容管理
              </template>
              <MenuItem name="6-1"><router-link to="/iwork/placementList">占位符管理</router-link></MenuItem>
            </Submenu>
          </Menu>
        </Sider>
        <Layout :style="{padding: '24px'}">
          <Content :style="{padding: '24px', minHeight: '480px', background: '#fff'}">
            <router-view/>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  </div>
</template>

<script>
  import {SaveProject,ImportProject} from "../../api"

  export default {
    name: "IWorkLayout",
    data(){
      return {
      }
    },
    methods:{
      saveProject:async function () {
        const result = await SaveProject();
        if(result.status == "SUCCESS"){
          this.$Message.success("保存成功!");
        }
      },
      importProject:async function () {
        const result = await ImportProject();
        if(result.status == "SUCCESS"){
          this.$Message.success("导入成功!");
        }
      }
    }
  }
</script>

<style scoped>
  .layout{
    border: 1px solid #d7dde4;
    background: #f5f7f9;
    position: relative;
    border-radius: 4px;
    overflow: hidden;
  }
  .layout-logo{
    width: 100px;
    height: 30px;
    background: #5b6270;
    border-radius: 3px;
    float: left;
    position: relative;
    top: 15px;
    left: 20px;
  }
  .layout-nav{
    width: 620px;
    margin: 0 auto;
    margin-right: 20px;
  }
</style>
