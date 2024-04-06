package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/cors"
	validatorFactory "goskeleton/app/http/validator/core/factory"
	"goskeleton/app/utils/gin_release"
	"net/http"
)

func InitApiRouter() *gin.Engine {
	var router *gin.Engine
	if variable.ConfigYml.GetBool("AppDebug") == false {
		router = gin_release.ReleaseRouter()
	} else {
		router = gin.Default()
		pprof.Register(router)
	}
	if variable.ConfigYml.GetInt("HttpServer.TrustProxies.IsOpen") == 1 {
		if err := router.SetTrustedProxies(variable.ConfigYml.GetStringSlice("HttpServer.TrustProxies.ProxyServerList")); err != nil {
			variable.ZapLog.Error(consts.GinSetTrustProxyError, zap.Error(err))
		}
	} else {
		_ = router.SetTrustedProxies(nil)
	}

	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Api 模块接口 hello word！")
	})

	router.Static("/public", "./public")

	vApi := router.Group("/api/")
	{
		ocean := vApi.Group("ad/")
		{
			ocean.GET("monitorOcean", validatorFactory.Create(consts.ValidatorPrefix+"Store"))
		}
	}
	return router
}
