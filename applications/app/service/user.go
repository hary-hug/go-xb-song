package service

import (
	"go-xb-song/applications/app/model"
	"go-xb-song/applications/app/pkg/util"
	"time"
)


func GetUserByToken(token string) (user model.User, err error) {

	claims, err := util.ParseToken(token)

	if err != nil {
		return
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return
	}

	db := model.Db.Model(model.User{})
	db = db.Where("uid = ?", claims.Uid)
	if err = db.First(&user).Error; err != nil {
		return
	}

	return
}

