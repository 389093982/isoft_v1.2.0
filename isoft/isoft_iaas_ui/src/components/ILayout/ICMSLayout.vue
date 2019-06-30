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
              系统支持
            </MenuItem>
          </div>
        </Menu>
      </Header>
      <Layout>
        <Sider hide-trigger :style="{background: '#fff'}">
          <Menu active-name="1-2" theme="light" width="auto" :open-names="['1']">
            <Submenu name="1">
              <template slot="title">
                <Icon type="ios-navigate"></Icon>
                分类管理
              </template>
              <MenuItem name="1-1"><router-link to="/cms/carousel_list">轮播图管理</router-link></MenuItem>
              <MenuItem name="1-2"><router-link to="/cms/catalog_list">列表管理</router-link></MenuItem>
            </Submenu>
            <Submenu name="2">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                广告位管理
              </template>
              <MenuItem name="2-1"><router-link to="/iwork/runLogList">日志列表</router-link></MenuItem>
              <MenuItem name="2-1"><router-link to="/iwork/workHistoryList">编辑历史</router-link></MenuItem>
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
  export default {
    name: "ICMSLayout"
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
    width: 520px;
    margin: 0 auto;
    margin-right: 20px;
  }
</style>
