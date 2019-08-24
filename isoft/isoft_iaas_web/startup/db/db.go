package db

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	_ "github.com/mattn/go-sqlite3"
	"isoft/isoft_iaas_web/models/cms"
	"isoft/isoft_iaas_web/models/common"
	"isoft/isoft_iaas_web/models/iblog"
	"isoft/isoft_iaas_web/models/ilearning"
	"isoft/isoft_iaas_web/models/share"
	"isoft/isoft_iaas_web/models/sso"
	"net/url"
)

// 数据库同步模式,支持 FLYWAY 和 AUTO
const RunSyncdbMode = "AUTO"

func InitDb() {

	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	timezone := beego.AppConfig.String("db.timezone")

	Dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true&charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)

	if timezone != "" {
		Dsn = Dsn + "&loc=" + url.QueryEscape(timezone)
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", Dsn)
	orm.SetMaxIdleConns("default", 1000) // SetMaxIdleConns用于设置闲置的连接数
	orm.SetMaxOpenConns("default", 2000) // SetMaxOpenConns用于设置最大打开的连接数,默认值为0表示不限制

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	registerModel()
	createTable()
}

func registerModel() {
	orm.RegisterModel(new(iblog.Catalog))
	orm.RegisterModel(new(iblog.Blog))
	orm.RegisterModel(new(ilearning.Course))
	orm.RegisterModel(new(ilearning.CourseVideo))
	orm.RegisterModel(new(ilearning.Favorite))
	orm.RegisterModel(new(ilearning.CommentTheme))
	orm.RegisterModel(new(ilearning.CommentReply))
	orm.RegisterModel(new(ilearning.Note))
	orm.RegisterModel(new(cms.Configuration))
	orm.RegisterModel(new(cms.CommonLink))
	orm.RegisterModel(new(share.Share))
	orm.RegisterModel(new(common.History))

	orm.RegisterModel(new(sso.User))
	orm.RegisterModel(new(sso.AppRegister))
	orm.RegisterModel(new(sso.LoginRecord))
	orm.RegisterModel(new(sso.UserToken))

	orm.RegisterModel(new(cms.Element))
	orm.RegisterModel(new(cms.Placement))
}

// 自动建表
func createTable() {
	force := false  // 不强制建数据库
	verbose := true // 打印建表过程
	if err := orm.RunSyncdb("default", force, verbose); err != nil {
		beego.Error(err)
	}
}
