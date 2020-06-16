package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/service"
	"net/http"
)

func GetUserDetail(ctx *gin.Context)  {

	token := com.StrTo(ctx.GetHeader("x-token")).String()

	user, err := service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "获取数据失败" + err.Error(),
			"data" : make(map[string]interface{}),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : user,
	})
}
