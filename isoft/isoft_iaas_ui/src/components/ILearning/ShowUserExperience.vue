<template>
  <div style="margin: 10px 0 30px 0;">
    <div style="width:690px;text-align: center;background: rgba(220,220,220,0.28);margin: 10px 0 10px 0;padding: 10px;">
      <p style="font-size: 20px;" class="hovered hvr-grow hoverLinkColor">用户体验</p>
    </div>
    <Row>
      <Col span="18">
        <a :href="this.userexs[currentIndex].link">
          <img :src="this.userexs[currentIndex].img" width="525px;" height="360px;" style="margin-top: 5px;"/>
          <span style="position: absolute;left: 0px;bottom: 7px;height: 50px;width: 525px;padding-left: 20px;
          background-color: rgba(0,0,0,0.6);line-height: 50px;color: white;font-size: 14px;">
          {{userexs[currentIndex].label}}
        </span>
        </a>
      </Col>
      <Col span="6" style="padding-left: 12px;">
        <ul>
          <li v-for="(userex,index) in userexs" style="list-style: none;display: block;margin-top: 5px;">
            <span style="width: 13px;height: 17px;float: left;position: relative;left: -3px;top: 35px;">
              <img v-show="userex.choosed" src="../../assets/arrow_l.jpg"/>
            </span>
            <HoverBigImg @mouseenter="overShow(index)" :srcImg="userex.img" width="144px"/>
          </li>
        </ul>
      </Col>
    </Row>
  </div>
</template>

<script>
  import HoverBigImg from "../Common/img/HoverBigImg";
  export default {
    name: "ShowUserExperience",
    components: {HoverBigImg}, // 用户体验展示
    data(){
      return {
        currentIndex:0,
        userexs:[
          {
            img:require("../../assets/1.jpg"),
            link:"http://www.baidu.com",
            label:"eeeeeee11111111111111111111",
            choosed:true,
          },
          {
            img:require("../../assets/default.png"),
            link:"http://www.baidu.com",
            label:"eeeeeee22222222222222222",
            choosed:false,
          },
          {
            img:require("../../assets/1.jpg"),
            link:"http://www.baidu.com",
            label:"eeeeeee333333333333333333333333",
            choosed:false,
          },
          {
            img:require("../../assets/default.png"),
            link:"http://www.baidu.com",
            label:"eeeeeee44444444444444444444444444444444",
            choosed:false,
          }
        ]
      }
    },
    methods:{
      overShow:function (index) {
        this.currentIndex = index;
        var _userexs = this.userexs;
        for(var i=0; i<this.userexs.length; i++){
          if(i == index){
            _userexs[i].choosed = true;
          }else{
            _userexs[i].choosed = false;
          }
        }
        this.$set(this.userexs, _userexs);
      }
    }
  }
</script>

<style scoped>

</style>
