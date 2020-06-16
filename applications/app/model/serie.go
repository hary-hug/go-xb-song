package model


// SerieBook is Struct of table 'xb_seriebook'
type SerieBook struct {
	SerieId  int     `gorm:"column:serie_id" json:"SerieId"`
	Name     string  `gorm:"column:name" json:"name"`
	Image    string  `gorm:"column:image" json:"image"`
	Intro    string  `gorm:"column:intro" json:"intro"`
	Status   int     `gorm:"column:status" json:"status"`
}

// TableName returns table name of struct 'SerieBook'
func (SerieBook) TableName() string {
	return "xb_seriebook"
}

