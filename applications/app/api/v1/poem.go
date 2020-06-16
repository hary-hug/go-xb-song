package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/pkg/util"
	"go-xb-song/applications/app/service"
	"net/http"
)

func GetPoemsBySerie(ctx *gin.Context)  {

	var (
		err  error
		res  []interface{}
	)

	res, err = service.GetpoemsBySerieName(ctx)
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


func GetPoemDetail(ctx *gin.Context)  {

	var (
		err  error
	)

	res := make(map[string]interface{})

	sign := com.StrTo(ctx.Query("sign")).String()
	contentId := com.StrTo(ctx.Param("id")).MustInt()
	// 校验sign是否正确
	if util.GetPoemIdSign(contentId) != sign {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code" : 0,
			"msg"  : "invalid sign",
			"data" : make(map[string]interface{}),
		})
		return
	}

	res, err = service.GetPoemById(contentId)
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
		"data" : res,
	})

}
