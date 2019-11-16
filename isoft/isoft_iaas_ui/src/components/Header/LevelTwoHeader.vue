<template>
  <div>
    <div style="margin: 5px 20px 0 20px;float: right;" class="clear">
      <ul>
        <li><span @click="$router.push({path:'/ilearning/advise'})">我要吐槽/建议</span></li>
        <li>
          网站导航 <Icon type="ios-arrow-down" />
          <div class="subArea" style="width: 1000px;right: 110px;">
            <div class="arrow"></div>
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
        <li>
          热点 APP <Icon type="ios-arrow-down" />
          <div class="subArea" style="width: 1000px;right: 30px;">
            <div class="arrow"></div>
            BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    name: "LevelTwoHeader"
  }
</script>

<style scoped>
  @import "../../assets/css/isoft_common.css";

  li{
    float: left;
    list-style: none;
    margin-right: 10px;
    display: inline-block;
    cursor: pointer;
  }
  .subArea{
    background-color: white;
    position: absolute;
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
  .arrow{
    position: absolute;
    right: 25px;
    top: -10px;
    width: 0;
    height: 0;
    border-width: 5px;
    border-style: solid;
    border-color: transparent #e1e1e1 transparent transparent;
    transform: rotate(90deg); /*顺时针旋转90°*/
  }
</style>
