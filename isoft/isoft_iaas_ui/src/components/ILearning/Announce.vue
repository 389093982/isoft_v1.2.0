<template>
  <div class="textBox">
    <transition name="slide">
      <p class="text" :key="text.id">{{text.val}}</p>
    </transition>
  </div>
</template>

<script>
  export default {
    name: 'Announce',
    data () {
      return {
        textArr: [
          '1 第一条公告',
          '2 第二条公告第二条公告',
          '3 第三条公告第三条公告第三条公告'
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
          if (this.number === 2) {
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
  h2 {
    padding: 20px 0
  }
  .textBox {
    width: 100%;
    height: 40px;
    margin: 0 auto;
    overflow: hidden;
    position: relative;
    text-align: center;
  }
  .text {
    width: 100%;
    position: absolute;
    bottom: 0;
  }
  .slide-enter-active, .slide-leave-active {
    transition: all 0.5s linear;
  }
  .slide-enter{
    transform: translateY(20px) scale(1);
    opacity: 1;
  }
  .slide-leave-to {
    transform: translateY(-20px) scale(0.8);
    opacity: 0;
  }
</style>
