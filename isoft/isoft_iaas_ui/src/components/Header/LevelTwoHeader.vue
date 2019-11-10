<template>
  <div>
    <div style="margin: 5px 20px 0 20px;float: right;" class="clear">
      <ul>
        <li>
          网站导航 <Icon type="ios-arrow-down" />
          <div class="subArea" style="width: 1000px;right: 110px;">
            BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
          </div>
        </li>
        <li>
          热点 APP <Icon type="ios-arrow-down" />
          <div class="subArea" style="width: 1000px;right: 30px;">
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
  }
</style>
