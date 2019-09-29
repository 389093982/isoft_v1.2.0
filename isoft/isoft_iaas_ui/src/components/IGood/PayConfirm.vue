<template>
  <div>
    <div v-if="good">
      <Row>
        <Col span="12"><img :src="good_images[0]" width="100%" height="400px"/></Col>
        <Col span="12">{{good.good_desc}}</Col>
      </Row>
      <Row style="text-align: right;" v-if="orderInfo">
        <Button v-if="orderInfo.payment_status == 1">付款</Button>
      </Row>
    </div>
  </div>
</template>

<script>
  import {GetGoodDetail,GetOrderDetail} from "../../api"

  export default {
    name: "PayConfirm",
    data(){
      return {
        good:null,
        good_images:[],   // 商品图片
        orderInfo:null,
      }
    },
    methods:{
      refreshGoodDetail:async function () {
        const result = await GetGoodDetail(this.$route.query.good_id);
        if(result.status == "SUCCESS"){
          this.good = result.good;
          this.good_images = JSON.parse(result.good.good_images);
        }
      },
      refreshOrderInfo:async function () {
        const result = await GetOrderDetail(this.$route.query.orderCode);
        if(result.status == "SUCCESS"){
          this.orderInfo = result.orderInfo;
        }
      }
    },
    mounted(){
      this.refreshGoodDetail();
      this.refreshOrderInfo();
    }
  }
</script>

<style scoped>

</style>
