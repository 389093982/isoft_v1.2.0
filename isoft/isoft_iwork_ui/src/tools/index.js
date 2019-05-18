import cronValidate from "./cron"

// 获取 cookie 值
export const getCookie = function getCookie(c_name) {
  if (document.cookie.length > 0) {
    var c_start = document.cookie.indexOf(c_name + "=");
    if (c_start != -1) {
      c_start = c_start + c_name.length+1;
      var c_end=document.cookie.indexOf(";", c_start);
      if (c_end == -1) {
        c_end = document.cookie.length;
      }
      return unescape(document.cookie.substring(c_start,c_end));
    }
  }
  return "";
};

//删除cookie
export const delCookie = function delCookie(name) {
  document.cookie = name+"=;expires="+(new Date(0)).toGMTString();
};

export const checkEmpty = function checkEmpty(checkStr){
  if(checkStr == null || checkStr == undefined || checkStr == ""){
    return true;
  }
  return false;
};

// 跨域设置 cookie
export function setCookie (c_name,value,expiredays,domain){
  if(checkContainsInString(domain, "localhost")){
    // instead for localhost you should use false
    domain = false
  }
  var exdate = new Date();
  exdate.setDate(exdate.getDate() + expiredays);
  //判断是否需要跨域存储
  if (domain) {
    // egg：path=/;domain=xueersi.com";
    document.cookie = c_name+ "=" +escape(value)+((expiredays==null) ? "" : ";expires="+exdate.toGMTString())+";path=/;domain=" + domain;
  } else {
    document.cookie = c_name+ "=" +escape(value)+((expiredays==null) ? "" : ";expires="+exdate.toGMTString())+";path=/";
  }
}

// 判断值 value 是否是列表 validList 中
export function oneOf (value, validList) {
  for (let i = 0; i < validList.length; i++) {
    if (value === validList[i]) {
      return true;
    }
  }
  return false;
}

// 判断字符串是否包含子串
export function checkContainsInString(str, subStr) {
  return str.indexOf(subStr) != -1
}

// 根据正则验证字符串
export function validatePatternForString(pattern, str) {
  return pattern.test(str);
}

// 校验只能含有字母数字和下划线
export function validateCommonPatternForString(str) {
  var uPattern = /^[a-zA-Z0-9_]{1,}$/;
  return uPattern.test(str);
}

// 校验密码
export function validateUserName(username) {
  // 6至20位，以字母开头，字母，数字，减号，下划线!
  var uPattern = /^[a-zA-Z]([-_a-zA-Z0-9]{5,19})+$/;
  return uPattern.test(username);
}

// 校验用户名
export function validatePasswd(passwd) {
  // 最少6位，至少1个大小写字母，数字和特殊字符!
  var pPattern = /^.*(?=.{6,})(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*? ]).*$/;
  return pPattern.test(passwd);
}

// 时间格式化
export function formatDate (date, fmt) {
  let o = {
    'M+': date.getMonth() + 1, // 月份
    'd+': date.getDate(), // 日
    'h+': date.getHours(), // 小时
    'm+': date.getMinutes(), // 分
    's+': date.getSeconds(), // 秒
    'S': date.getMilliseconds() // 毫秒
  }
  if (/(y+)/.test(fmt)) {
    fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length));
  }
  for (var k in o) {
    if (new RegExp('(' + k + ')').test(fmt)) {
      fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length)));
    }
  }
  return fmt;
}

export const validateCron = (cron) => cronValidate(cron);

// 字符串分割函数
export function strSplit (str, sep) {
  return str.split(sep);
}

// 字符串重复 n 次
export function getRepeatStr(str, n){
  return new Array(n+1).join(str);
}

/**
 * 数组元素交换位置
 * @param {array} arr 数组
 * @param {number} index1 添加项目的位置
 * @param {number} index2 删除项目的位置
 * index1和index2分别是两个数组的索引值，即是两个要交换元素位置的索引值，如1，5就是数组中下标为1和5的两个元素交换位置
 */
export function swapArray(arr, index1, index2) {
  arr[index1] = arr.splice(index2, 1, arr[index1])[0];
  return arr;
}
