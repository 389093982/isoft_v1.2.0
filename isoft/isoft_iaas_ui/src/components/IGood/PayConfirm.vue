<template>
  <div>
    <div v-if="good">
      <Row>
        <Col span="12"><img :src="good_images[0]" width="100%" height="100%"/></Col>
        <Col span="12">{{good.good_desc}}</Col>
      </Row>
      <Row style="text-align: right;">
        <Button>付款</Button>
      </Row>
    </div>
  </div>
</template>

<script>
  import {GetGoodDetail} from "../../api"

  export default {
    name: "PayConfirm",
    data(){
      return {
        good:null,
        good_images:[],   // 商品图片
      }
    },
    methods:{
      refreshGoodDetail:async function () {
        const result = await GetGoodDetail(this.$route.query.good_id);
        if(result.status == "SUCCESS"){
          this.good = result.good;
          this.good_images = JSON.parse(result.good.good_images);
        }
      }
    },
    mounted(){
      this.refreshGoodDetail();
    }
  }
</script>

<style scoped>

</style>
