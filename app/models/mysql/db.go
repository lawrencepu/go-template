package mysql

import (
	"go-template/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	DB = InitDB(dbConfig.MaxIdleConn, dbConfig.MaxOpenConn, dbConfig.ConnMaxLifetime, dbConfig.Dsn)
}



// init db
func InitDB(maxIdleConn, maxOpenConn, connMaxLifetime int, dsn string) *gorm.DB {
	return mysqlConn(maxIdleConn, maxOpenConn, connMaxLifetime, dsn)
}

// init mysql pool
func mysqlConn(maxIdleConn, maxOpenConn, connMaxLifetime int, dsn string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(dsn))
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