package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-template/app/http/middleware"
	"go-template/app/models/mysql"
	"go-template/app/models/redis"
	"go-template/config"
	"go-template/routes"
	"go-template/tools"
	"time"
)

type App struct {
	WebServer *gin.Engine
	BaseConfig *config.BaseConf
}


func InitApp() *App {
	app := new(App)
	app.RegisterConfig("./config/app.json")
	app.Timezone()
	app.InitLogger()
	app.RegisterWebServer()
	app.RegisterDB()
	app.RegisterMiddleware()
	app.RegisterRouter()
	return app
}

//注册路由
func (app *App) RegisterRouter()  {
	func(rr func(c *gin.Engine)) {
		rr(app.WebServer)
	}(routes.Register)
}

//注册基础中间件
func (app *App) RegisterMiddleware()  {
	app.WebServer.Use(
		middleware.RequestTrace,
		middleware.SetLog,
	)
}

// 全局时区设置
func (app *App) Timezone()  {
	loc, _ := time.LoadLocation(app.BaseConfig.Timezone)
	time.Local = loc
}

// 注册配置文件
func (app *App) RegisterConfig(cfg string) {
	err := config.InitConfig(cfg)
	if err != nil {
		panic(err)
	}
	baseConfig, err := config.GetBaseConf()
	if err != nil {
		panic(err)
	}
	app.BaseConfig = baseConfig
}

// 初始化微博服务提供
func (app *App) RegisterWebServer() {
	app.WebServer = gin.Default()
}

// 开启web服务
func (app *App) Run() {
	app.WebServer.Run(app.BaseConfig.HttpListen)
}

// sql/nosql初始化
func (app *App) RegisterDB() {
	mysql.Init()
	redis.Init()
}

// 日志格式初始化
func (app *App) InitLogger() {
	tools.InitLogger()
}