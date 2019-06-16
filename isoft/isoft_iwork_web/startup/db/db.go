package db

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	_ "github.com/mattn/go-sqlite3"
	"isoft/isoft_iwork_web/models"
)

// 数据库同步模式,支持 FLYWAY 和 AUTO
const RunSyncdbMode = "AUTO"

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "iwork.db") // iwork 项目使用

	//dbhost := beego.AppConfig.String("db.host")
	//dbport := beego.AppConfig.String("db.port")
	//dbname := beego.AppConfig.String("db.name")
	//dbuser := beego.AppConfig.String("db.user")
	//dbpass := beego.AppConfig.String("db.pass")
	//timezone := beego.AppConfig.String("db.timezone")
	//
	//Dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true&charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)
	//
	//if timezone != "" {
	//	Dsn = Dsn + "&loc=" + url.QueryEscape(timezone)
	//}
	//
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", Dsn)
	//orm.SetMaxIdleConns("default", 1000) // SetMaxIdleConns用于设置闲置的连接数
	//orm.SetMaxOpenConns("default", 2000) // SetMaxOpenConns用于设置最大打开的连接数,默认值为0表示不限制

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	registerModel()

	//if RunSyncdbMode == "FLYWAY" {
	//	// ilearning 模块
	//	flyway.MigrateToDB(Dsn, "./conf/migrations/migrations.sql")
	//	// sso 模块
	//	flyway.MigrateToDB(Dsn, "./conf/migrations/sso_migrations.sql")
	//} else {
	//	createTable()
	//}
	createTable()
}

func registerModel() {
	orm.RegisterModel(new(models.CronMeta))
	orm.RegisterModel(new(models.Resource))
	orm.RegisterModel(new(models.Work))
	orm.RegisterModel(new(models.WorkStep))
	orm.RegisterModel(new(models.RunLogRecord))
	orm.RegisterModel(new(models.RunLogDetail))
	orm.RegisterModel(new(models.Entity))
	orm.RegisterModel(new(models.ValidateLogRecord))
	orm.RegisterModel(new(models.ValidateLogDetail))
	orm.RegisterModel(new(models.WorkHistory))
	orm.RegisterModel(new(models.TableMigrate))
	orm.RegisterModel(new(models.GlobalVar))
	orm.RegisterModel(new(models.Template))
}

// 自动建表
func createTable() {
	force := false  // 不强制建数据库
	verbose := true // 打印建表过程
	if err := orm.RunSyncdb("default", force, verbose); err != nil {
		beego.Error(err)
	}
	//if err := orm.RunSyncdb("iwork", force, verbose); err != nil{
	//	beego.Error(err)
	//}
}
