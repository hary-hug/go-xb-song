package model

type Task struct {
	TaskId  int  `gorm:"column:task_id; primary_key" json:"task_id"`
	Uid     int  `gorm:"column:uid" json:"uid"`
	PoemId  int  `gorm:"column:poem_id" json:"poem_id"`
	Read    int `gorm:"column:read" json:"read"`
	Write   int `gorm:"column:write" json:"write"`
	Mrite   int `gorm:"column:mrite" json:"mrite"`
	Day     string  `gorm:"column:day" json:"day"`
	ReadRequire int `gorm:"column:read_require" json:"read_require"`
	WriteRequire int `gorm:"column:write_require" json:"write_require"`
	MriteRequire int `gorm:"column:mrite_require" json:"mrite_require"`
}

type TaskSetting struct {
	Uid     int  `gorm:"column:uid; primary_key" json:"uid"`
	Read    int `gorm:"column:read" json:"read"`
	Write   int `gorm:"column:write" json:"write"`
	Mrite   int `gorm:"column:mrite" json:"mrite"`
}

// TableName returns table name of struct 'Task'
func (Task) TableName() string {
	return "xb_task"
}

// TableName returns table name of struct 'TaskSetting'
func (TaskSetting) TableName() string {
	return "xb_task_setting"
}


// CheckTaskExists returns task's id
func CheckTaskExists(uid int, poemId int, day string) (task Task) {

	db := Db.Model(Task{})
	db = db.Where("uid = ? and poem_id = ? and day = ?", uid, poemId, day)
	db.Find(&task)

	return
}