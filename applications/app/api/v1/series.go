package v1

import (
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/service"
	"net/http"
)

func GetSeries(ctx *gin.Context)  {

	var (
		err  error
		res  []interface{}
	)

	// query condition of service
	qc := make(map[string]interface{})


	if res, err = service.GetSeries(qc); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "获取数据失败" + err.Error(),
			"data" : make([]interface{}, 0),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : res,
	})

}