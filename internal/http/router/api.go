package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"jim/config"
	"jim/internal/http/controller"
	"jim/internal/http/middleware"
	"jim/pkg"
	"sync"
)

func registerApi(r *gin.Engine) {
	index := controller.IndexController{}
	user := controller.UserController{}
	api := r.Group("/api")
	{
		api.GET("/ping", index.Pong)
		api.POST("/user/register", user.Register)
		api.POST("/usr/login", user.Login)
	}
}

func NewGin(conf *config.Config) *gin.Engine {
	var once sync.Once
	var instance *gin.Engine
	once.Do(func() {
		env := conf.App.Env
		switch env {
		case config.EnvLocal:
			gin.SetMode(gin.DebugMode)
			break
		case config.EnvTest:
			gin.SetMode(gin.TestMode)
			break
		case config.EnvStage:
			gin.SetMode(gin.TestMode)
		case config.EnvProduction:
		default:
			gin.SetMode(gin.ReleaseMode)
			break
		}
		instance = gin.Default()
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			if err := v.RegisterValidation("mobile", pkg.MobileValidator); err != nil {
				panic(err)
			}
		}

		instance.Use(cors.Default())
		instance.Use(middleware.Recovery(), middleware.AppSetting(), middleware.AccessLog(), middleware.Jaeger())
		registerApi(instance)
	})
	return instance
}
