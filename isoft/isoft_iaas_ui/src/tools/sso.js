import {getCookie} from "./index"
import {checkContainsInString} from "./index"
import {checkEmpty} from "./index"

const _checkHasLogin = function(){
  let userName = getCookie("userName");
  var isLogin = getCookie("isLogin");
  var token = getCookie("tokenString");
  return !checkEmpty(userName) && !checkEmpty(isLogin) && !checkEmpty(token) && isLogin == "isLogin";
};

const _getLoginUserName = function(){
  return getCookie("userName");
};

// sso 登陆拦截
const _checkSSOLogin = function(to, from, next) {
  // 非免登录白名单,并且不含登录标识的需要重新跳往登录页面
  if(_mustLogin(to.path) && !_checkHasLogin()){
    // 跳往登录页面
    window.location.href = "/#/sso/login/?redirectUrl=" + window.location.href;
  }else{
    next();
  }
};

const _mustLogin = function(target){
  // 包含 /mine/ 是必须要检查登录状态的,其它地址是免登陆的
  return checkContainsInString(target, "/mine/");
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
export const getLoginUserName = () => _getLoginUserName();

