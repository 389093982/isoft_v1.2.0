<template>
  <div style="margin:15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;min-height: 500px;">
    <!-- 内外边距：上右下左 -->
    <Row style="padding: 15px 10px 10px 25px;">
      <Col span="2">
        <IBeautifulLink2 @onclick="$router.push({path:'/igood/good_list'})">热销商品</IBeautifulLink2>
      </Col>
      <Col span="2">
        <IBeautifulLink2 @onclick="$router.push({path:'/igood/good_list'})">我的店铺商品</IBeautifulLink2>
      </Col>
    </Row>

    <div style="margin: 20px;">
      <Row :gutter="10">
        <Col span="12" v-for="good in goods">
          <Row style="margin-bottom: 10px;">
            <Col span="12" style="text-align: center;">
              <router-link :to="{path:'/igood/good_detail',query:{id:good.id}}">
                <!-- 长度大于 2 排除空数组 [] -->
                <img v-if="good.good_images.length > 2" :src="good.good_images | filterFirst" width="180px" height="180px"/>
                <img v-else src="../../assets/default.png" width="180px" height="180px"/>
              </router-link>
            </Col>
            <Col span="12">
              <div style="height: 150px;width: 100%;word-wrap:break-word;word-break:break-all;overflow: hidden;">
                <p>商品名称：<span><IBeautifulLink2>{{good.good_name}}</IBeautifulLink2></span></p>
                <p>商品描述：<span><IBeautifulLink2 :in-line="false">{{good.good_desc}}</IBeautifulLink2></span></p>
                <p>商品价格：<span style="color: red;font-weight: bold;">￥{{good.good_price}}</span></p>
              </div>
              <div>
                <Button>加入购物车</Button>
                <Button>立即购买</Button>
              </div>
            </Col>
          </Row>
        </Col>
      </Row>
    </div>

    <div style="text-align: right;margin: 20px 100px 50px 0;">
      <Button @click="$router.push({path:'/igood/good_edit'})">发布商品</Button>
    </div>
  </div>
</template>

<script>
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2";
  import {GoodList} from "../../api"

  export default {
    name: "GoodList",
    components: {IBeautifulLink2},
    data(){
      return {
        showGoodEditModal:false,
        goods:[],
      }
    },
    methods:{
      refreshGoodList:async function () {
        const result = await GoodList();
        if(result.status == "SUCCESS"){
          this.goods = result.goods;
        }
      }
    },
    filters:{
      filterFirst:function (good_images) {
        let arr = JSON.parse(good_images);
        return arr[0];
      }
    },
    mounted(){
      this.refreshGoodList();
    }
  }
</script>

<style scoped>

</style>
