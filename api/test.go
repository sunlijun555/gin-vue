package api

import (
	"go-vue/result"

	"github.com/gin-gonic/gin"
)

// Success 成功测试
// @Summary Success测试接口
// @Tags 测试相关接口
// @Produce json
// @Description 成功测试接口
// @Success 200 {object} result.Result
// @router /api/success [get]
func Success(c *gin.Context) {
	result.Success(c, 200)
}

// Failed 失败测试
// @Summary Failed测试接口
// @Tags 测试相关接口
// @Produce json
// @Description 失败测试接口
// @Success 200 {object} result.Result
// @router /api/failed [get]
func Failed(c *gin.Context) {
	result.Failed(c, int(result.Apicode.Failed), result.Apicode.GetMessage(result.Apicode.Failed))

}
