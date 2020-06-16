package model

type Feedback struct {
	Uid     int      `gorm:"column:uid"`
	Content string   `gorm:"column:content"`
	CreateTime int   `gorm:"column:create_time"`
}

func (Feedback) TableName() string {
	return "xb_feedback"
}
