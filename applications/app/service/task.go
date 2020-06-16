package service

import (
	"go-xb-song/applications/app/model"
)

// InitTask initial a task of today
func InitTask(uid int, poemId int, day string) (task model.Task) {

	//t := time.Now()
	//day := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())

	if task = model.CheckTaskExists(uid, poemId, day); task.TaskId > 0 {
		return task
	}

	var (
		taskSeting model.TaskSetting
	)

	// 获取我的任务设置
	db := model.Db.Model(model.TaskSetting{})
	db = db.Where("uid = ?", uid)
	db.Find(&taskSeting)


	db = model.Db.Model(model.Task{})
	task.Uid = uid
	task.PoemId = poemId
	task.Day = day
	task.Read = 0
	task.Write = 0
	task.Mrite = 0
	if taskSeting.Read >= 0 {
		task.ReadRequire = taskSeting.Read
	} else {
		// 默认
		task.ReadRequire = 5
	}

	if taskSeting.Write >= 0 {
		task.WriteRequire = taskSeting.Write
	} else {
		task.WriteRequire = 5
	}

	if taskSeting.Mrite >= 0 {
		task.MriteRequire = taskSeting.Mrite
	} else {
		task.MriteRequire = 5
	}

	db.Create(&task)

	return task

}

// AddCount update the value of column 'read', 'write', 'mrite'
func AddCount(taskId int, column int) error {

	var (
		err error
		task model.Task
	)

	db := model.Db.Model(model.Task{})
	db = db.Where("task_id = ?", taskId)
	db.Find(&task)
	switch column {
	case 1:
		// read count
		err = db.Update("read", task.Read + 1).Error
		break
	case 2:
		// write count
		err = db.Update("write", task.Write + 1).Error
		break
	case 3:
		// mrite count
		err = db.Update("mrite", task.Mrite + 1).Error
		break
	}

	if err != nil {
		return err
	}

	return nil
}

// GetTaskSetting returns user task setting
func GetTaskSetting(uid int) (taskSeting model.TaskSetting) {

	db := model.Db.Model(model.TaskSetting{})
	db = db.Where("uid = ?", uid)
	db.Find(&taskSeting)

	return
}
