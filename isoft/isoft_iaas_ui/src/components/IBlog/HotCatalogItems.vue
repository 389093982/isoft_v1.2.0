<template>
  <ElementsLoader :placement_name="placement_name" @onLoadElement="onLoadElement"
                  style="margin-left: 50px;margin-right: 50px;margin-bottom: 20px;">
    <Row>
      <Col span="2" style="text-align: center;padding-top: 25px;">
        <a href="javascript:;" @click="previous" v-show="showPrevious"><img src="/static/images/arrow_left.png"/></a>
      </Col>
      <Col span="5" v-for="element in getCurrentPage">
        <a href="javascript:;" style="color: #999;" @click="chooseItem(element.linked_refer)">
          <div class="item" style="padding:10px; height: 100px;">
            <Row>
              <Col span="6">
                <img :src="element.img_path" :alt="element.element_label" width="50px" height="50px" @error="defImg()"/>
              </Col>
              <Col span="18" style="padding-left: 5px;">
                <p class="share_catalog_name">{{element.element_label}}</p>
                <p style="font-size: 12px;">{{element.content}}</p>
              </Col>
            </Row>
          </div>
        </a>
      </Col>
      <Col span="2" style="text-align: center;padding-top: 25px;">
        <a href="javascript:;" @click="next"  v-show="showNext"><img src="/static/images/arrow_right.png"/></a>
      </Col>
    </Row>
  </ElementsLoader>
</template>

<script>
  import {FilterElementByPlacement} from "../../api"
  import ElementsLoader from "../Background/CMS/ElementsLoader";

  export default {
    name: "HotCatalogItems",
    components: {ElementsLoader},
    data(){
      return {
        placement_name:this.GLOBAL.placement_host_recommend_blog_tpyes,
        // 热门分享类型
        elements: [],
        // 当前页
        currentPageNo:1,
        defaultImg: require('../../assets/default.png'),
      }
    },
    computed:{
      getCurrentPage:function(){
        return this.pagination(this.currentPageNo,4,this.elements);
      },
      showPrevious:function () {
        const total_page = Math.ceil(this.elements.length / 4);
        return this.currentPageNo > 1;
      },
      showNext:function () {
        const total_page = Math.ceil(this.elements.length / 4);
        return this.currentPageNo < total_page;
      }
    },
    methods:{
      defImg(){
        let img = event.srcElement;
        img.src = this.defaultImg;
        img.onerror = null; //防止闪图
      },
      // 获取前一页
      previous: function(){
        if(this.currentPageNo > 1){
          this.currentPageNo = this.currentPageNo - 1;
        }
      },
      // 获取后一页
      next: function(){
        const total_page = Math.ceil(this.elements.length / 4);
        if(this.currentPageNo < total_page){
          this.currentPageNo = this.currentPageNo + 1;
        }
      },
      // 对数组进行也分
      pagination: function (pageNo, pageSize, array) {
        let offset = (pageNo - 1) * pageSize;
        return (offset + pageSize >= array.length) ? array.slice(offset, array.length) : array.slice(offset, offset + pageSize);
      },
      chooseItem:function (share_name) {
        this.$emit('chooseItem',share_name);
      },
      onLoadElement:function (placement_label, elements) {
        this.elements = elements;
      }
    },
  }
</script>

<style scoped>
  *{
    font-family: 'Helvetica Neue', 'STHeiti', '微软雅黑', 'Microsoft YaHei', Helvetica,Arial,sans-serif;
  }
  .item:hover{
    background: #ebebeb;
    border-radius: 10px;
  }
  .share_catalog_name{
    color:green;
  }
  .share_catalog_name:hover{
    color:red;
  }
</style>
