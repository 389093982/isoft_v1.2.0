<template>
 <div>
   <!-- 热门分类 -->
   <HotCatalogItems @chooseItem="chooseItem"/>
   <div style="margin: 0 15px;background-color: #fff;border: 1px solid #e6e6e6;border-radius: 4px;">
    <Row>
      <Col span="16" style="padding: 0 0 20px;border-right: 1px solid #e6e6e6;">
        <div style="border-bottom: 1px solid #e6e6e6;padding: 20px;height: 62px;">
          <Row class="_search">
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
        </div>
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
                <a style="color: #499ef3;font-weight: bold;" @click="chooseItem(searchblog.catalog_name)">所属分类：{{ searchblog.catalog_name }}</a>

                <span style="float: right;" v-if="isAdmin">
                  <span v-if="searchblog.blog_status == 1" style="color: #f16aff;">已启用</span>
                  <span v-else style="color: #f16aff;">已禁用</span>
                  <IBeautifulLink2 @onclick="deleteBlog(searchblog.id, searchblog.blog_status)">启/禁用</IBeautifulLink2>
                </span>
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
        <Col span="8" style="padding: 10px;">
          <HotUser/>
          <CatalogList/>
        </Col>
    </Row>
   </div>

   <HorizontalLinks :placement_name="GLOBAL.want_to_find" style="margin: 20px;"/>
 </div>
</template>

<script>
  import HotCatalogItems from "./HotCatalogItems"
  import {BlogList,UpdateBlogStatus} from "../../api"
  import CatalogList from "./CatalogList"
  import HotUser from "../User/HotUser"
  import HorizontalLinks from "../Elementviewers/HorizontalLinks";
  import IBeautifulLink2 from "../Common/link/IBeautifulLink2";
  import {CheckAdminLogin} from "../../tools"

  export default {
    name: "BlogList",
    components:{IBeautifulLink2, HorizontalLinks, CatalogList,HotCatalogItems,HotUser},
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
        if(this.search_type == "_all"){
          search_type = "";
        }
        const result = await BlogList(this.offset,this.current_page, search_type);
        if(result.status=="SUCCESS"){
          this.searchblogs = result.blogs;
          this.total = result.paginator.totalcount;
        }
      },
      deleteBlog:async function(blog_id, blog_status){
        blog_status = blog_status == 1 ? -1 : 1;
        const result = await UpdateBlogStatus(blog_status, blog_id);
        if(result.status=="SUCCESS"){
          this.refreshBlogList();
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
    },
    computed:{
      isAdmin:function () {
        return CheckAdminLogin();
      }
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
  ._search a{
    color: #155faa;
  }
  ._search a:hover{
    color: #6cb0ca;
  }
</style>
