<template>
  <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;min-height: 500px;">
    <!-- 内外边距：上右下左 -->
    <Row style="padding: 15px 10px 10px 25px;">
      <Col span="2">
        <IBeautifulLink2 @onclick="$router.push({path:'/igood/good_list'})">热销商品</IBeautifulLink2>
      </Col>
      <Col span="2">
        <IBeautifulLink2 @onclick="$router.push({path:'/igood/good_list'})">我的店铺商品</IBeautifulLink2>
      </Col>
    </Row>

    <div style="text-align: right;margin: 20px 100px 50px 0;">
      <Button @click="$router.push({path:'/igood/good_edit'})">发布商品</Button>
    </div>
  </div>
</template>

<script>
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2";
  export default {
    name: "GoodList",
    components: {IBeautifulLink2},
    data(){
      return {
        showGoodEditModal:false,
      }
    }
  }
</script>

<style scoped>

</style>
