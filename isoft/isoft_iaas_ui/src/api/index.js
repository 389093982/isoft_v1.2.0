/*
包含n个接口请求函数的模块
函数的返回值: promise对象
 */
import ajax from './ajax'

const BASE_URL = '/api'
const BASE_URL_2 = '/api2'

// 编辑或者新增博客分类
// export const CatalogEdit = (catalog_name, catalog_desc) => ajax(BASE_URL+'/catalog/edit',{catalog_name, catalog_desc},'POST');
export const CatalogEdit = (catalog_name, catalog_desc) => ajax(BASE_URL_2+'/iwork/httpservice/CatalogEdit2',{catalog_name, catalog_desc},'POST');

// 获取我的所有博客分类
// export const GetMyCatalogs = () => ajax(BASE_URL+'/catalog/getMyCatalogs',{},'GET');
export const GetMyCatalogs = () => ajax(BASE_URL_2+'/iwork/httpservice/GetMyCatalogs2',{},'GET');

// 获取我的所有博客文章
// export const GetMyBlogs = () => ajax(BASE_URL+'/blog/getMyBlogs',{},'GET');
export const GetMyBlogs = () => ajax(BASE_URL_2+'/iwork/httpservice/GetMyBlogs2',{},'GET');

// 编辑或者新增博客文章
// export const BlogEdit = (blog_title, short_desc, key_words, catalog_id, content) => ajax(BASE_URL+'/blog/edit',{blog_title, short_desc, key_words, catalog_id, content},'POST');
export const BlogEdit = (blog_title, short_desc, key_words, catalog_id, content) => ajax(BASE_URL_2+'/iwork/httpservice/BlogEdit2',{blog_title, short_desc, key_words, catalog_id, content},'POST');

// 热门博客分页列表
// export const BlogList = (offset,current_page) => ajax(BASE_URL+'/blog/blogList',{offset,current_page},'GET');
export const BlogList = (offset,current_page) => ajax(BASE_URL_2+'/iwork/httpservice/BlogList2',{offset,current_page},'GET');

// 根据 blog_id 查询 blog 详细信息
// export const ShowBlogDetail = (blog_id) => ajax(BASE_URL+'/blog/showBlogDetail',{blog_id},'GET');
export const ShowBlogDetail = (blog_id) => ajax(BASE_URL_2+'/iwork/httpservice/ShowBlogDetail2',{blog_id},'GET');

// 新建课程
// export const NewCourse = (course_name,course_type,course_sub_type,course_short_desc) =>
//   ajax(BASE_URL+'/ilearning/newCourse',{course_name,course_type,course_sub_type,course_short_desc},'GET');
export const NewCourse = (course_name,course_type,course_sub_type,course_short_desc) =>
  ajax(BASE_URL_2+"/iwork/httpservice/NewCourse2",{course_name,course_type,course_sub_type,course_short_desc},'GET');

// 分页查询我的课程清单
// export const GetMyCourseList = (userName) => ajax(BASE_URL+'/ilearning/getMyCourseList',{userName},'GET');
export const GetMyCourseList = (userName) => ajax(BASE_URL_2+"/iwork/httpservice/GetMyCourseList2",{userName},'GET');

// 完结视频更新
// export const EndUpdate = (course_id) => ajax(BASE_URL+'/ilearning/endUpdate',{course_id},'GET');
export const EndUpdate = (course_id) => ajax(BASE_URL_2+"/iwork/httpservice/EndUpdate2",{course_id},'GET');

// 显示课程详细信息
// export const ShowCourseDetail = (course_id) => ajax(BASE_URL+'/ilearning/showCourseDetail',{course_id},'GET');
export const ShowCourseDetail = (course_id) => ajax(BASE_URL_2+"/iwork/httpservice/ShowCourseDetail2",{course_id},'GET');

// 切换收藏点赞
// export const ToggleFavorite = (favorite_id, favorite_type) => ajax(BASE_URL+'/ilearning/toggle_favorite',{favorite_id, favorite_type},'GET');
export const ToggleFavorite = (favorite_id, favorite_type) => ajax(BASE_URL_2+"/iwork/httpservice/ToggleFavorite2",{favorite_id, favorite_type},'GET');

// 查询评论主题
// export const FilterCommentTheme = (comment_id, theme_type) => ajax(BASE_URL+'/ilearning/filterCommentTheme',{comment_id, theme_type},'GET');
export const FilterCommentTheme = (comment_id, theme_type) => ajax(BASE_URL_2+'/iwork/httpservice/FilterCommentTheme2',{comment_id, theme_type},'GET');

// 添加评论
// export const AddCommentReply = (parent_id, reply_content, comment_id, theme_type, reply_comment_type, refer_user_name) =>
//   ajax(BASE_URL+'/ilearning/addCommentReply',{parent_id, reply_content, comment_id, theme_type, reply_comment_type, refer_user_name},'GET');
export const AddCommentReply = (parent_id, reply_content, comment_id, theme_type, reply_comment_type, refer_user_name) =>
  ajax(BASE_URL_2+'/iwork/httpservice/AddCommentReply2',{parent_id, reply_content, comment_id, theme_type, reply_comment_type, refer_user_name},'GET');

// 获取评论列表
// export const FilterCommentReply = (comment_id, theme_type, parent_id, reply_comment_type) =>
//   ajax(BASE_URL+'/ilearning/filterCommentReply',{comment_id, theme_type, parent_id, reply_comment_type},'GET');
export const FilterCommentReply = (comment_id, theme_type, parent_id, reply_comment_type) =>
  ajax(BASE_URL_2+'/iwork/httpservice/FilterCommentReply2',{comment_id, theme_type, parent_id, reply_comment_type},'GET');

// 获取所有课程类型
// export const GetAllCourseType = () => ajax(BASE_URL+'/ilearning/getAllCourseType',{},'GET');
export const GetAllCourseType = () => ajax(BASE_URL_2+"/iwork/httpservice/GetAllCourseType2",{},'GET');

// export const ShowCourseHistory = (offset,current_page) => ajax(BASE_URL+"/common/showCourseHistory", {offset,current_page},'GET');
export const ShowCourseHistory = (offset,current_page) => ajax(BASE_URL_2+"/iwork/httpservice/ShowCourseHistory2", {offset,current_page},'GET');

// 获取热门推荐的课程
// export const GetHotCourseRecommend = () => ajax(BASE_URL+'/ilearning/getHotCourseRecommend',{},'GET');
export const GetHotCourseRecommend = () => ajax(BASE_URL_2+"/iwork/httpservice/GetHotCourseRecommend2",{},'GET');

// 根据课程名称获取所有子类型名称
// export const GetAllCourseSubType = (course_type) => ajax(BASE_URL+'/ilearning/getAllCourseSubType',{course_type},'GET');
export const GetAllCourseSubType = (course_type) => ajax(BASE_URL_2+'/iwork/httpservice/GetAllCourseSubType2',{course_type},'GET');

// 课程搜索
// export const SearchCourseList = (search) => ajax(BASE_URL+'/ilearning/searchCourseList',{search},'GET');
export const SearchCourseList = (search) => ajax(BASE_URL_2+'/iwork/httpservice/SearchCourseList2',{search},'GET');

// 添加配置项
// export const AddConfiguration = (parent_id, configuration_name, configuration_value) =>
//   ajax(BASE_URL+'/cms/addConfiguration',{parent_id, configuration_name, configuration_value},'GET');
export const AddConfiguration = (parent_id, configuration_name, configuration_value) =>
  ajax(BASE_URL_2+'/iwork/httpservice/AddConfiguration2',{parent_id, configuration_name, configuration_value},'GET');

// 根据名称查询配置项
// export const QueryAllConfigurations = (configuration_name) => ajax(BASE_URL+'/cms/queryAllConfigurations',{configuration_name},'GET');

// 分页查询配置项信息
// export const FilterConfigurations = (search,offset,current_page) => ajax(BASE_URL+'/cms/filterConfigurations',{search, offset,current_page},'GET');
export const FilterConfigurations = (search,offset,current_page) => ajax(BASE_URL_2+'/iwork/httpservice/FilterConfigurations2',{search, offset,current_page},'GET');

// 获取随机数量的友情链接地址
// export const QueryRandomCommonLink = (link_type) => ajax(BASE_URL+'/cms/queryRandomCommonLink',{link_type},'GET');
export const QueryRandomCommonLink = (link_type) => ajax(BASE_URL_2+'/iwork/httpservice/QueryRandomCommonLink2',{link_type},'GET');

// 分页查询友情链接地址
// export const FilterCommonLinks = (offset,current_page,search) => ajax(BASE_URL+'/cms/filterCommonLinks',{offset,current_page,search},'GET');
export const FilterCommonLinks = (offset,current_page,search) => ajax(BASE_URL_2+'/iwork/httpservice/FilterCommonLinks2',{offset,current_page,search},'GET');

// 添加友情链接地址
// export const AddCommonLink = (link_type, link_name, link_addr) => ajax(BASE_URL+'/cms/addCommonLink',{link_type, link_name, link_addr},'GET');
export const AddCommonLink = (link_type, link_name, link_addr) => ajax(BASE_URL_2+'/iwork/httpservice/AddCommonLink2',{link_type, link_name, link_addr},'GET');

// export const FilterElements = (offset,current_page,search) => ajax(BASE_URL+'/cms/filterElements',{offset,current_page,search},'GET');
export const FilterElements = (offset,current_page,search) => ajax(BASE_URL_2+'/iwork/httpservice/FilterElements2',{offset,current_page,search},'GET');

// export const AddElement = (placement, title, content, imgpath, linked_refer) => ajax(BASE_URL+'/cms/addElement',{placement, title, content, imgpath, linked_refer},'GET');
export const AddElement = (placement, title, content, imgpath, linked_refer) => ajax(BASE_URL_2+'/iwork/httpservice/AddElement2',{placement, title, content, imgpath, linked_refer},'GET');

export const UpdateElementStatus = (id, status) => ajax(BASE_URL+'/cms/updateElementStatus',{id, status},'GET');

// export const AddPlacement = (placement_name, placement_desc) => ajax(BASE_URL+'/cms/addPlacement',{placement_name, placement_desc},'GET');
export const AddPlacement = (placement_name, placement_desc) => ajax(BASE_URL_2+'/iwork/httpservice/AddPlacement2',{placement_name, placement_desc},'GET');

// export const FilterPlacement = (offset,current_page,search) => ajax(BASE_URL+'/cms/filterPlacement',{offset,current_page,search},'GET');
export const FilterPlacement = (offset,current_page,search) => ajax(BASE_URL_2+'/iwork/httpservice/FilterPlacement2',{offset,current_page,search},'GET');

// export const DeletePlacementById = (id) => ajax(BASE_URL+'/cms/deletePlacementById',{id},'GET');
export const DeletePlacementById = (id) => ajax(BASE_URL_2+'/iwork/httpservice/DeletePlacementById2',{id},'GET');

// export const FilterElementByPlacement = (placement) => ajax(BASE_URL+'/cms/filterElementByPlacement',{placement},'GET');
export const FilterElementByPlacement = (placement) => ajax(BASE_URL_2+'/iwork/httpservice/FilterElementByPlacement2',{placement},'GET');

// 根据 blog_id 查询 blog 详细信息
// export const ShowShareDetail = (share_id) => ajax(BASE_URL+'/share/showShareDetail',{share_id},'GET');
export const ShowShareDetail = (share_id) => ajax(BASE_URL_2+'/iwork/httpservice/ShowShareDetail2',{share_id},'GET');

// 获取ShareList 信息
// export const FilterShareList = (offset,current_page,search_type) => ajax(BASE_URL+'/share/filterShareList',{offset,current_page,search_type},'GET');
export const FilterShareList = (offset,current_page,search_type) => ajax(BASE_URL_2+'/iwork/httpservice/FilterShareList2',{offset,current_page,search_type},'GET');

// 新增共享链接
// export const AddNewShare = (share_type,share_desc,link_href,content) => ajax(BASE_URL+'/share/addNewShare',{share_type,share_desc,link_href,content},'GET');
export const AddNewShare = (share_type,share_desc,link_href,content) => ajax(BASE_URL_2+'/iwork/httpservice/AddNewShare2',{share_type,share_desc,link_href,content},'GET');

// 登录接口
// export const Login = (username,passwd) => ajax(BASE_URL+"/sso/user/login", {username,passwd},'POST');
export const Login = (username,passwd) => ajax(BASE_URL_2+"/iwork/httpservice/PostLogin2", {username,passwd},'POST');

// 注册接口
// export const Regist = (username,passwd) => ajax(BASE_URL+"/sso/user/regist", {username,passwd},'POST');
export const Regist = (username,passwd) => ajax(BASE_URL_2+"/iwork/httpservice/Regist2", {username,passwd},'POST');

export const GetHotUsers = () => ajax(BASE_URL_2+"/iwork/httpservice/GetHotUsers", {},'POST');

// 系统注册分页查询
// export const AppRegisterList = (offset,current_page,search) => ajax(BASE_URL+"/sso/app/appRegisterList", {offset,current_page,search},'POST');
export const AppRegisterList = (offset,current_page,search) => ajax(BASE_URL_2+"/iwork/httpservice/AppRegisterList2", {offset,current_page,search},'POST');

// 添加系统注册
// export const AddAppRegister = (app_address) => ajax(BASE_URL+"/sso/app/addAppRegister", {app_address},'POST');
export const AddAppRegister = (app_address) => ajax(BASE_URL_2+"/iwork/httpservice/AddAppRegister2", {app_address},'POST');

// 登录记录分页查询
// export const LoginRecordList = (offset,current_page,search) => ajax(BASE_URL+"/sso/user/loginRecordList", {offset,current_page,search},'POST');
export const LoginRecordList = (offset,current_page,search) => ajax(BASE_URL_2+"/iwork/httpservice/LoginRecordList2", {offset,current_page,search},'POST');


// 跨模块使用,模块化部署时需要使用 nginx 代理
export const LoginAddr = "/sso/login/";
