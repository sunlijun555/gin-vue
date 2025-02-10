package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应结果
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(Apicode.Success)
	res.Message = Apicode.GetMessage(Apicode.Success)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func Failed(c *gin.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}
