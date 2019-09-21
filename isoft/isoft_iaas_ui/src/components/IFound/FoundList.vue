<template>
  <div style="padding:20px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;">
    <div><IndexCarousel/></div>
    <div><IBeautifulTabLink/></div>
    <div>
      <Row :gutter="10">
        <Col span="12"><IndexCarousel/></Col>
        <Col span="12"><IndexCarousel/></Col>
      </Row>
    </div>
  </div>
</template>

<script>
  import IndexCarousel from "../ILearning/IndexCarousel";
  import IBeautifulTabLink from "../Common/link/IBeautifulTabLink";
  export default {
    name: "FoundList",
    components: {IBeautifulTabLink, IndexCarousel}
  }
</script>

<style scoped>

</style>
