import {oneOf} from "../tools"
import {checkSSOLogin} from "./sso"

// 前台开启的模块 sso,ilearning,iwork
const open_modules = ["ilearning", "sso"];

export function modulesCheck(moduleName) {
  return oneOf(moduleName, open_modules);
}

export const CheckSSOLogin = (to, from, next) => checkSSOLogin(to, from, next);

