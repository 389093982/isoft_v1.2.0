/*
包含n个接口请求函数的模块
函数的返回值: promise对象
 */
import ajax from './ajax'
import store from "../store"

const BASE_URL = '/api'

// 编辑或者新增博客分类
export const CatalogEdit = (catalog_name, catalog_desc) => ajax(BASE_URL+'/catalog/edit',{catalog_name, catalog_desc},'POST');

// 获取我的所有博客分类
export const GetMyCatalogs = () => ajax(BASE_URL+'/catalog/getMyCatalogs',{},'GET');

// 获取我的所有博客文章
export const GetMyBlogs = () => ajax(BASE_URL+'/blog/getMyBlogs',{},'GET');

// 编辑或者新增博客文章
export const BlogEdit = (blog_title, short_desc, key_words, catalog_id, content) => ajax(BASE_URL+'/blog/edit',{blog_title, short_desc, key_words, catalog_id, content},'POST');

// 热门博客分页列表
export const BlogList = (offset,current_page) => ajax(BASE_URL+'/blog/blogList',{offset,current_page},'GET');

// 根据 blog_id 查询 blog 详细信息
export const ShowBlogDetail = (blog_id) => ajax(BASE_URL+'/blog/showBlogDetail',{blog_id},'GET');

// 新建课程
export const NewCourse = (course_name,course_type,course_sub_type,course_short_desc) =>
  ajax(BASE_URL+'/ilearning/newCourse',{course_name,course_type,course_sub_type,course_short_desc},'GET');

// 分页查询我的课程清单
export const GetMyCourseList = (userName) => ajax(BASE_URL+'/ilearning/getMyCourseList',{userName},'GET');

// 完结视频更新
export const EndUpdate = (course_id) => ajax(BASE_URL+'/ilearning/endUpdate',{course_id},'GET');

// 显示课程详细信息
export const ShowCourseDetail = (course_id) => ajax(BASE_URL+'/ilearning/showCourseDetail',{course_id},'GET');

// 切换收藏点赞
export const ToggleFavorite = (favorite_id, favorite_type) => ajax(BASE_URL+'/ilearning/toggle_favorite',{favorite_id, favorite_type},'GET');

// 查询评论主题
export const FilterCommentTheme = (comment_id, theme_type) => ajax(BASE_URL+'/ilearning/filterCommentTheme',{comment_id, theme_type},'GET');

// 添加评论
export const AddCommentReply = (parent_id, reply_content, comment_id, theme_type, reply_comment_type, refer_user_name) =>
  ajax(BASE_URL+'/ilearning/addCommentReply',{parent_id, reply_content, comment_id, theme_type, reply_comment_type, refer_user_name},'GET');

// 获取评论列表
export const FilterCommentReply = (comment_id, theme_type, parent_id, reply_comment_type) =>
  ajax(BASE_URL+'/ilearning/filterCommentReply',{comment_id, theme_type, parent_id, reply_comment_type},'GET');

// 获取所有课程类型
export const GetAllCourseType = () => ajax(BASE_URL+'/ilearning/getAllCourseType',{},'GET');

// 获取热门推荐的课程
export const GetHotCourseRecommend = () => ajax(BASE_URL+'/ilearning/getHotCourseRecommend',{},'GET');

// 根据课程名称获取所有子类型名称
export const GetAllCourseSubType = (course_type) => ajax(BASE_URL+'/ilearning/getAllCourseSubType',{course_type},'GET');

// 课程搜索
export const SearchCourseList = (search) => ajax(BASE_URL+'/ilearning/searchCourseList',{search},'GET');

// 添加配置项
export const AddConfiguration = (parent_id, configuration_name, configuration_value) =>
  ajax(BASE_URL+'/cms/addConfiguration',{parent_id, configuration_name, configuration_value},'GET');

// 根据名称查询配置项
export const QueryAllConfigurations = (configuration_name) => ajax(BASE_URL+'/cms/queryAllConfigurations',{configuration_name},'GET');

// 分页查询配置项信息
export const FilterConfigurations = (search,offset,current_page) => ajax(BASE_URL+'/cms/filterConfigurations',{search, offset,current_page},'GET');

// 获取随机数量的友情链接地址
export const QueryRandomCommonLink = (link_type) => ajax(BASE_URL+'/cms/queryRandomCommonLink',{link_type},'GET');

// 分页查询友情链接地址
export const FilterCommonLinks = (offset,current_page,search) => ajax(BASE_URL+'/cms/filterCommonLinks',{offset,current_page,search},'GET');

// 添加友情链接地址
export const AddCommonLink = (link_type, link_name, link_addr) => ajax(BASE_URL+'/cms/addCommonLink',{link_type, link_name, link_addr},'GET');

// 根据 blog_id 查询 blog 详细信息
export const ShowShareDetail = (share_id) => ajax(BASE_URL+'/share/showShareDetail',{share_id},'GET');

// 获取ShareList 信息
export const FilterShareList = (offset,current_page,search_type) => ajax(BASE_URL+'/share/filterShareList',{offset,current_page,search_type},'GET');


// 新增共享链接
export const AddNewShare = (share_type,share_desc,link_href,content) => ajax(BASE_URL+'/share/addNewShare',{share_type,share_desc,link_href,content},'GET');

export const ShowCourseHistory = (offset,current_page) => ajax(BASE_URL+"/common/showCourseHistory", {offset,current_page},'GET');

// 注册应用心跳检测项
export const RegisterHeartBeat = (addr) => ajax(BASE_URL+"/monitor/registerHeartBeat", {addr},'GET');

// 分页获取应用心跳检测项
export const FilterPageHeartBeat = (offset,current_page) => ajax(BASE_URL+"/monitor/filterPageHeartBeat", {offset,current_page},'GET');

// 分页获取 ifile 清单
export const FilterPageIFiles = (search_name,offset,current_page) => ajax(BASE_URL+"/ifile/filterPageIFiles", {search_name,offset,current_page},'GET');

// 登录接口
export const Login = (username,passwd) => ajax(BASE_URL+"/sso/user/login", {username,passwd},'POST');

// 注册接口
export const Regist = (username,passwd) => ajax(BASE_URL+"/sso/user/regist", {username,passwd},'POST');

// 系统注册分页查询
export const AppRegisterList = (offset,current_page,search) => ajax(BASE_URL+"/sso/app/appRegisterList", {offset,current_page,search},'POST');

// 添加系统注册
export const AddAppRegister = (app_address) => ajax(BASE_URL+"/sso/app/addAppRegister", {app_address},'POST');

// 登录记录分页查询
export const LoginRecordList = (offset,current_page,search) => ajax(BASE_URL+"/sso/user/loginRecordList", {offset,current_page,search},'POST');


// 跨模块使用,模块化部署时需要使用 nginx 代理
export const LoginAddr = "/sso/login/";
