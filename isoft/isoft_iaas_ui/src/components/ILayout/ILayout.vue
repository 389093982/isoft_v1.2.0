<template>
  <div>
    <Layout :style="{minHeight: '550px'}">
      <Header>
        <Header/>
      </Header>
      <Content :style="{padding: '18px', paddingTop: '5px', minHeight: '550px', background: '#f5f6f7', lineHeight: '24px', fontSize: '14px'}">
        <router-view/>
      </Content>
      <Footer>
        <Footer/>
      </Footer>
    </Layout>

    <LeftSuspensionMenu/>
  </div>
</template>

<script>

  import Header from '../Header/Header'
  import Footer from '../Footer/Footer'
  import LeftSuspensionMenu from "./RightSuspensionMenu";

  export default {
    name: 'ILayout',
    components: {
      LeftSuspensionMenu,
      Header,
      Footer
    },
  }
</script>

<style lang="stylus" type="text/stylus" rel="stylesheet/stylus">

</style>
