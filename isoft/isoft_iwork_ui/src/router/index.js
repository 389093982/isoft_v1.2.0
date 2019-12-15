import Vue from 'vue'
import Router from 'vue-router'

// es6 import 异步语法,使用异步组件加载机制减少耗时操作
const WorkList = () => import("@/components/IWork/IWork/WorkList");
const WorkStepList = () => import("@/components/IWork/IWorkStep/WorkStepList");
const RunLogList = () => import("@/components/IWork/IRunLog/RunLogList");
const WorkHistoryList = () => import("@/components/IWork/IHistory/WorkHistoryList");
const RunLogDetail = () => import("@/components/IWork/IRunLog/RunLogDetail");
const IWorkLayout = () => import("@/components/ILayout/IWorkLayout");
const QuartzList = () => import("@/components/IWork/IQuartz/QuartzList");
const ResourceList = () => import("@/components/IWork/IResource/ResourceList");
const MigrateList = () => import("@/components/IWork/IMigrate/MigrateList");
const EditMigrate = () => import("@/components/IWork/IMigrate/EditMigrate");
const QuickSql = () => import("@/components/IWork/IAssistant/QuickSql");
const Format = () => import("@/components/IWork/IAssistant/Format");
const GlobalVarList = () => import("@/components/IWork/IGlobalVar/GlobalVarList");
const File = () => import("@/components/IWork/IFile/File");
const DashBoard = () => import("@/components/IWork/IDashBoard/DashBoard");
const IModuleList = () => import("@/components/IWork/IModule/IModuleList");
const IFilterList = () => import("@/components/IWork/IFilter/IFilterList");
import AuditTaskList from "../components/IWork/Audit/AuditTaskList"
import AuditDetail from "../components/IWork/Audit/AuditDetail"

const Element = () => import("@/components/IWork/IPlacement/Element");
const EditElement = () => import("@/components/IWork/IPlacement/EditElement");
const Placement = () => import("@/components/IWork/IPlacement/Placement");
const EditPlacement = () => import("@/components/IWork/IPlacement/EditPlacement");

Vue.use(Router);

const IWorkRouter = [
  {
    path: '/',
    redirect: '/iwork/workList'
  },
  {
    path: '/iwork',
    component: IWorkLayout,
    children: [
      {path: 'moduleList',component: IModuleList},
      {path: 'quartzList',component: QuartzList},
      {path: 'resourceList',component: ResourceList},
      {path: 'workList',component: WorkList},
      {path: 'filterList',component: IFilterList},
      {path: 'workstepList',component: WorkStepList},
      {path: 'runLogList',component: RunLogList},
      {path: 'workHistoryList',component: WorkHistoryList},
      {path: 'runLogDetail',component: RunLogDetail},
      {path: 'migrateList',component: MigrateList},
      {path: 'editMigrate',component: EditMigrate},
      {path: 'quickSql',component: QuickSql},
      {path: 'format',component: Format},
      {path: 'globalVarList',component: GlobalVarList},
      {path: 'files',component: File},
      {path: 'dashboard',component: DashBoard},
      {path: 'audit',component: AuditTaskList},
      {path: 'audit_detail',component: AuditDetail},
      {path: 'elementEdit',component: EditElement},
      {path: 'elementList',component: Element},
      {path: 'placementList',component: Placement},
      {path: 'placementEdit',component: EditPlacement},
    ]
  }
];


export default new Router({
  // # 主要用来区分前后台应用, history 模式需要使用 nginx 代理
  // History 模式,去除vue项目中的 #
  // mode: 'history',
  routes: IWorkRouter,
  // 页面跳转时,让页面滚动在顶部
  scrollBehavior(to,from,savedPosition){
    if(savedPosition){
      return savedPosition;
    }else{
      return {x:0,y:0}
    }
  },
})
