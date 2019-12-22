<template>
  <div>
    <Layout :style="{minHeight: '550px'}">
      <Header>
        <Header/>
      </Header>
      <Content :style="{margin: '0px 100px 18px 100px', paddingTop: '5px', minHeight: '550px', background: '#f5f6f7', lineHeight: '24px', fontSize: '14px'}">
        <router-view/>
      </Content>
      <Footer>
        <Footer/>
      </Footer>
    </Layout>

    <RightSuspensionMenu/>
    <RandomAdmt/>
  </div>
</template>

<script>

  import Header from '../Header/Header'
  import Footer from '../Footer/Footer'
  import RightSuspensionMenu from "./RightSuspensionMenu"
  import RandomAdmt from "../Advertisement/RandomAdmt"

  export default {
    name: 'ILayout',
    components: {
      RandomAdmt,
      RightSuspensionMenu,
      Header,
      Footer
    },
  }
</script>

<style scoped>
  .bgArea{
  }
</style>
