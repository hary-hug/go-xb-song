package model


type Author struct {
	AuthorId  int      `gorm:"column:author_id" json:"author_id"`
	Name      string   `gorm:"column:name" json:"name"`
	Image     string   `gorm:"column:image" json:"image"`
	Dynasty   string   `gorm:"column:dynasty" json:"dynasty"`
	Intro     string   `gorm:"column:intro" json:"intro"`
}

// TableName returns table name of struct 'Author'
func (Author) TableName() string {
	return "xb_author"
}