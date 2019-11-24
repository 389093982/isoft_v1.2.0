<template>
  <div>
      <form>
        <div>{{onView}}</div>
        <!--标题-->
        <div class="SingleArticle-title">
          <Input v-model="article.title" placeholder="Enter title..." style="width: 100%" />
        </div>
        <!--内容-->
        <div class="SingleArticle-content">
          <Input type="textarea" v-model="article.content" :rows="rows" :maxlength="max" show-word-limit="true" placeholder="Enter content" style="width:100%"/>
        </div>
        <!--按钮-->
            <Button type="primary">发布</Button>
            <Button type="success">保存</Button>
        <div>

        </div>
        <slot name="title"></slot>
        <slot name="content"></slot>
      </form>
  </div>
</template>

<script>
    export default {
      name: "ArticleEdit",
      components:{},
      data(){
          return{
            onView:"正在预览",
            article:{
              title:"送东阳马生序",
              content:"        余幼时即嗜学。家贫，无从致书以观，每假借于藏书之家，" +
              "手自笔录，计日以还。天大寒，砚冰坚，手指不可屈伸，弗之怠。录毕，走送之" +
              "，不敢稍逾约。以是人多以书假余，余因得遍观群书。既加冠，益慕圣贤之道 ，" +
              "又患无硕师、名人与游，尝趋百里外，从乡之先达执经叩问。先达德隆望尊，" +
              "门人弟子填其室，未尝稍降辞色。余立侍左右，援疑质理，俯身倾耳以请；或" +
              "遇其叱咄，色愈恭，礼愈至，不敢出一言以复；俟其欣悦，则又请焉。故余虽" +
              "愚，卒获有所闻。\n" +
              "        当余之从师也，负箧曳屣，行深山巨谷中，穷冬烈风，大雪深数尺，足" +
              "肤皲裂而不知。至舍，四支僵劲不能动，媵人持汤沃灌，以衾拥覆，久而乃和。" +
              "寓逆旅，主人日再食，无鲜肥滋味之享。同舍生皆被绮绣，戴朱缨宝饰之帽，腰白" +
              "玉之环，左佩刀，右备容臭，烨然若神人；余则缊袍敝衣处其间，略无慕艳意。" +
              "以中有足乐者，不知口体之奉不若人也。盖余之勤且艰若此。"
            },
            rows:15,
            max:2000,
          }
      }
    }
</script>

<style scoped>
  .SingleArticle-title{color: #ff6900}
  .SingleArticle-content{color: #1b1b44}
</style>
