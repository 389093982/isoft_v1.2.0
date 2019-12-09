<template>
  <div style="height: 850px;padding-top: 10px;">
    <h2 class="good_rank">精品排行</h2>
    <a href="#" v-for="(rank, index) in ranks">
      <Row>
        <Col span="2" :class="index < 3 ? 'rank_red_index' : 'rank_grey_index'">{{index + 1}}</Col>
        <Col span="17">
            <span class="rank_name">{{rank.rank_name}}</span>
          <span class="rank_desc">{{rank.rank_desc}}</span>
        </Col>
        <Col span="5" class="rank_label">{{rank.rank_label}}</Col>
      </Row>
    </a>
  </div>
</template>

<script>
  export default {
    name: "GoodRank",
    data(){
      return {
        ranks:[
          {rank_name:'哪吒之魔童降世', rank_desc:'东方英雄逆天改命', rank_label:'2554.0万'},
          {rank_name:'哪吒之魔童降世', rank_desc:'东方英雄逆天改命', rank_label:'2554.0万'},
          {rank_name:'哪吒之魔童降世', rank_desc:'东方英雄逆天改命', rank_label:'2554.0万'},
          {rank_name:'哪吒之魔童降世', rank_desc:'东方英雄逆天改命', rank_label:'2554.0万'},
          {rank_name:'哪吒之魔童降世', rank_desc:'东方英雄逆天改命', rank_label:'2554.0万'},
          {rank_name:'哪吒之魔童降世', rank_desc:'东方英雄逆天改命', rank_label:'2554.0万'},
          {rank_name:'哪吒之魔童降世', rank_desc:'东方英雄逆天改命', rank_label:'2554.0万'},
          {rank_name:'哪吒之魔童降世', rank_desc:'东方英雄逆天改命', rank_label:'2554.0万'},
        ],
      }
    }
  }
</script>

<style scoped>
  .good_rank{
    margin: 0 10px 10px 10px;
    position: relative;
    height: 40px;
    color: #111;
    font-size: 20px;
    font-weight: 400;
    line-height: 40px;
    white-space: nowrap;
  }
  a{
    margin: 0 20px;
    padding: 10px 10px 5px 10px;
    border-top: 1px solid #eee;
    display: block;
    position: relative;
    cursor: pointer;
  }
  .rank_red_index{
    color: #ff183e;
    font-size: 20px;
  }
  .rank_grey_index{
    color: #999;
    font-size: 20px;
  }
  .rank_name{
    color: #111;
    font-size: 16px;
    font-weight: 500;
    line-height: 20px;
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  .rank_name:hover{
    color: red;
  }
  .rank_desc{
    color: #999;
    font-size: 13px;
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
  .rank_label{
    color: #999;
    font-size: 13px;
    white-space: nowrap;
  }
</style>
