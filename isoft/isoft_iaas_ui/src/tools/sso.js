import {getCookie} from "./index"
import {checkContainsInString} from "./index"
import {checkEmpty} from "./index"

const _checkAdminLogin = function () {
  let roleName = getCookie("roleName");
  return roleName == "admin";

}

const _checkHasLogin = function(){
  let userName = getCookie("userName");
  var isLogin = getCookie("isLogin");
  var token = getCookie("tokenString");
  var expireSecond = getCookie("expireSecond");
  if(expireSecond == null || new Date().getTime() > expireSecond){    // 判断是否过期
    return false;
  }
  return !checkEmpty(userName) && !checkEmpty(isLogin) && !checkEmpty(token) && isLogin == "isLogin";
};

const _getLoginUserName = function(){
  return getCookie("userName");
};

// sso 登陆拦截
const _checkSSOLogin = function(to, from, next) {
  // 1、必须要登录但是没登录
  // 2、必须要 admin 登录但是没登录或者不是 admin 登录
  if((_mustLogin(to.path) && !_checkHasLogin()) || (_mustAdminLogin(to.path)  && (!_checkHasLogin() || !_checkAdminLogin()))){
    if(!checkContainsInString(from.path, "/sso/login/")){
      // 跳往登录页面
      window.location.href = "/#/sso/login/?redirectUrl=" + window.location.href;
    }
  }else{
    next();
  }
};

const _mustLogin = function(target){
  // 包含 /mine/ 是必须要检查登录状态的,其它地址是免登陆的
  return checkContainsInString(target, "/mine/");
}

const _mustAdminLogin = function(target){
  // 包含 /background/ 必须 admin 登录
  return checkContainsInString(target, "/background/");
}

const _checkNotLogin = function (){
  if(checkContainsInString(window.location.href, "/#/sso/login") || checkContainsInString(window.location.href, "/#/sso/regist")){
    return true;
  }
  return false;
}

export const checkSSOLogin = (to, from, next) => _checkSSOLogin(to, from, next);
export const checkNotLogin = () => _checkNotLogin();
export const checkHasLogin = () => _checkHasLogin();
export const checkAdminLogin = () => _checkAdminLogin();
export const getLoginUserName = () => _getLoginUserName();

