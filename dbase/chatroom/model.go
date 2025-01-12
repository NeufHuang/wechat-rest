package chatroom

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建群聊

type CreateParam struct {
	Roomid     string `binding:"required"`
	Name       string
	Level      int32
	Remark     string
	JoinArgot  string
	RevokeMsg  string
	WelcomeMsg string
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Chatroom{
		Roomid:     data.Roomid,
		Name:       data.Name,
		Level:      data.Level,
		Remark:     data.Remark,
		JoinArgot:  data.JoinArgot,
		RevokeMsg:  data.RevokeMsg,
		WelcomeMsg: data.WelcomeMsg,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新群聊

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Chatroom{
			Roomid: data.Roomid,
		}).
		Updates(tables.Chatroom{
			Name:       data.Name,
			Level:      data.Level,
			Remark:     data.Remark,
			JoinArgot:  data.JoinArgot,
			RevokeMsg:  data.RevokeMsg,
			WelcomeMsg: data.WelcomeMsg,
		})

	return result.Error

}

// 合并群聊

type MigrateParam = CreateParam

func Migrate(data *MigrateParam) error {

	item, err := Fetch(&FetchParam{
		Roomid: data.Roomid,
	})

	if err == nil && item.Rd > 0 {
		err = Update(data)
	} else {
		_, err = Create(data)
	}

	return err

}

// 删除群聊

type DeleteParam struct {
	Roomid string `binding:"required"`
}

func Delete(data *DeleteParam) error {

	var item *tables.Chatroom

	result := dborm.Db.
		Where(&tables.Chatroom{
			Roomid: data.Roomid,
		}).
		Delete(&item)

	return result.Error

}

// 获取群聊

type FetchParam struct {
	Roomid string `binding:"required"`
}

func Fetch(data *FetchParam) (*tables.Chatroom, error) {

	var item *tables.Chatroom

	result := dborm.Db.
		Where(&tables.Chatroom{
			Roomid: data.Roomid,
		}).
		First(&item)

	if item == nil {
		item = &tables.Chatroom{Roomid: data.Roomid}
	}

	return item, result.Error

}

// 获取群聊列表

type FetchAllParam struct {
	Level int32
}

func FetchAll(data *FetchAllParam) ([]*tables.Chatroom, error) {

	var items []*tables.Chatroom

	result := dborm.Db.
		Where(&tables.Chatroom{
			Level: data.Level,
		}).
		Find(&items)

	return items, result.Error

}

// 获取群聊总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Chatroom{}).
		Where(&tables.Chatroom{
			Level: data.Level,
		}).
		Count(&count)

	return count, result.Error

}
