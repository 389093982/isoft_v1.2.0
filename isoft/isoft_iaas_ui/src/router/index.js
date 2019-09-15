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
const BlogEdit = () => import("@/components/IBlog/BlogEdit");
const BookEdit = () => import("@/components/IBlog/Book/BookEdit");
const BookList2 = () => import("@/components/IBlog/Book/BookList2");
const BookDetail = () => import("@/components/IBlog/Book/BookDetail");
const UserDetail = () => import("@/components/User/UserDetail");
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
    {path: 'book_list',component: BookList2},
    {path: 'book_detail',component: BookDetail},
    {path: 'blog_detail',component: BlogDetail},
    {path: 'mine/blog_edit',component: BlogEdit},
    {path: 'mine/book_edit',component: BookEdit},
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
      redirect: '/ilearning/mine/course_space/newCourse',
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
    {path: 'appRegist',component: AppRegist},
    {path: 'loginRecord',component: LoginRecord},
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
  ]
}];

const IUserReouter = [{
  path: '/user',
  component: ILayout,
  children: [
    {path: 'detail',component: UserDetail},
    {path: 'mine/detail',component: UserDetail},
  ]
}];

function getAllRouters() {
  let allRouters = [];
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
})
