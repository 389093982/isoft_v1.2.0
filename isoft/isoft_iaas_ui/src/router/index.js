import Vue from 'vue'
import Router from 'vue-router'

import Login from "../components/SSO/Login/Login"
import Regist from "../components/SSO/Login/Regist"
import AppRegist from "../components/SSO/AppRegist"
import LoginRecord from "../components/SSO/LoginRecord"
import IEmptyLayout from "../components/ILayout/IEmptyLayout"
import ICMSLayout from "../components/ILayout/ICMSLayout"
import Element from "../components/Background/CMS/Element"
import Catalog from "../components/Background/CMS/Catalog"
import Placement from "../components/Background/CMS/Placement"
import IBlog from '../components/IBlog/IBlog'
import BlogList from '../components/IBlog/BlogList'
import CatalogAdd from '../components/IBlog/CatalogAdd'
import BlogAdd from '../components/IBlog/BlogAdd'
import BlogDetail from '../components/IBlog/BlogDetail'
import ILearningIndex from '../components/ILearning/Index'
import CourseSpace from '../components/ILearning/CourseSpace/CourseSpace'
import NewCourse from '../components/ILearning/CourseSpace/NewCourse'
import RecentlyViewed from '../components/ILearning/CourseSpace/RecentlyViewed'
import MyCourseList from '../components/ILearning/CourseSpace/MyCourseList'
import CourseDetail from '../components/ILearning/Course/CourseDetail'
import VideoPay from '../components/ILearning/Course/VideoPay'
import Configuration from '../components/Background/CMS/Configuration'
import CourseSearch from "../components/ILearning/Course/CourseSearch"
import ShareAdd from "../components/Share/ShareAdd"
import ShareList from "../components/Share/ShareList"
import ShareDetail from "../components/Share/ShareDetail"
import CommonLinkList from "../components/Background/CMS/CommonLinkList"
import ILayout from "../components/ILayout/ILayout"


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
    {path: 'blog_add',component: BlogAdd},
    {path: 'blog_list',component: BlogList},
    {path: 'blog_detail',component: BlogDetail},
  ]
};

const ShareListRouter = {
  path: '/share',
  component: ILayout,
  children: [
    {path: 'add',component: ShareAdd,},
    {path: 'list',component: ShareList,},
    {path: 'detail',component: ShareDetail,},
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
      path: 'course_space',
      component: CourseSpace,
      redirect: '/ilearning/course_space/newCourse',
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

const ILearningRouters = [IBlogRouter, ILearningRouter, ShareListRouter];

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
    {path: 'cms/element_list',component: Element},
    {path: 'cms/catalog_list',component: Catalog},
    {path: 'cms/placement_list',component: Placement},
    {path: 'cms/commonLinkList',component: CommonLinkList},
    {path: 'cms/configuration',component: Configuration},
  ]
}];

function getAllRouters() {
  let allRouters = [];
  allRouters = joinArray(allRouters, ILearningRouters);
  allRouters = joinArray(allRouters, ISSOReouter);
  allRouters = joinArray(allRouters, ICMSReouter);
  allRouters = joinArray(allRouters, getRootRouters());
  return allRouters;
}

export default new Router({
  // # 主要用来区分前后台应用, history 模式需要使用 nginx 代理
  // History 模式,去除vue项目中的 #
  mode: 'history',
  routes: getAllRouters(),
})
