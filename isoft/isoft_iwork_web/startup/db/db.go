package db

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // _ 的作用,并不需要把整个包都导入进来,仅仅是是希望它执行init()函数而已
	"isoft/isoft/common/chiperutil"
	"isoft/isoft_iwork_web/models"
	"net/url"
)

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	timezone := beego.AppConfig.String("db.timezone")
	aesChiperKey := beego.AppConfig.String("isoft.aes.cipher.key")
	// 对数据库密码进行解密
	dbport = chiperutil.AesDecryptToStr(dbport, aesChiperKey)
	dbuser = chiperutil.AesDecryptToStr(dbuser, aesChiperKey)
	dbpass = chiperutil.AesDecryptToStr(dbpass, aesChiperKey)

	Dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true&charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)

	if timezone != "" {
		Dsn = Dsn + "&loc=" + url.QueryEscape(timezone)
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", Dsn)
	orm.SetMaxIdleConns("default", 100) // SetMaxIdleConns用于设置闲置的连接数
	orm.SetMaxOpenConns("default", 200) // SetMaxOpenConns用于设置最大打开的连接数,默认值为0表示不限制
	db, _ := orm.GetDB("default")
	db.SetConnMaxLifetime(100)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	registerModel()
	createTable()
}

func registerModel() {
	orm.RegisterModel(new(models.CronMeta))
	orm.RegisterModel(new(models.Resource))
	orm.RegisterModel(new(models.Work))
	orm.RegisterModel(new(models.WorkStep))
	orm.RegisterModel(new(models.RunLogRecord))
	orm.RegisterModel(new(models.RunLogDetail))
	orm.RegisterModel(new(models.ValidateLogRecord))
	orm.RegisterModel(new(models.ValidateLogDetail))
	orm.RegisterModel(new(models.WorkHistory))
	orm.RegisterModel(new(models.SqlMigrate))
	orm.RegisterModel(new(models.SqlMigrateLog))
	orm.RegisterModel(new(models.GlobalVar))
	orm.RegisterModel(new(models.Template))
	orm.RegisterModel(new(models.Module))
	orm.RegisterModel(new(models.Filters))
	orm.RegisterModel(new(models.AuditTask))
	orm.RegisterModel(new(models.Placement))
	orm.RegisterModel(new(models.Element))
}

// 自动建表
func createTable() {
	force := false  // 不强制建数据库
	verbose := true // 打印建表过程
	if err := orm.RunSyncdb("default", force, verbose); err != nil {
		beego.Error(err)
	}
}
