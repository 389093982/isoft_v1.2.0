import Login from "../components/SSO/Login/Login"
import Regist from "../components/SSO/Login/Regist"
import AppRegist from "../components/SSO/AppRegist"
import LoginRecord from "../components/SSO/LoginRecord"
import IEmptyLayout from "../components/ILayout/IEmptyLayout"

export const ISSOReouter = [{
  path: '/sso',
  component: IEmptyLayout,
  children: [
    {path: 'login',component: Login},
    {path: 'regist',component: Regist},
    {path: 'appRegist',component: AppRegist},
    {path: 'loginRecord',component: LoginRecord},
  ]
}];


