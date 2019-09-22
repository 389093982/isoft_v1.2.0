<template>
  <div style="margin:15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;min-height: 500px;">
    <div v-if="good" style="padding: 50px;">
      <Row>
        <Col span="12">
          AAAAAAAAAAAAAAAA
        </Col>
        <Col  span="12">
          <p>商品描述：<span><IBeautifulLink2>{{good.good_desc}}</IBeautifulLink2></span></p>
          <p>商品价格：<span style="color: red;font-weight: bold;">￥{{good.good_price}}</span></p>
        </Col>
      </Row>

    </div>
  </div>
</template>

<script>
  import {GetGoodDetail} from "../../api"
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2";

  export default {
    name: "GoodDetail",
    components:{IBeautifulLink2},
    data(){
      return {
        good:null,
      }
    },
    methods:{
      refreshGoodDetail:async function (good_id) {
        const result = await GetGoodDetail(good_id);
        if(result.status == "SUCCESS"){
          this.good = result.good;
        }
      }
    },
    mounted(){
      if(this.$route.query.id != undefined && this.$route.query.id != null){
        this.refreshGoodDetail(this.$route.query.id);
      }
    }
  }
</script>

<style scoped>

</style>
