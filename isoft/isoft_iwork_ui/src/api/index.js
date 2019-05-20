/*
包含n个接口请求函数的模块
函数的返回值: promise对象
 */
import ajax from './ajax'

const BASE_URL = '/api'

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
export const AddWorkStep = (work_id, work_step_id, work_step_type) => ajax(BASE_URL+"/iwork/addWorkStep", {work_id, work_step_id, work_step_type},'POST');

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
export const FilterPageEntity = (search, offset,current_page) => ajax(BASE_URL+"/iwork/filterPageEntity", {search, offset,current_page},'POST');

// 编辑 entity
export const EditEntity = (entity_id,entity_name,entity_type) => ajax(BASE_URL+"/iwork/editEntity", {entity_id,entity_name,entity_type},'POST');

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

export const GlobalVarList = (offset,current_page, search) => ajax(BASE_URL+"/iwork/globalVarList", {offset,current_page, search},'POST');

export const EditGlobalVar = (id, globalVarName, globalVarValue) => ajax(BASE_URL+"/iwork/editGlobalVar", {id, globalVarName, globalVarValue},'POST');

export const DeleteGlobalVarById = (id) => ajax(BASE_URL+"/iwork/deleteGlobalVarById", {id},'POST');
