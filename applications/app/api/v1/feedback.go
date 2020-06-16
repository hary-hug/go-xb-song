package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/model"
	"go-xb-song/applications/app/service"
	"net/http"
	"time"
)

func AddFeedback(ctx *gin.Context)  {

	var (
		err       error
		user      model.User
		feedback  model.Feedback
	)

	// 请求参数
	reqParma := new(struct {
		Content  string  `form:"content" json:"content" binding:"required"`
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

	db := model.Db.Model(model.Feedback{})
	feedback.Uid = user.Uid
	feedback.Content = reqParma.Content
	feedback.CreateTime = int(time.Now().Unix())
	db.Create(&feedback)

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : make(map[string]interface{}),
	})
}
