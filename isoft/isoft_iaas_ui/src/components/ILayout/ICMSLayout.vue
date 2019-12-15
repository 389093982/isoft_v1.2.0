<template>
  <div>
    <Layout :style="{minHeight: '550px'}">
      <Header>
        <Header/>
      </Header>
      <Layout>
        <Sider hide-trigger :style="{background: '#fff'}">
          <Menu active-name="1-1" theme="light" width="auto" :open-names="['1']">
            <Submenu name="1">
              <template slot="title">
                <Icon type="ios-navigate"></Icon>
                分类管理
              </template>
              <MenuItem name="1-1"><router-link to="/background/cms/configuration">查看配置项</router-link></MenuItem>
            </Submenu>
            <Submenu name="2">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                广告位管理
              </template>
            </Submenu>
            <Submenu name="3">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                评论/建议管理
              </template>
              <MenuItem name="3-1"><router-link to="#">评论列表</router-link></MenuItem>
              <MenuItem name="3-1"><router-link to="/background/advise_list">意见建议列表</router-link></MenuItem>
            </Submenu>
            <Submenu name="4">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                流量计费统计
              </template>
              <MenuItem name="4-1"><router-link to="#">外部链接跳转上报信息统计</router-link></MenuItem>
            </Submenu>
            <Submenu name="5">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                日志管理
              </template>
              <MenuItem name="5-1"><router-link to="/background/sso/loginRecord">登录记录</router-link></MenuItem>
              <MenuItem name="5-2">用户查询</MenuItem>
              <MenuItem name="5-3">权限管理</MenuItem>
              <MenuItem name="5-4"><router-link to="/background/sso/appRegist">注册管理</router-link></MenuItem>
            </Submenu>
          </Menu>
        </Sider>
        <Content :style="{padding: '18px', minHeight: '550px', background: '#f5f6f7', lineHeight: '24px', fontSize: '14px'}">
          <router-view/>
        </Content>
      </Layout>
      <Footer>
        <Footer/>
      </Footer>
    </Layout>
  </div>
</template>

<script>
  import Header from '../Header/Header'

  export default {
    name: "ICMSLayout",
    components:{Header},
  }
</script>

<style scoped>

</style>
