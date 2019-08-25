<template>
  <div>
    <Layout :style="{minHeight: '550px'}">
      <Header>
        <Header/>
      </Header>
      <Layout>
        <Sider hide-trigger :style="{background: '#fff'}">
          <Menu active-name="1-2" theme="light" width="auto" :open-names="['1']">
            <Submenu name="1">
              <template slot="title">
                <Icon type="ios-navigate"></Icon>
                分类管理
              </template>
              <MenuItem name="1-1"><router-link to="/cms/placement_list">占位符管理</router-link></MenuItem>
              <MenuItem name="1-2"><router-link to="/cms/element_list">轮播图管理</router-link></MenuItem>
              <MenuItem name="1-3"><router-link to="/cms/catalog_list">列表管理</router-link></MenuItem>
              <MenuItem name="1-4"><router-link to="/cms/configuration">查看配置项</router-link></MenuItem>
              <MenuItem name="1-5"><router-link to="/cms/commonLinkList">友情链接</router-link></MenuItem>
            </Submenu>
            <Submenu name="2">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                广告位管理
              </template>
            </Submenu>
            <Submenu name="2">
              <template slot="title">
                <Icon type="ios-barcode"></Icon>
                评论管理
              </template>
              <MenuItem name="2-1"><router-link to="#">评论列表</router-link></MenuItem>
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
