package router

import (
	"go-vue/api"
	"go-vue/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	//设置启动模式
	gin.SetMode(config.Config.System.Env)
	router := gin.New()
	// 宕机恢复
	router.Use(gin.Recovery())
	// register注册路由
	register(router)
	return router
}

// 路由接口
func register(router *gin.Engine) {
	//路由接口url
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/success", api.Success)
	router.GET("/api/failed", api.Failed)
	router.POST("/api/sysMenu/add", api.CreateSysMenu)
	router.GET("api/sysMenu/list", api.GetSysMenuList)
	router.GET("/api/sysMenu/info", api.GetSysMenu)
	router.PUT("/api/sysMenu/update", api.UpdateSysMenu)
	router.DELETE("/api/sysMenu/delete", api.DeleteSysMenu)
	router.POST("/api/sysRole/add", api.CreateSysRole)
	router.GET("/api/sysRole/list", api.GetSysRoleList)
	router.GET("/api/sysRole/info", api.GetSysRole)
	router.PUT("/api/sysRole/update", api.UpdateSysRole)
	router.DELETE("/api/sysRole/delete", api.DeleteSysRole)
	router.PUT("/api/sysRole/updateStatus", api.UpdateSysRoleStatus)
	router.GET("/api/sysRole/vo/list", api.GetSysRoleVoList)
}
