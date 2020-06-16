package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/model"
	"go-xb-song/applications/app/pkg/util"
	"go-xb-song/applications/app/pkg/wxbizdatacrypt"
	"gopkg.in/resty.v1"
	"net/http"
	"time"
)


func WxLogin(ctx *gin.Context) {

	// 请求参数
	wxParma := new(struct {
		Code           string `form:"code" json:"code" binding:"required"`
		Iv             string `form:"iv" json:"iv" binding:"required"`
		EncryptedData  string `form:"encrypted_data" json:"encrypted_data" binding:"required"`
	})

	if err := ctx.ShouldBindJSON(&wxParma); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	appid := "wxd8952542e15d0e4e"
	appSecret := "494a3a19c86b69004983f7405f0967a2"

	// get openid and session_key from  wechat api
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"appid": appid,
			"secret": appSecret,
			"js_code": wxParma.Code,
			"grant_type": "authorization_code",
		}).
		Get("https://api.weixin.qq.com/sns/jscode2session")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	session := new(struct {
		SessionKey  string   `json:"session_key"`
		Openid      string   `json:"openid"`
	})

	err = json.Unmarshal(resp.Body(), &session)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	wc := wxbizdatacrypt.NewWXBizDataCrypt(appid, session.SessionKey)
	userInfo, err := wc.Decrypt(wxParma.EncryptedData, wxParma.Iv)
	
	if userInfo.OpenID == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"msg"  : "获取用户信息" + err.Error(),
			"data" : make([]interface{}, 0),
		})
		return
	}

	var (
		user model.User
	)

	db := model.Db.Model(model.User{})
	db = db.Where("openid = ?", userInfo.OpenID)
	db.First(&user)

	if user.Uid <= 0 {
		// 注册
		user = model.User{
			Openid: userInfo.OpenID,
			Nickname: userInfo.NickName,
			Avatar: userInfo.AvatarURL,
			Gender: userInfo.Gender,
			Country: userInfo.Country,
			Province: userInfo.Province,
			City: userInfo.City,
			CreateTime: int(time.Now().Unix()),
			LoginTime: int(time.Now().Unix()),
		}
		model.Db.Model(model.User{}).Create(&user)
	}

	// 生成token
	token, err := util.GenerateToken(user.Uid)

	ctx.JSON(http.StatusOK, gin.H{
		"code" : 1,
		"msg"  : "success",
		"data" : token,
	})

}
