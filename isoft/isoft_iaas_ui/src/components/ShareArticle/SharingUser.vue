<template>
  <div>
    <Row v-if="article.title">
      <Col span="2">
        <a class="demo-avatar">
          <Avatar style="background-color: #87d068" icon="ios-person" />
        </a>
      </Col>
      <Col span="14">
        <p>二月：<a>{{article.title}}</a></p>
        <p style="color:#9ea7b4; font-size: 12px">7 人关注 • 20 个回复 • 208 次浏览 • 1 小时前</p>
      </Col>
      <Col span="8">
        <Tooltip content="好评" placement="top-end">
          <a><Icon type="ios-thumbs-up" style="color: lightsalmon;" size="20"/></a>
        </Tooltip>
        <Tooltip content="差评" placement="top-end">
          <a><Icon type="ios-thumbs-down" style="color: lightsalmon;" size="20"/></a>
        </Tooltip>
        <Tooltip content="收藏" placement="top-end">
          <a><Icon type="md-cube" style="color: lightsalmon;" size="20"/></a>
        </Tooltip>
        <Tooltip content="评论" placement="top-end">
          <a><Icon type="md-text" style="color: lightsalmon;" size="20"/></a>
        </Tooltip>
        <Tooltip content="转发" placement="top-end">
          <a><Icon type="ios-send" style="color: lightsalmon;" size="20"/></a>
        </Tooltip>
      </Col>
    </Row>
    <Col></Col>
  </div>
</template>

<script>
	export default {
	  props:{
	    article:{
	      title:String,
        content:String,
      },
    },
		name: "SharingUser",
    data(){
		  return{

      }
    },
	}
</script>

<style scoped>

</style>
