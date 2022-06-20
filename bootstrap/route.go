package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//SetupRoute路由初始化
func SetupRoute(router *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleWare(router)

	//注册API路由
	routes.RegisterAPIRoutes(router)

	//配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	//注册中间件
	router.Use(gin.Logger(), gin.Recovery())
}

func setup404Handler(router *gin.Engine) {
	//处理404请求

	//处理404请求
	router.NoRoute(func(ctx *gin.Context) {
		//获取标头信息Accept信息
		acceptString := ctx.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			//如果是HTML的话
			ctx.String(http.StatusNotFound, "页面返回404")
		} else {
			//默认返回JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认url和请求方法是否正确",
			})
		}
	})
}
