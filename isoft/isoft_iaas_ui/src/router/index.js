import Vue from 'vue'
import Router from 'vue-router'

// es6 import 异步语法,使用异步组件加载机制减少耗时操作
const Login = () => import("@/components/SSO/Login/Login");
const Regist = () => import("@/components/SSO/Login/Regist");
const AppRegist = () => import("@/components/SSO/AppRegist");
const LoginRecord = () => import("@/components/SSO/LoginRecord");
const IEmptyLayout = () => import("@/components/ILayout/IEmptyLayout");
const ICMSLayout = () => import("@/components/ILayout/ICMSLayout");
const Element = () => import("@/components/Background/CMS/Element");
const EditElement = () => import("@/components/Background/CMS/EditElement");
const Catalog = () => import("@/components/Background/CMS/Catalog");
const Placement = () => import("@/components/Background/CMS/Placement");
const EditPlacement = () => import("@/components/Background/CMS/EditPlacement");
const IBlog = () => import("@/components/IBlog/IBlog");
const BlogList = () => import("@/components/IBlog/BlogList");
const BlogDetail = () => import("@/components/IBlog/BlogDetail");
const ArticleEdit = () => import("@/components/IBlog/ArticleEdit");
const BookCatalogEdit = () => import("@/components/IBlog/Book/BookCatalogEdit");
const BookList = () => import("@/components/IBlog/Book/BookList");
const BookArticleList = () => import("@/components/IBlog/Book/BookArticleList");
const UserDetail = () => import("@/components/User/UserDetail");
const UserGuide = () => import("@/components/User/UserGuide");
const ILearningIndex = () => import("@/components/ILearning/Index");
const CourseSpace = () => import("@/components/ILearning/CourseSpace/CourseSpace");
const NewCourse = () => import("@/components/ILearning/CourseSpace/NewCourse");
const RecentlyViewed = () => import("@/components/ILearning/CourseSpace/RecentlyViewed");
const MyCourseList = () => import("@/components/ILearning/CourseSpace/MyCourseList");
const CourseDetail = () => import("@/components/ILearning/Course/CourseDetail");
const VideoPay = () => import("@/components/ILearning/Course/VideoPay");
const Configuration = () => import("@/components/Background/CMS/Configuration");
const CourseSearch = () => import("@/components/ILearning/Course/CourseSearch");
const CommonLinkList = () => import("@/components/Background/CMS/CommonLinkList");
const FoundList = () => import("@/components/IFound/FoundList");
const GoodList = () => import("@/components/IGood/GoodList");
const GoodEdit = () => import("@/components/IGood/GoodEdit");
const GoodDetail = () => import("@/components/IGood/GoodDetail");
const PayConfirm = () => import("@/components/IGood/PayConfirm");
const ILayout = () => import("@/components/ILayout/ILayout");


import {joinArray} from "../tools"

Vue.use(Router);

function getRootRouters () {
  return [{
    path: '/',
    redirect: '/ilearning/index'
  }]
};

const IBlogRouter = {
  path: '/iblog',
  component: ILayout,
  // 二级路由的配置
  children: [
    {path: 'blog_index',component: IBlog},
    {path: 'blog_list',component: BlogList},
    {path: 'book_list',component: BookList},
    {path: 'book_detail',component: BookArticleList},
    {path: 'blog_detail',component: BlogDetail},
    {path: 'mine/blog_edit',component: ArticleEdit},
    {path: 'mine/book_edit',component: BookCatalogEdit},
    {path: 'mine/book_list',component: BookList},
  ]
};

const ILearningRouter = {
  path: '/ilearning',
  component: ILayout,
  // 二级路由的配置
  children: [
    {
      path: 'index',
      component: ILearningIndex,
    },
    {
      path: 'mine/course_space',
      component: CourseSpace,
      redirect: '/ilearning/mine/course_space/myCourseList',
      children: [
        {path: 'newCourse',component: NewCourse,},
        {path: 'myCourseList',component: MyCourseList,},
        {path: 'RecentlyViewed',component: RecentlyViewed,},
      ]
    },
    {
      path: 'course_detail',
      component: CourseDetail,
    },
    {
      path: 'video_play',
      component: VideoPay,
    },
    {
      // this.$router.push({ name: 'xxx'});
      // this.$router.push({ path: 'xxx'});
      name:'course_search',
      path: 'course_search',
      component: CourseSearch,
    },
  ]
};

const ILearningRouters = [IBlogRouter, ILearningRouter];

const ISSOReouter = [{
  path: '/sso',
  component: IEmptyLayout,
  children: [
    {path: 'login',component: Login},
    {path: 'regist',component: Regist},
  ]
}];

const ICMSReouter = [{
  path: '/background',
  component: ICMSLayout,
  children: [
    {path: 'cms/element_edit',component: EditElement},
    {path: 'cms/element_list',component: Element},
    {path: 'cms/catalog_list',component: Catalog},
    {path: 'cms/placement_list',component: Placement},
    {path: 'cms/placement_edit',component: EditPlacement},
    {path: 'cms/commonLinkList',component: CommonLinkList},
    {path: 'cms/configuration',component: Configuration},
    {path: 'sso/appRegist',component: AppRegist},
    {path: 'sso/loginRecord',component: LoginRecord},
  ]
}];

const IUserReouter = [{
  path: '/user',
  component: ILayout,
  children: [
    {path: 'detail',component: UserDetail},
    {path: 'guide',component: UserGuide},
    {path: 'mine/detail',component: UserDetail},
  ]
}];

const IFoundReouter = [{
  path: '/ifound',
  component: ILayout,
  children: [
    {path: 'found_list',component: FoundList},
  ]
}];

const IGoodReouter = [{
  path: '/igood',
  component: ILayout,
  children: [
    {path: 'good_list',component: GoodList},
    {path: 'mine/good_list',component: GoodList},
    {path: 'mine/good_edit',component: GoodEdit},
    {path: 'good_detail',component: GoodDetail},
    {path: 'pay_confirm',component: PayConfirm},
  ]
}];

function getAllRouters() {
  let allRouters = [];
  allRouters = joinArray(allRouters, IFoundReouter);
  allRouters = joinArray(allRouters, IGoodReouter);
  allRouters = joinArray(allRouters, IUserReouter);
  allRouters = joinArray(allRouters, ILearningRouters);
  allRouters = joinArray(allRouters, ISSOReouter);
  allRouters = joinArray(allRouters, ICMSReouter);
  allRouters = joinArray(allRouters, getRootRouters());
  return allRouters;
}

export default new Router({
  // # 主要用来区分前后台应用, history 模式需要使用 nginx 代理
  // History 模式,去除vue项目中的 #
  // mode: 'history',
  routes: getAllRouters(),
  // 页面跳转时,让页面滚动在顶部
  scrollBehavior(to,from,savedPosition){
    // from 和 to 相同路由页面不滚动到顶部
    if(from.path === to.path){
      return;
    }
    if(savedPosition){
      return savedPosition;
    }else{
      return {x:0,y:0}
    }
  },
})
