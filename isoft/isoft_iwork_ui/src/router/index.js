import Vue from 'vue'
import Router from 'vue-router'

import WorkList from "../components/IWork/IWork/WorkList"
import WorkStepList from "../components/IWork/IWorkStep/WorkStepList"
import RunLogList from "../components/IWork/IRunLog/RunLogList"
import WorkHistoryList from "../components/IWork/IHistory/WorkHistoryList"
import RunLogDetail from "../components/IWork/IRunLog/RunLogDetail"
import IWorkLayout from "../components/ILayout/IWorkLayout"
import QuartzList from "../components/IWork/IQuartz/QuartzList"
import ResourceList from "../components/IWork/IResource/ResourceList"
import MigrateList from "../components/IWork/IMigrate/MigrateList"
import EditMigrate from "../components/IWork/IMigrate/EditMigrate"
import QuickSql from "../components/IWork/IAssistant/QuickSql"
import Format from "../components/IWork/IAssistant/Format"
import Template from "../components/IWork/IAssistant/Template"
import GlobalVarList from "../components/IWork/IGlobalVar/GlobalVarList"
import EntityList from "../components/IWork/IEntity/EntityList"
import File from "../components/IWork/IFile/File"

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
      {path: 'quartzList',component: QuartzList},
      {path: 'resourceList',component: ResourceList},
      {path: 'workList',component: WorkList},
      {path: 'workstepList',component: WorkStepList},
      {path: 'runLogList',component: RunLogList},
      {path: 'workHistoryList',component: WorkHistoryList},
      {path: 'runLogDetail',component: RunLogDetail},
      {path: 'migrateList',component: MigrateList},
      {path: 'editMigrate',component: EditMigrate},
      {path: 'quickSql',component: QuickSql},
      {path: 'format',component: Format},
      {path: 'template',component: Template},
      {path: 'globalVarList',component: GlobalVarList},
      {path: 'entityList',component: EntityList},
      {path: 'files',component: File},
    ]
  }
];


export default new Router({
  // # 主要用来区分前后台应用, history 模式需要使用 nginx 代理
  // History 模式,去除vue项目中的 #
  // mode: 'history',
  routes: IWorkRouter,
})
