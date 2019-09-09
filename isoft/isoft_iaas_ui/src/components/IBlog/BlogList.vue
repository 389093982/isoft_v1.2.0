<template>
 <div style="margin-left: 50px;margin-right: 50px;">

   <!-- 热门分类 -->
   <HotCatalogItems @chooseItem="chooseItem"/>

    <Row>
      <Col span="18">
        热门博客 <router-link style="float: right;" :to="{path:'/iblog/mine/blog_edit'}">我也要发布</router-link>
        <ul>
          <li v-for="searchblog in searchblogs" style="list-style:none;padding: 10px 10px;background: #fff;border-bottom: 1px solid #f4f4f4;">
            <Row>
              <Col span="12">
                <router-link to="" style="float: left;">
                  <Avatar size="small" src="https://i.loli.net/2017/08/21/599a521472424.jpg" />
                </router-link>
                <!-- 使用v-bind动态绑定id传递给目标路径 -->
                <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:searchblog.id}}">
                  <h4>{{searchblog.blog_title}}</h4>
                </router-link>
              </Col>
              <Col span="12">
                <router-link :to="{path:'/iblog/blog_list',query:{blog_id:searchblog.id}}">
                  <span style="color: #499ef3;font-weight: bold;">所属分类：{{ searchblog.catalog_name }}</span>
                </router-link>
              </Col>
            </Row>
            <p style="margin-bottom: 4px;font-size: 14px;color: #8a8a8a;line-height: 24px;">
              {{searchblog.short_desc | filterLimitFunc}}
            </p>
            <p>
              <Row>
                <Col span="17">
                  <!-- 作者详情 -->
                  <router-link :to="{path:'/iblog/author',query:{author:searchblog.author}}">{{searchblog.author}}</router-link>
                  发布于:<Time :time="searchblog.created_time" type="datetime" style="color:red;"/>&nbsp;
                  更新于:<Time :time="searchblog.last_updated_time" type="datetime" style="color:red;"/>&nbsp;
                </Col>
                <Col span="2">
                  <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:searchblog.id}}">
                    <span style="color: red;">{{searchblog.views}}</span>阅读
                  </router-link>
                </Col>
                <Col span="2">
                  <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:searchblog.id}}">
                    <span style="color: red;">0</span>条评论
                  </router-link>
                </Col>
                <Col span="3">
                  <router-link :to="{path:'/iblog/mine/blog_edit'}">我也要发布</router-link>
                </Col>
              </Row>
            </p>
          </li>
        </ul>

        <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
              @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
      </Col>
      <Col span="6">
        <CatalogList/>
      </Col>
    </Row>
 </div>
</template>

<script>
  import HotCatalogItems from "../Share/HotCatalogItems"
  import {BlogList} from "../../api"
  import CatalogList from "./CatalogList"

  export default {
    name: "BlogList",
    components:{CatalogList,HotCatalogItems},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:10,
        searchblogs:[],
      }
    },
    methods:{
      chooseItem:function(item_name){

      },
      handleChange(page){
        this.current_page = page;
        this.refreshBlogList();
      },
      handlePageSizeChange(pageSize){
        this.offset = pageSize;
        this.refreshBlogList();
      },
      refreshBlogList:async function () {
        const result = await BlogList(this.offset,this.current_page);
        if(result.status=="SUCCESS"){
          this.searchblogs = result.blogs;
          this.total = result.paginator.totalcount;
        }
      }
    },
    mounted: function () {
      this.refreshBlogList();
    },
    filters:{
      // 内容超长则显示部分
      filterLimitFunc:function (value) {
        if(value && value.length > 100) {
          value= value.substring(0,100) + '...';
        }
        return value;
      },
    }
  }
</script>

<style scoped>
  a{
    color: #3d3d3d;
  }
  a:hover{
    color: red;
  }
</style>
