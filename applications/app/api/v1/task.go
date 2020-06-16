package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/model"
	"go-xb-song/applications/app/service"
	"net/http"
)

func InitTask(ctx *gin.Context)  {

	var (
		err   error
		user  model.User
	)

	// 请求参数
	reqParma := new(struct {
		PoemId  int  `form:"poem_id" json:"poem_id" binding:"required"`
		Day     string   `form:"day" json:"day" binding:"required"`
	})

	if err := ctx.ShouldBindJSON(&reqParma); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := com.StrTo(ctx.GetHeader("x-token")).String()

	user, err = service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "token: " + err.Error(),
			"data" : make(map[string]interface{}),
		})
		return
	}

	// 初始化当日任务
	task := service.InitTask(user.Uid, reqParma.PoemId, reqParma.Day)

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : task,
	})

}


func AddCount(ctx *gin.Context)  {

	var (
		err   error
	)

	// 请求参数
	reqParma := new(struct {
		TaskId    int  `form:"task_id" json:"task_id" binding:"required"`
		Column    int  `form:"column" json:"column" binding:"required"`
	})

	if err := ctx.ShouldBindJSON(&reqParma); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := com.StrTo(ctx.GetHeader("x-token")).String()

	_, err = service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "出错了: " + err.Error(),
			"data" : make(map[string]interface{}),
		})
		return
	}

	err = service.AddCount(reqParma.TaskId, reqParma.Column)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "出错了: " + err.Error(),
			"data" : make(map[string]interface{}),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : make(map[string]interface{}),
	})

}


func GetTaskSetting(ctx *gin.Context)  {

	var (
		err   error
		user  model.User
		taskSeting  model.TaskSetting
	)

	token := com.StrTo(ctx.GetHeader("x-token")).String()
	user, err = service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "出错了: " + err.Error(),
			"data" : make(map[string]interface{}),
		})
		return
	}

	taskSeting = service.GetTaskSetting(user.Uid)

	if taskSeting.Uid <= 0 {
		// 创建默认设置
		db := model.Db.Model(model.TaskSetting{})
		taskSeting.Uid = user.Uid
		taskSeting.Read = 5
		taskSeting.Write = 5
		taskSeting.Mrite = 5
		db.Create(&taskSeting)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : taskSeting,
	})
}


func SaveTaskSetting(ctx *gin.Context)  {

	var (
		err   error
		user  model.User
		taskSeting  model.TaskSetting
	)

	// 请求参数
	reqParma := new(struct {
		Read    int  `form:"read" json:"read" binding:"required"`
		Write    int  `form:"write" json:"write" binding:"required"`
		Mrite    int  `form:"mrite" json:"mrite" binding:"required"`
	})


	if err := ctx.ShouldBindJSON(&reqParma); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := com.StrTo(ctx.GetHeader("x-token")).String()
	user, err = service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "出错了: " + err.Error(),
			"data" : make(map[string]interface{}),
		})
		return
	}

	db := model.Db.Model(model.TaskSetting{})
	db = db.Where("uid = ?", user.Uid)
	db.Find(&taskSeting)

	// 不使用struct，因为0至会无法更新
	db.Update(map[string]interface{}{
		"read": reqParma.Read,
		"write": reqParma.Write,
		"mrite": reqParma.Mrite,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : make(map[string]interface{}),
	})
}
