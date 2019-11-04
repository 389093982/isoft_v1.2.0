<template>
  <div style="margin:15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;min-height: 500px;">
    <!-- 内外边距：上右下左 -->
    <Row style="padding: 15px 10px 10px 25px;">
      <Col span="2">
        <IBeautifulLink @onclick="$router.push({path:'/igood/good_list'})">热销商品</IBeautifulLink>
      </Col>
      <Col span="2">
        <IBeautifulLink @onclick="$router.push({path:'/igood/mine/good_list',query:{type:'mine'}})">我的店铺商品</IBeautifulLink>
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
                <p>商品名称：<span><IBeautifulLink>{{good.good_name}}</IBeautifulLink></span></p>
                <p>商品描述：<span><IBeautifulLink :in-line="false">{{good.good_desc}}</IBeautifulLink></span></p>
                <p>商品价格：<span style="color: red;font-weight: bold;">￥{{good.good_price}}</span></p>
                <p>卖家姓名：{{good.good_seller}}</p>
                <p>卖家联系方式：{{good.seller_contact}}</p>
              </div>
              <div>

                <Button v-if="editable(good)" @click="$router.push({path:'/igood/mine/good_edit',query:{id:good.id}})">编辑商品</Button>
                <span v-else>
                  <Button @click="payConfirm(good)">立即购买</Button>
                </span>
              </div>
            </Col>
          </Row>
        </Col>
      </Row>
    </div>

    <div style="text-align: right;margin: 20px 100px 50px 0;">
      <Button @click="$router.push({path:'/igood/mine/good_edit'})">发布商品</Button>
    </div>
  </div>
</template>

<script>
  import IBeautifulLink from "../Common/link/IBeautifulLink";
  import {GoodList,NewOrder} from "../../api"
  import {CheckHasLogin,GetLoginUserName} from "../../tools"

  export default {
    name: "GoodList",
    components: {IBeautifulLink},
    data(){
      return {
        showGoodEditModal:false,
        goods:[],
      }
    },
    methods:{
      payConfirm:async function(good){
        const result = await NewOrder(good.id);
        if(result.status=="SUCCESS"){
          this.$router.push({path:'/igood/pay_confirm', query:{"good_id":good.id, "orderCode":result.orderCode}});
        }
      },
      refreshGoodList:async function () {
        const result = await GoodList();
        if(result.status == "SUCCESS"){
          this.goods = result.goods;
        }
      },
      editable:function (good) {
        return this.$route.query.type == 'mine' && CheckHasLogin() && GetLoginUserName() == good.good_seller;
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
    },
    watch: {
      '$route':'refreshGoodList'
    },
  }
</script>

<style scoped>

</style>
