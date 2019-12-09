<template>
  <IBeautifulCard title="娱乐版块">
    <div slot="content" style="padding: 20px;">
      <Row :gutter="10">
        <Col span="12"><IndexCarousel/></Col>
        <Col span="12"><IndexCarousel/></Col>
      </Row>
      <Row :gutter="10">
        <Col span="12"><IndexCarousel/></Col>
        <Col span="12"><IndexCarousel/></Col>
      </Row>
    </div>
  </IBeautifulCard>
</template>

<script>
  import IndexCarousel from "./IndexCarousel"
  import IBeautifulCard from "../Common/card/IBeautifulCard"

  export default {
    name: "Fun",
    components:{IBeautifulCard,IndexCarousel}
  }
</script>

<style scoped>

</style>
