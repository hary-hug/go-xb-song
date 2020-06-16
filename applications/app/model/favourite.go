package model

// the structure of table 'xb_favourite'
type Favourite struct {
	FavId      int     `gorm:"column:fav_id;primary_key" json:"fav_id"`
	Uid        int     `gorm:"column:uid" json:"uid"`
	Fgroup     string  `gorm:"column:fgroup" json:"fgroup"`
	PoemId     int     `gorm:"column:poem_id" json:"poem_id"`
	Weight     int     `gorm:"column:weight" json:"weight"`
	CreateTime int     `gorm:"column:create_time" json:"create_time"`
}

type FavouriteDetail struct {
	FavId      int   `gorm:"column:fav_id;primary_key" json:"fav_id"`
	Uid        int   `gorm:"column:uid" json:"uid"`
	PoemId     int   `gorm:"column:poem_id" json:"poem_id"`
	Weight     int   `gorm:"column:weight" json:"weight"`
	CreateTime int   `gorm:"column:create_time" json:"create_time"`
	Poem
}

// return table name of struct 'Favourite'
func (Favourite) TableName() string {
	return "xb_favourite"
}

// return table name of struct 'FavouriteDetail'
func (FavouriteDetail) TableName() string {
	return "xb_favourite"
}

// CheckFavouriteExists returns favourite's id
func CheckFavouriteExists(uid int, poemId int) bool {

	var (
		favourite Favourite
	)

	db := Db.Model(Favourite{})
	db = db.Select("fav_id")
	db = db.Where("uid = ? and poem_id = ?", uid, poemId)
	db.Find(&favourite)

	if favourite.FavId > 0 {
		return true
	}

	return false
}

// GetUserFavouriteIds returns a slice of user's favourite poem's id
func GetUserFavouriteIds(uid int) (ids []int, err error) {

	var favourites []Favourite

	db := Db.Model(Favourite{})
	db = db.Where("uid = ?", uid)

	if err = db.Find(&favourites).Error; err != nil {
		return
	}

	for i := range favourites {
		ids = append(ids, favourites[i].PoemId)
	}

	return
}