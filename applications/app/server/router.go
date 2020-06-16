package server

import (
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/api/v1"
)

func InitRoutes(e *gin.Engine) {

	g1 := e.Group("/api/v1")
	{
		g1.GET("/series", v1.GetSeries)
		g1.GET("/serie/poems", v1.GetPoemsBySerie)
		g1.GET("/poem/:id", v1.GetPoemDetail)
		g1.GET("/search", v1.SearchContentsByKeyword)
		g1.POST("/wxlogin", v1.WxLogin)
		g1.GET("/user", v1.GetUserDetail)
		g1.GET("/user/favourites", v1.GetFavourites)
		g1.GET("/user/favourite/groups", v1.GetFavouriteGroup)
		g1.POST("/user/favourite/add", v1.AddFavourite)
		g1.POST("/user/favourite/cancel", v1.CancelFavourite)
		g1.GET("/user/task/setting", v1.GetTaskSetting)
		g1.POST("/user/task/setting/save", v1.SaveTaskSetting)
		g1.POST("/task/init", v1.InitTask)
		g1.POST("/task/count/add", v1.AddCount)
		g1.POST("/user/feedback/add", v1.AddFeedback)

	}

}
