package sysconfig

import "github.com/astaxie/beego"

var SYSCONFIG_VARS_USAGE_LOGGABLE = beego.AppConfig.DefaultBool("sysconfig_vars_usage_loggable", false)
