package service

import (
	"go-xb-song/applications/app/model"
	"strings"
	"time"
)

// GetFavourites returns user's favourite poems
func GetFavourites(uid int, group string, page int) (res []interface{}, err error) {

	var (
		favourites  []*model.FavouriteDetail
	)

	offset := 0
	limit  := 10
	if page > 0 {
		offset = int(page-1) * limit
	}

	db := model.Db.Model(model.FavouriteDetail{})
	db = db.Select("xb_poem.*, xb_favourite.*")
	db = db.Where("xb_poem.status = ?", 1)
	if group != "" {
		db = db.Where("xb_favourite.fgroup = ?", group)
	}

	db = db.Where("xb_favourite.uid = ?", uid)
	db = db.Joins("left join xb_poem on xb_favourite.poem_id=xb_poem.poem_id")
	db = db.Order("xb_favourite.weight desc, xb_favourite.create_time desc")
	db = db.Offset(offset).Limit(limit)


	if err = db.Find(&favourites).Error; err != nil {
		return
	}

	for i := range favourites {

		item := make(map[string]interface{})

		item["poem_id"] = favourites[i].PoemId
		item["weight"]  = favourites[i].Weight
		item["title"]   = favourites[i].Poem.Title
		item["author"]  = favourites[i].Poem.Author
		item["dynasty"] = favourites[i].Poem.Dynasty
		item["digest"]  = favourites[i].Poem.Digest

		res = append(res, item)
	}

	return
}

// AddFavourite add user's favourite poem with weight
// returns favourite's id
func AddFavourite(uid int, poemId int, weight int) int {

	var (
		favourite   model.Favourite
		poem        model.Poem
	)

	if model.CheckFavouriteExists(uid, poemId) {
		return 0
	}


	db := model.Db.Model(model.Poem{})
	db = db.Where("poem_id = ?", poemId)

	if err := db.First(&poem).Error; err != nil {
		return 0
	}
	// 暂时添加分组
	group := ""
	if strings.Index(poem.Seriebook, "唐诗三百首") >= 0 {
		group = "唐诗三百首"
	} else if strings.Index(poem.Seriebook, "宋词三百首") >= 0 {
		group = "宋词三百首"
	} else if strings.Index(poem.Seriebook, "诗经") >= 0 {
		group = "诗经"
	}


	db = model.Db.Model(model.Favourite{})
	favourite.Uid = uid
	favourite.Fgroup = group
	favourite.PoemId = poemId
	favourite.Weight = weight
	favourite.CreateTime = int(time.Now().Unix())
	db.Create(&favourite)

	return favourite.FavId
}

// CancelFavourite delete user's favourite poem
// returns error or nil
func CancelFavourite(uid int, poemId int) error  {

	db := model.Db.Where("uid = ? and poem_id = ?", uid, poemId)

	if err := db.Delete(&model.Favourite{}).Error; err != nil {
		return err
	}

	return nil
}