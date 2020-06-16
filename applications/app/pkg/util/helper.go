package util

import (
	"crypto/md5"
	"fmt"
	"go-xb-song/applications/app/pkg/conf"
	"log"
	"strings"
)

// get image hostname from app.ini
func GetImageHost() string {

	sec, err := conf.IniFile.GetSection("static")

	if err != nil {
		log.Fatalln("Fail to get section 'static': ", err)
	}

	return sec.Key("IMG_HOST").String()
}

// get abs path of image
func GetImageUrl(img string, suffix string) string {

	if img == "" {
		return img
	}

	// 非外部图片，可加裁剪参数
	if strings.Index(img, "http") < 0 {
		img = img + suffix
	}

	return fmt.Sprintf("%s%s", GetImageHost(), img)

}

// 获取poemId md5加密的sign
func GetPoemIdSign(poemId int) string {

	str := fmt.Sprintf("%d%s", poemId, "E81vlldNEBKnRcTB")
	data := []byte(str)

	return fmt.Sprintf("%x", md5.Sum(data))
}

