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

// 定时任务分页查询
export const QuartzList = (offset,current_page,search) => ajax(BASE_URL+"/iwork/filterPageQuartz", {offset,current_page,search},'POST');

// 编辑 quartz
export const EditQuartz = (task_name, operate) => ajax(BASE_URL+"/iwork/editQuartz", {task_name, operate},'POST');

// 添加 quartz 记录
export const AddQuartz = (task_name,task_type,cron_str) => ajax(BASE_URL+"/iwork/addQuartz", {task_name,task_type,cron_str},'POST');

// resource 分页查询
export const ResourceList = (offset,current_page,search) => ajax(BASE_URL+"/iwork/filterPageResource", {offset,current_page,search},'POST');

// 添加 resource 记录
export const AddResource = (resource_name,resource_type,resource_url,resource_dsn,resource_username,resource_password) => ajax(BASE_URL+"/iwork/addResource", {resource_name,resource_type,resource_url,resource_dsn,resource_username,resource_password},'POST');

// 删除 resource 记录
export const DeleteResource = (id) => ajax(BASE_URL+"/iwork/deleteResource", {id},'POST');

// 验证 resource
export const ValidateResource = (id) => ajax(BASE_URL+"/iwork/validateResource", {id},'POST');

export const GetAllResource = (resource_type) => ajax(BASE_URL+"/iwork/getAllResource", {resource_type},'POST');

// work 分页查询
export const WorkList = (offset,current_page,search) => ajax(BASE_URL+"/iwork/filterPageWork", {offset,current_page,search},'POST');

// 编辑 work 记录
export const EditWork = (work_id, work_name, work_desc) => ajax(BASE_URL+"/iwork/editWork", {work_id, work_name, work_desc},'POST');

// 删除 work 记录
export const DeleteWorkById = (id) => ajax(BASE_URL+"/iwork/deleteWorkById", {id},'POST');

// workstep 分页查询
export const WorkStepList = (work_id) => ajax(BASE_URL+"/iwork/filterWorkStep", {work_id},'POST');

// 添加 workstep 记录
export const AddWorkStep = (work_id, work_step_id) => ajax(BASE_URL+"/iwork/addWorkStep", {work_id, work_step_id},'POST');

// 运行 work
export const RunWork = (work_id) => ajax(BASE_URL+"/iwork/runWork", {work_id},'POST');

// 保存为历史版本
export const SaveHistory = (work_id) => ajax(BASE_URL+"/iwork/saveHistory", {work_id},'POST');

// 分页查询历史版本信息
export const FilterPageWorkHistory = (offset,current_page) => ajax(BASE_URL+"/iwork/filterPageWorkHistory", {offset,current_page},'POST');

// 编辑 workstep 基本信息
export const EditWorkStepBaseInfo = (work_id,work_step_id,work_step_name,work_step_desc,work_step_type, is_defer) => ajax(BASE_URL+"/iwork/editWorkStepBaseInfo", {work_id,work_step_id,work_step_name,work_step_desc,work_step_type, is_defer},'POST');

// 编辑 workstep 记录
export const EditWorkStepParamInfo = (work_id,work_step_id,paramInputSchemaStr, paramMappingsStr) => ajax(BASE_URL+"/iwork/editWorkStepParamInfo", {work_id,work_step_id,paramInputSchemaStr,paramMappingsStr},'POST');

// 加载 workstep 记录
export const LoadWorkStepInfo = (work_id,work_step_id) => ajax(BASE_URL+"/iwork/loadWorkStepInfo", {work_id,work_step_id},'POST');

// 加载前置节点输出参数,包括全局参数
export const LoadPreNodeOutput = (work_id,work_step_id) => ajax(BASE_URL+"/iwork/loadPreNodeOutput", {work_id,work_step_id},'POST');

// 删除 workstep 记录
export const DeleteWorkStepByWorkStepId = (work_id, work_step_id) => ajax(BASE_URL+"/iwork/deleteWorkStepByWorkStepId", {work_id, work_step_id},'POST');

// 交换 workstep 顺序
export const ChangeWorkStepOrder = (work_id,work_step_id,type) => ajax(BASE_URL+"/iwork/changeWorkStepOrder", {work_id,work_step_id,type},'POST');

// 分页查询运行记录
export const FilterPageLogRecord = (work_id,offset,current_page) => ajax(BASE_URL+"/iwork/filterPageLogRecord", {work_id,offset,current_page},'POST');

// 获取最后一次运行日志详情
export const GetLastRunLogDetail = (tracking_id) => ajax(BASE_URL+"/iwork/getLastRunLogDetail", {tracking_id},'POST');

// 获取相关流程
export const GetRelativeWork = (work_id) => ajax(BASE_URL+"/iwork/getRelativeWork", {work_id},'POST');

// 分页查询 entity 信息
export const FilterPageEntity = (offset,current_page) => ajax(BASE_URL+"/iwork/filterPageEntity", {offset,current_page},'POST');

// 编辑 entity
export const EditEntity = (entity_id,entity_name,entity_field_str) => ajax(BASE_URL+"/iwork/editEntity", {entity_id,entity_name,entity_field_str},'POST');

// 删除 entity
export const DeleteEntity = (entity_id) => ajax(BASE_URL+"/iwork/deleteEntity", {entity_id},'POST');

// 校验整个工程
export const ValidateAllWork = () => ajax(BASE_URL+"/iwork/validateAllWork", {},'POST');

// 显示校验结果
export const LoadValidateResult = () => ajax(BASE_URL+"/iwork/loadValidateResult", {},'POST');

export const RefactorWorkStepInfo = (work_id, refactor_worksub_name,selections) => ajax(BASE_URL+"/iwork/refactorWorkStepInfo", {work_id, refactor_worksub_name,selections},'POST');

export const BatchChangeIndent = (work_id, mod,selections) => ajax(BASE_URL+"/iwork/batchChangeIndent", {work_id, mod,selections},'POST');

export const SubmitMigrate = (tableName, table_migrate_sql, tableColunms, id, operateType) => ajax(BASE_URL+"/iwork/submitMigrate", {tableName, table_migrate_sql, tableColunms, id, operateType},'POST');

export const FilterPageMigrate = (filterTableName, offset,current_page) => ajax(BASE_URL+"/iwork/filterPageMigrate", {filterTableName, offset,current_page},'POST');

export const GetMigrateInfo = (id) => ajax(BASE_URL+"/iwork/getMigrateInfo", {id},'POST');

export const ExecuteMigrate = (resource_name, forceClean) => ajax(BASE_URL+"/iwork/executeMigrate", {resource_name, forceClean},'POST');

export const BuildInstanceSql = (tableName, tableColunms, id, operateType) => ajax(BASE_URL+"/iwork/buildInstanceSql", {tableName, tableColunms, id, operateType},'POST');

export const LoadQuickSqlMeta = (resource_id) => ajax(BASE_URL+"/iwork/loadQuickSqlMeta", {resource_id},'POST');

export const ParseToMultiValue = (pureText, value) => ajax(BASE_URL+"/iwork/parseToMultiValue", {pureText, value},'POST');

// 跨模块使用,模块化部署时需要使用 nginx 代理
export const LoginAddr = "/sso/login/";
