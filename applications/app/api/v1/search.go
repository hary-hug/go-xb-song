package v1

import (
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/service"
	"net/http"
)

func SearchContentsByKeyword(ctx *gin.Context)  {

	var (
		err  error
		res  []interface{}
	)

	res, err = service.GetPoemByKeyword(ctx)
	if  err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "获取数据失败" + err.Error(),
			"data" : make([]interface{}, 0),
		})
		return
	}


	if len(res) <= 0 {
		res = make([]interface{}, 0)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : res,
	})

}
