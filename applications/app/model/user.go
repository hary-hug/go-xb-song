package model


type User struct {
	Uid        int     `gorm:"column:uid;primary_key" json:"uid"`
	Openid     string  `gorm:"column:openid" json:"openid"`
	Nickname   string  `gorm:"column:nickname" json:"nickname"`
	Avatar     string  `gorm:"column:avatar" json:"avatar"`
	Gender     int     `gorm:"column:gender" json:"gender"`
	Country    string  `gorm:"column:country" json:"country"`
	Province   string  `gorm:"column:province" json:"province"`
	City       string  `gorm:"column:city" json:"city"`
	CreateTime int     `gorm:"column:create_time" json:"create_time"`
	LoginTime  int     `gorm:"column:login_time" json:"login_time"`
}

func (User) TableName() string {
	return "xb_user"
}