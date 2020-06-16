package model

// Poem is Struct of table 'xb_poem'
type Poem struct {
	PoemId       int      `gorm:"column:poem_id;primary_key" json:"poem_id"`
	Status       int      `gorm:"column:status" json:"status"`
	Title        string   `gorm:"column:title" json:"title"`
	Author       string   `gorm:"column:author" json:"author"`
	Dynasty      string   `gorm:"column:dynasty" json:"dynasty"`
	Content      string   `gorm:"column:content" json:"content"`
	Seriebook    string   `gorm:"column:seriebook" json:"seriebook"`
	Tags         string   `gorm:"column:tags" json:"tags"`
	Audio        int      `gorm:"column:audio" json:"audio"`
	Digest       string   `gorm:"column:digest" json:"digest"`
	Translation  string   `gorm:"column:translation" json:"translation"`
	Remark       string   `gorm:"column:remark" json:"remark"`
	Pinyin       string   `gorm:"column:pinyin" json:"pinyin"`
	Appreciate   string   `gorm:"column:appreciate" json:"appreciate"`
}

// TableName returns table name of struct 'Poem'
func (Poem) TableName() string {
	return "xb_poem"
}
