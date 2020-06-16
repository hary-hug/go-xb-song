package service

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-xb-song/applications/app/model"
	"strings"
)

// return a content's list query by serie's name
func GetpoemsBySerieName(ctx *gin.Context) (res []interface{}, err error) {

	page, _ := com.StrTo(ctx.Query("page")).Int()
	serie := com.StrTo(ctx.Query("serie")).String()
	token := com.StrTo(ctx.GetHeader("x-token")).String()

	var (
		poems  []*model.Poem
		favouriteIds []int
	)

	offset := 0
	limit  := 10
	if page > 0 {
		offset = int(page-1) * limit
	}

	db := model.Db.Model(model.Poem{})
	db = db.Where("status = ?", 1)
	db = db.Where("seriebook like ?", "%" + serie + "%")
	db = db.Order("sort desc")
	db = db.Offset(offset).Limit(limit)


	if err = db.Find(&poems).Error; err != nil {
		return
	}

	// 获取我的收藏
	user, _ := GetUserByToken(token)
	favouriteIds = make([]int, 0)
	if user.Uid != 0 {
		favouriteIds, _ = model.GetUserFavouriteIds(user.Uid)
	}

	for i := range poems {

		item := make(map[string]interface{})
		// 判断是否已被收藏
		isFavour := 0
		for f := range favouriteIds {
			if poems[i].PoemId == favouriteIds[f] {
				isFavour = 1
				break
			}
		}

		item["poem_id"] = poems[i].PoemId
		item["title"]   = poems[i].Title
		item["author"]  = poems[i].Author
		item["dynasty"] = poems[i].Dynasty
		item["digest"]  = strings.Trim(poems[i].Digest, "\n")
		item["tags"] = strings.Split(poems[i].Tags, ",")
		item["is_favour"] = isFavour

		res = append(res, item)
	}

	return
}

// return poems query by keyword
func GetPoemByKeyword(ctx *gin.Context) (res []interface{}, err error) {

	page := com.StrTo(ctx.Query("page")).MustInt()
	keyword := com.StrTo(ctx.Query("k")).String()
	token := com.StrTo(ctx.GetHeader("x-token")).String()

	var (
		poems  []*model.Poem
		favouriteIds []int
	)

	offset := 0
	limit  := 10
	if page > 0 {
		offset = int(page-1) * limit
	}

	db := model.Db.Model(model.Poem{})
	db = db.Where("status = ?", 1)
	db = db.Where("title like ? or tags like ? or content like ? or author like ?",
		"%" + keyword +"%",
		"%" + keyword +"%",
		"%" + keyword +"%",
		"%" + keyword +"%")
	db = db.Offset(offset).Limit(limit)


	if err = db.Find(&poems).Error; err != nil {
		return
	}

	// 获取我的收藏
	user, _ := GetUserByToken(token)
	favouriteIds = make([]int, 0)
	if user.Uid != 0 {
		favouriteIds, _ = model.GetUserFavouriteIds(user.Uid)
	}


	for i := range poems {

		item := make(map[string]interface{})
		// 判断是否已被收藏
		isFavour := 0
		for f := range favouriteIds {
			if poems[i].PoemId == favouriteIds[f] {
				isFavour = 1
				break
			}
		}

		item["poem_id"] = poems[i].PoemId
		item["title"]   = poems[i].Title
		item["author"]  = poems[i].Author
		item["dynasty"] = poems[i].Dynasty
		item["digest"]  = strings.Trim(poems[i].Digest, "\n")
		item["tags"] = strings.Split(poems[i].Tags, ",")
		item["is_favour"] = isFavour

		res = append(res, item)
	}

	return
}

// return a poem's detail query by poem's id
func GetPoemById(poemId int) (res map[string]interface{}, err error) {

	var (
		poem   model.Poem
		tags   []string
	)

	db := model.Db.Model(model.Poem{})
	db = db.Where("status = ?", 1)
	db = db.Where("poem_id = ?", poemId)
	if err = db.First(&poem).Error; err != nil {
		return
	}

	// tag
	tags = make([]string, 0)
	if poem.Tags != "" {
		tags = strings.Split(poem.Tags, ",")
	}

	fmt.Println(tags)

	item := make(map[string]interface{})
	item["poem_id"] = poem.PoemId
	item["title"] = poem.Title
	item["content"] = poem.Content
	item["author"] = poem.Author
	item["dynasty"] = poem.Dynasty
	item["translation"] = poem.Translation
	item["remark"] = poem.Remark
	item["appreciate"] = poem.Appreciate
	item["pinyin"] = poem.Pinyin
	item["audio"] = poem.Audio
	item["tags"] = tags

	res = item

	return
}