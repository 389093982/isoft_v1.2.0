import {getCookie} from "./index"
import {checkContainsInString} from "./index"
import {checkEmpty} from "./index"

// sso 登陆拦截
const _checkSSOLogin = function(to, from, next) {
  // 登录判断
  var userName = getCookie("userName");
  var isLogin = getCookie("isLogin");
  var token = getCookie("tokenString");
  // 非免登录白名单,并且不含登录标识的需要重新跳往登录页面
  if(!_checkNotLogin() && (checkEmpty(userName) || checkEmpty(isLogin) || checkEmpty(token) || isLogin != "isLogin")){
    // 跳往登录页面
    window.location.href = "/sso/login/?redirectUrl=" + window.location.href;
  }else{
    next();
  }
};

const _checkNotLogin = function (){
  if(checkContainsInString(window.location.href, "/sso/login") || checkContainsInString(window.location.href, "/sso/regist")){
    return true;
  }
  return false;
}

export const checkSSOLogin = (to, from, next) => _checkSSOLogin(to, from, next);
export const checkNotLogin = () => _checkNotLogin();

