package mysql

import (
	"go-template/config"
	"go-template/tools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// DB is sql *db
var DB *gorm.DB

const dbKey = "mysql"

// 初始化数据库组件
func Init() {
	dbConfig, err := config.GetDBConf(dbKey)
	if err != nil {
		panic(err)
	}
	DB = InitDB(dbConfig.MaxIdleConn, dbConfig.MaxOpenConn, dbConfig.ConnMaxLifetime, dbConfig.Dsn, dbConfig.Prefix)
}

// init db
func InitDB(maxIdleConn, maxOpenConn, connMaxLifetime int, dsn string, prefix string) *gorm.DB {
	return mysqlConn(maxIdleConn, maxOpenConn, connMaxLifetime, dsn, prefix)
}

// init mysql pool
func mysqlConn(maxIdleConn, maxOpenConn, connMaxLifetime int, dsn string, prefix string) *gorm.DB {
	l := logger.New(tools.Logger, logger.Config{
		//慢SQL阈值
		SlowThreshold: time.Second,
		IgnoreRecordNotFoundError: true,
		Colorful: false,
		//设置日志级别，只有Warn以上才会打印sql
		LogLevel: logger.Info,
	})
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: true,
		},
		Logger: l,
	})

	if err != nil {
		log.Printf("[db] mysql fail, err=%s", err)
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("[db] mysql fail, err=%s", err)
		panic(err)
	}
	DB.Callback().Create().Before("gorm:create").Register("created_at", UpdateTimeForCreateCallback)
	DB.Callback().Update().Before("gorm:update").Register("update_at", UpdateTimeForUpdatedCallback)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(connMaxLifetime))

	err = sqlDB.Ping()
	if err != nil {
		log.Printf("[db] mysql fail, err:%s", err.Error())
		panic(err)
	}
	log.Printf("[db] mysql success")
	return DB
}

//创建时自动维护字段
func UpdateTimeForCreateCallback(db *gorm.DB) {
	db.Statement.SetColumn("CreatedAt", time.Now())
	db.Statement.SetColumn("UpdatedAt", time.Now())
}

//更新时设置修改时间
func UpdateTimeForUpdatedCallback(db *gorm.DB) {
	db.Statement.SetColumn("UpdatedAt", time.Now())
}
