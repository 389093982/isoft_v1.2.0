<template>
  <div class="textBox" style="padding-left: 50px;">
    <img src="../../assets/notice.png" style="position: relative; top: 5px;left: -30px;"/>
    <transition name="slide">
      <p class="text" :key="text.id" v-html="text.val"></p>
    </transition>
  </div>
</template>

<script>
  export default {
    name: 'Announce',
    data () {
      return {
        textArr: [
          "1 第一条公告第一条公告第一条公告第一<span style='color:red;font-weight: bold;'>条公告第一条公告第一条</span>公告第一条公告",
          "2 第二条公告第二条公告第一条公告第一条公告第一条公告第一条公告第一条公告第一条公告",
          "4 使用 Beego 的产品<a>通过 GitHub 提交案例</a> ，如果使用GitHub有任何问题，你可以直接发送产品信息到邮箱 <a>xiemengjun@gmail.com</a>",
          "3 第三条公告第三条公告第三条公告第一条公告第一条公告第一条公告第一条公告第一条公告",
        ],
        number: 0
      }
    },
    computed: {
      text () {
        return {
          id: this.number,
          val: this.textArr[this.number]
        }
      }
    },
    mounted () {
      this.startMove()
    },
    methods: {
      startMove () {
        // eslint-disable-next-line
        let timer = setTimeout(() => {
          if (this.number === 3) {        // number 最大行数索引
            this.number = 0;
          } else {
            this.number += 1;
          }
          this.startMove();
        }, 2000); // 滚动不需要停顿则将2000改成动画持续时间
      }
    }
  }
</script>

<style scoped>
  .textBox {
    width: 100%;
    height: 40px;
    overflow: hidden;
    position: relative;
    border: 2px solid rgba(220,220,220,0.28);
    margin: 5px 0 5px 0;
    padding: 5px;
  }
  .text {
    width: 100%;
    position: absolute;
    bottom: 0;
  }
  .slide-enter-active, .slide-leave-active {
    transition: all 0.2s linear;
  }
  .slide-enter{
    transform: translateY(20px) scale(1);
    opacity: 1;
  }
  .slide-leave-to {
    /*transform: translateY(-20px) scale(0.8);*/
    transform: translateY(-20px) scale(1);
    opacity: 0;
  }
</style>
