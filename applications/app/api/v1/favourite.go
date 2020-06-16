package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/model"
	"go-xb-song/applications/app/service"
	"net/http"
)

func GetFavourites(ctx *gin.Context)  {

	var (
		err   error
		user  model.User
		res  []interface{}
	)

	token := com.StrTo(ctx.GetHeader("x-token")).String()
	page := com.StrTo(ctx.Query("page")).MustInt()
	group := com.StrTo(ctx.Query("group")).String()


	user, err = service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : err.Error(),
			"data" : make([]interface{}, 0),
		})
		return
	}

	res, err = service.GetFavourites(user.Uid, group, page)

	if len(res) <= 0 {
		res = make([]interface{}, 0)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : res,
	})

}


func GetFavouriteGroup(ctx *gin.Context)  {

	type rs struct{
		Fgroup  string  `json:"group"`
		Total   int     `json:"total"`
	}

	var (
		err        error
		user       model.User
		rss        []rs
		res        []interface{}
	)

	token := com.StrTo(ctx.GetHeader("x-token")).String()

	user, err = service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : err.Error(),
			"data" : make([]interface{}, 0),
		})
		return
	}


	db := model.Db.Table("xb_favourite")
	db = db.Select("distinct(fgroup) as fgroup, count(fav_id) as total")
	db = db.Where("uid = ?", user.Uid)
	db = db.Group("fgroup")

	if err = db.Find(&rss).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "success",
			"data" : make([]interface{}, 0),
		})
		return
	}

	for i := range rss {
		if rss[i].Fgroup == "" {
			continue
		}
		item := make(map[string]interface{})
		item["group"] = rss[i].Fgroup
		item["total"] = rss[i].Total
		res = append(res, item)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 0,
		"msg"  : "success",
		"data" : res,
	})
	return
}


func AddFavourite(ctx *gin.Context)  {

	var (
		err   error
		user  model.User
	)

	reqParma := new(struct{
		PoemId    int  `form:"poem_id" json:"poem_id" binding:"required"`
		Weight    int  `form:"weight" json:"weight" binding:"required"`
	})

	token := com.StrTo(ctx.GetHeader("x-token")).String()

	if err := ctx.ShouldBindJSON(&reqParma); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "token: " + err.Error(),
			"data" : make(map[string]interface{}),
		})
		return
	}

	// 新增
	favId := service.AddFavourite(user.Uid, reqParma.PoemId, reqParma.Weight)

	if favId == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "已收藏",
			"data" : make(map[string]interface{}),
		})
		return
	}

	res := make(map[string]interface{})
	res["fav_id"] = favId

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : res,
	})

}


func CancelFavourite(ctx *gin.Context)  {
	var (
		err   error
		user  model.User
	)

	reqParma := new(struct{
		PoemId    int  `form:"poem_id" json:"poem_id" binding:"required"`
	})

	token := com.StrTo(ctx.GetHeader("x-token")).String()

	if err := ctx.ShouldBindJSON(&reqParma); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = service.GetUserByToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "出错了: " + err.Error(),
			"data" : make(map[string]interface{}),
		})
		return
	}
	// 删除
	err = service.CancelFavourite(user.Uid, reqParma.PoemId)
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