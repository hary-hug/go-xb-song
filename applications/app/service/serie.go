package service

import (
	"go-xb-song/applications/app/model"
)


// GetSeries return a list of serie
// qs is the query conditions of model
func GetSeries(qs map[string]interface{}) (res []interface{}, err error) {

	var (
		series  []*model.SerieBook
	)

	db := model.Db.Model(model.SerieBook{})
	db = db.Where("status = ?" , 1)


	if err = db.Find(&series).Error; err != nil {
		return
	}

	for i := range series {

		item := make(map[string]interface{})

		item["serie_id"] = series[i].SerieId
		item["name"]   = series[i].Name
		item["image"]   = series[i].Image
		item["intro"]  = series[i].Intro

		res = append(res, item)
	}

	return

}


