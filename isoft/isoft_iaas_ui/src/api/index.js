/*
包含n个接口请求函数的模块
函数的返回值: promise对象
 */
import ajax from './ajax'

const BASE_URL = '/api'

// 编辑或者新增博客分类
export const CatalogEdit = (catalog_name, catalog_desc) => ajax(BASE_URL+'/iwork/httpservice/CatalogEdit2',{catalog_name, catalog_desc},'POST');

// 获取我的所有博客分类
export const GetMyCatalogs = () => ajax(BASE_URL+'/iwork/httpservice/GetMyCatalogs2',{},'GET');

// 获取我的所有博客文章
export const GetMyBlogs = () => ajax(BASE_URL+'/iwork/httpservice/GetMyBlogs2',{},'GET');

// 编辑或者新增博客文章
export const BlogEdit = (blog_title, short_desc, key_words, catalog_id, content) => ajax(BASE_URL+'/iwork/httpservice/BlogEdit2',{blog_title, short_desc, key_words, catalog_id, content},'POST');

// 热门博客分页列表
export const BlogList = (offset,current_page) => ajax(BASE_URL+'/iwork/httpservice/BlogList2',{offset,current_page},'GET');

// 根据 blog_id 查询 blog 详细信息
export const ShowBlogDetail = (blog_id) => ajax(BASE_URL+'/iwork/httpservice/ShowBlogDetail2',{blog_id},'GET');

// 新建课程
export const NewCourse = (course_name,course_type,course_sub_type,course_short_desc) =>
  ajax(BASE_URL+"/iwork/httpservice/NewCourse2",{course_name,course_type,course_sub_type,course_short_desc},'GET');

// 分页查询我的课程清单
export const GetMyCourseList = (userName) => ajax(BASE_URL+"/iwork/httpservice/GetMyCourseList2",{userName},'GET');

// 完结视频更新
export const EndUpdate = (course_id) => ajax(BASE_URL+"/iwork/httpservice/EndUpdate2",{course_id},'GET');

// 显示课程详细信息
export const ShowCourseDetail = (course_id) => ajax(BASE_URL+"/iwork/httpservice/ShowCourseDetail2",{course_id},'GET');

// 切换收藏点赞
export const ToggleFavorite = (favorite_id, favorite_type) => ajax(BASE_URL+"/iwork/httpservice/ToggleFavorite2",{favorite_id, favorite_type},'GET');

// 添加评论
export const AddComment = (parent_id, content, theme_pk, theme_type, comment_type, refer_user_name) =>
  ajax(BASE_URL+'/iwork/httpservice/AddComment2',{parent_id, content, theme_pk, theme_type, comment_type, refer_user_name},'GET');

// 获取评论列表
export const FilterComment = (theme_pk, theme_type, parent_id, comment_type, offset,current_page) =>
  ajax(BASE_URL+'/iwork/httpservice/FilterComment2',{theme_pk, theme_type, parent_id, comment_type, offset,current_page},'GET');

// 获取所有课程类型
export const GetAllCourseType = () => ajax(BASE_URL+"/iwork/httpservice/GetAllCourseType2",{},'GET');

export const ShowCourseHistory = (offset,current_page) => ajax(BASE_URL+"/iwork/httpservice/ShowCourseHistory2", {offset,current_page},'GET');

// 获取热门推荐的课程
export const GetHotCourseRecommend = () => ajax(BASE_URL+"/iwork/httpservice/GetHotCourseRecommend2",{},'GET');

// 根据课程名称获取所有子类型名称
export const GetAllCourseSubType = (course_type) => ajax(BASE_URL+'/iwork/httpservice/GetAllCourseSubType2',{course_type},'GET');

// 课程搜索
export const SearchCourseList = (search) => ajax(BASE_URL+'/iwork/httpservice/SearchCourseList2',{search},'GET');

// 添加配置项
export const AddConfiguration = (parent_id, configuration_name, configuration_value) =>
  ajax(BASE_URL+'/iwork/httpservice/AddConfiguration2',{parent_id, configuration_name, configuration_value},'GET');


// 分页查询配置项信息
export const FilterConfigurations = (search,offset,current_page) => ajax(BASE_URL+'/iwork/httpservice/FilterConfigurations2',{search, offset,current_page},'GET');

// 获取随机数量的友情链接地址
export const QueryRandomCommonLink = (link_type) => ajax(BASE_URL+'/iwork/httpservice/QueryRandomCommonLink2',{link_type},'GET');

// 分页查询友情链接地址
export const FilterCommonLinks = (offset,current_page,search) => ajax(BASE_URL+'/iwork/httpservice/FilterCommonLinks2',{offset,current_page,search},'GET');

// 添加友情链接地址
export const AddCommonLink = (link_type, link_name, link_addr) => ajax(BASE_URL+'/iwork/httpservice/AddCommonLink2',{link_type, link_name, link_addr},'GET');

export const FilterElements = (offset,current_page,search) => ajax(BASE_URL+'/iwork/httpservice/FilterElements2',{offset,current_page,search},'GET');

export const EditElement = (id, placement, navigation_level, navigation_parent_id, title, content, md_content, imgpath, linked_refer) =>
  ajax(BASE_URL+'/iwork/httpservice/EditElement2',{id, placement, navigation_level, navigation_parent_id, title, content, md_content, imgpath, linked_refer},'GET');

export const UpdateElementStatus = (id, status) => ajax(BASE_URL+'/iwork/httpservice/UpdateElementStatus2',{id, status},'GET');

export const EditPlacement = (id, placement_name, placement_desc, placement_label) =>
  ajax(BASE_URL+'/iwork/httpservice/EditPlacement2',{id, placement_name, placement_desc, placement_label},'GET');

export const QueryPlacementById = (id) => ajax(BASE_URL+'/iwork/httpservice/QueryPlacementById',{id},'GET');

export const FilterPlacement = (offset,current_page,search) => ajax(BASE_URL+'/iwork/httpservice/FilterPlacement2',{offset,current_page,search},'GET');
export const GetAllPlacements = () => ajax(BASE_URL+'/iwork/httpservice/GetAllPlacements2',{},'GET');
export const CopyElement = (id) => ajax(BASE_URL+'/iwork/httpservice/CopyElement',{id},'GET');

export const DeletePlacementById = (id) => ajax(BASE_URL+'/iwork/httpservice/DeletePlacementById2',{id},'GET');
export const CopyPlacement = (id) => ajax(BASE_URL+'/iwork/httpservice/CopyPlacement',{id},'GET');

export const FilterElementByPlacement = (placement) => ajax(BASE_URL+'/iwork/httpservice/FilterElementByPlacement2',{placement},'GET');
export const QueryElementById = (id) => ajax(BASE_URL+'/iwork/httpservice/QueryElementById',{id},'GET');

// 根据 blog_id 查询 blog 详细信息
export const ShowShareDetail = (share_id) => ajax(BASE_URL+'/iwork/httpservice/ShowShareDetail2',{share_id},'GET');

// 获取ShareList 信息
export const FilterShareList = (offset,current_page,search_type) => ajax(BASE_URL+'/iwork/httpservice/FilterShareList2',{offset,current_page,search_type},'GET');

// 新增共享链接
export const AddNewShare = (share_type,share_desc,link_href,content) => ajax(BASE_URL+'/iwork/httpservice/AddNewShare2',{share_type,share_desc,link_href,content},'GET');

// 登录接口
export const Login = (username,passwd) => ajax(BASE_URL+"/iwork/httpservice/PostLogin2", {username,passwd},'POST');

// 注册接口
export const Regist = (username,passwd) => ajax(BASE_URL+"/iwork/httpservice/Regist2", {username,passwd},'POST');

export const GetHotUsers = () => ajax(BASE_URL+"/iwork/httpservice/GetHotUsers", {},'POST');

// 系统注册分页查询
export const AppRegisterList = (offset,current_page,search) => ajax(BASE_URL+"/iwork/httpservice/AppRegisterList2", {offset,current_page,search},'POST');

// 添加系统注册
export const AddAppRegister = (app_address) => ajax(BASE_URL+"/iwork/httpservice/AddAppRegister2", {app_address},'POST');

// 登录记录分页查询
export const LoginRecordList = (offset,current_page,search) => ajax(BASE_URL+"/iwork/httpservice/LoginRecordList2", {offset,current_page,search},'POST');


// 跨模块使用,模块化部署时需要使用 nginx 代理
export const LoginAddr = "/sso/login/";
