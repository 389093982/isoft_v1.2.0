<template>
  <div>
    <div style="margin: 5px 100px 0 100px;float: right;" class="clear">
      <ul>
        <li><span @click="$router.push({path:'/ilearning/advise'})">我要吐槽/建议</span></li>
        <li>
          网站导航 <Icon type="ios-arrow-down" />
          <div class="subArea">
            BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB

            <div class="linearTransitionBg"></div>

            <div>
              <span class="linearTransitionBox linearTransitionBorderBottom" @click="$router.push({path:'/igood/good_list'})">
                我要赚钱
              </span>
              <span class="linearTransitionBox linearTransitionBorderBottom" @click="$router.push({path:'/ifound/discount_list'})">
                聚优惠
              </span>
              <span class="linearTransitionBox linearTransitionBorderBottom" @click="$router.push({path:'/ifound/activity_list'})">
                聚活动
              </span>
              <span class="linearTransitionBox linearTransitionBorderBottom">AAAAAAAA</span>
              <span class="linearTransitionBox linearTransitionBorderBottom">AAAAAAAA</span>
              <span class="linearTransitionBox linearTransitionBorderBottom">AAAAAAAA</span>
              <span class="linearTransitionBox linearTransitionBorderBottom">AAAAAAAA</span>
            </div>
          </div>
        </li>
        <li @mouseenter="handleMouseenter"
            @mouseleave="handleMouseleave">
          热点 APP
          <Icon type="ios-arrow-up" v-if="up"/>
          <Icon type="ios-arrow-down" v-else/>
          <div class="subArea">
            BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    name: "LevelTwoHeader",
    data(){
      return {
        up:false,
      }
    },
    methods:{
      handleMouseenter () {
        this.up = true;
      },
      handleMouseleave () {
        this.up = false;
      },
    }
  }
</script>

<style scoped>
  @import "../../assets/css/isoft_common.css";

  li{
    float: left;
    list-style: none;
    margin-left: 10px;
    padding: 2px 5px;
    display: inline-block;
    cursor: pointer;
    background-color: rgba(255, 81, 0, 0.15);
  }
  .subArea{
    background-color: white;
    position: absolute;
    width: 84%;
    right: 8%;
    min-height: 300px;
    border: 1px solid #e1e1e1;
    z-index: 50;
    padding: 20px;
    display: none;
    border-radius:2px;
  }
  li:hover .subArea{
    display: block;
  }
</style>
