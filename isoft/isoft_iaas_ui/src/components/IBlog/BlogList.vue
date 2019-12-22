<template>
 <div>
   <!-- 热门分类 -->
   <HotCatalogItems @chooseItem="chooseItem"/>

   <div>
    <Row>
      <Col span="17" style="padding: 0 8px 0 0;">
        <div class="isoft_bg_white isoft_pd10">
          <Row class="_search" style="border-bottom: 1px solid #e6e6e6;padding: 20px;height: 62px;">
            <Col span="4" style="text-align: center;font-size: 20px;color: #333;">
              <span v-if="search_type==='_all'">全部分类</span>
              <span v-else-if="search_type==='_hot'">热门分享</span>
              <span v-else-if="search_type==='_personal'">我的分享</span>
              <span v-else>{{search_type}}</span>
            </Col>
            <Col span="3" offset="8" style="text-align: center;"><a href="javascript:;" @click="chooseItem('_all')">全部分类</a></Col>
            <Col span="3" style="text-align: center;"><a href="javascript:;" @click="chooseItem('_hot')">热门博客</a></Col>
            <Col span="3" style="text-align: center;"><a href="javascript:;" @click="chooseItem('_personal')">我的博客</a></Col>
            <Col span="3" style="text-align: center;"><router-link to="/iblog/mine/blog_edit">我也要发布</router-link></Col>
          </Row>

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
                  <Tag v-if="searchblog.to_top > 0" color="error">置顶</Tag>
                  <a style="color: rgba(43,0,206,0.4);font-weight: 600;" @click="chooseItem(searchblog.catalog_name)">所属分类：{{ searchblog.catalog_name }}</a>

                  <span v-if="searchblog.blog_status == -1" style="float: right;color: red;">审核不通过！</span>
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
                    发布于:<Time :time="searchblog.created_time" style="color:red;"/>&nbsp;
                    更新于:<Time :time="searchblog.last_updated_time" style="color:red;"/>&nbsp;
                  </Col>
                  <Col span="2">
                    <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:searchblog.id}}">
                      <span class="isoft_font12"><span style="color: red;">{{searchblog.views}}</span>阅读</span>
                    </router-link>
                  </Col>
                  <Col span="2">
                    <router-link :to="{path:'/iblog/blog_detail',query:{blog_id:searchblog.id}}">
                      <span class="isoft_font12"><span style="color: red;">0</span>条评论</span>
                    </router-link>
                  </Col>
                  <Col span="3">
                    <router-link :to="{path:'/iblog/mine/blog_edit'}">
                      <span class="isoft_font12">我也要发布</span>
                    </router-link>
                  </Col>
                </Row>
              </p>
            </li>
          </ul>

          <Page :total="total" :page-size="offset" show-total show-sizer :styles="{'text-align': 'center','margin-top': '10px'}"
                @on-change="handleChange" @on-page-size-change="handlePageSizeChange"/>
        </div>
      </Col>
        <Col span="7" class="isoft_bg_white isoft_pd10">
          <HotUser/>
          <RandomAdmt2/>
          <CatalogList/>
        </Col>
    </Row>
   </div>

   <div class="isoft_bg_white isoft_pd10" style="margin-top: 8px;">
     <HorizontalLinks :placement_name="GLOBAL.placement_want_to_find"/>
   </div>

 </div>
</template>

<script>
  import HotCatalogItems from "./HotCatalogItems"
  import {queryPageBlog} from "../../api"
  import CatalogList from "./CatalogList"
  import HotUser from "../User/HotUser"
  import HorizontalLinks from "../Elementviewers/HorizontalLinks";
  import IBeautifulLink from "../Common/link/IBeautifulLink";
  import RandomAdmt2 from "../Advertisement/RandomAdmt2";

  export default {
    name: "BlogList",
    components:{RandomAdmt2, IBeautifulLink, HorizontalLinks, CatalogList,HotCatalogItems,HotUser},
    data(){
      return {
        // 当前页
        current_page:1,
        // 总数
        total:0,
        // 每页记录数
        offset:20,
        searchblogs:[],
        search_type:'_all',
      }
    },
    methods:{
      chooseItem:function(item_name){
        if(this.search_type != item_name){
          this.search_type = item_name;
          this.current_page = 1;
          this.refreshBlogList();
        }
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
        var search_type = this.search_type;
        const result = await queryPageBlog(this.offset,this.current_page, search_type);
        if(result.status=="SUCCESS"){
          this.searchblogs = result.blogs;
          this.total = result.paginator.totalcount;
        }
      },
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
    },
  }
</script>

<style scoped>


  a{
    color: #3d3d3d;
  }
  a:hover{
    color: red;
  }
  ._search a{
    color: #155faa;
  }
  ._search a:hover{
    color: #6cb0ca;
  }
</style>
