package db

import (
	"database/sql"
	"time"
)

// Menu 菜单实体表
type Menu struct {
	Id         int64     `json:"id" db:"id"`
	RecordId   string    `json:"record_id" db:"record_id"`
	Name       string    `json:"name" db:"name"`
	Sequence   int64     `json:"sequence" db:"sequence"`
	Icon       string    `json:"icon" db:"icon"`
	Router     string    `json:"router" db:"router"`
	Hidden     int64     `json:"hidden" db:"hidden"`
	ParentId   string    `json:"parent_id" db:"parent_id"`
	ParentPath string    `json:"parent_path" db:"parent_path"`
	Creator    string    `json:"creator" db:"creator"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	DeletedAt  time.Time `json:"deleted_at" db:"deleted_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// 根据name查询指定菜单数据
func SelectMenuByName(name string) (*Menu, error) {
	menu := Menu{}
	err := db.Get(&menu, "SELECT * FROM menu where `name` = ?", name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &menu, err
}

// 根据record查询指定菜单数据
func SelectMenuByRecord(record string) (*Menu, error) {
	menu := Menu{}
	err := db.Get(&menu, "SELECT * FROM menu where `record_id` = ?", record)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &menu, err
}

// 根据ParentID父级内码查询指定菜单数据
func SelectMenuByParentID(parentId string) ([]*Menu, error) {
	var menu []*Menu
	err := db.Select(&menu, "SELECT * FROM menu where `parent_id` = ?",

		parentId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return menu, err
}

func InsertMenu(sequence, hidden int64, icon, router, parentId, parentPath, creator, record, name string) error {
	_, err := db.Exec("INSERT INTO `menu` ( `record_id`, `name`, `sequence`, `hidden`, `icon`, `router`, `parent_id`, `parent_path`, `creator`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		record, name, sequence, hidden, icon, router, parentId, parentPath, creator)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func DeleteMenu(record string) error {
	_, err := db.Exec("DELETE FROM `menu` WHERE `record_id` = ?", record)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateMenu(menu *Menu) error {
	_, err := db.Exec("UPDATE `menu` SET `name`=?, `sequence`=?, `hidden`=?, `icon`=?, `router`=?, `parent_id`=?, `parent_path`=?, `creator`=? WHERE `record_id` = ?",
		menu.Name, menu.Sequence, menu.Hidden, menu.Icon, menu.Router, menu.ParentId, menu.ParentPath, menu.Creator, menu.RecordId)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

// 获取菜单列表
func SelectMenuItems() ([]*Menu, error) {
	var menus []*Menu
	err := db.Select(&menus, "SELECT * FROM `menu`")
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return menus, err
}
