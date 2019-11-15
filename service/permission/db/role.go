package db

import (
	"database/sql"
	"time"
)

// 角色实体表
type Role struct {
	Id        int64     `json:"id" db:"id"`
	Record    string    `json:"record_id" db:"record_id"`
	Name      string    `json:"name" db:"name"`
	Sequence  int64     `json:"sequence" db:"sequence"`
	Memo      string    `json:"memo" db:"memo"`
	Creator   string    `json:"creator" db:"creator"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// 根据name查询指定角色数据
func SelectRoleByName(name string) (*Role, error) {
	role := Role{}
	err := db.Get(&role, "SELECT * FROM `role` WHERE `name` = ?", name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &role, err
}

// 根据record查询指定角色数据
func SelectRoleByRecord(record string) (*Role, error) {
	role := Role{}
	err := db.Get(&role, "SELECT  * FROM `role` WHERE `record_id` = ?", record)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &role, err
}

func InsertRole(role *Role) error {
	_, err := db.Exec("INSERT INTO `role` (`record_id`, `name`, `sequence`, `memo`, `creator`) VALUES (?, ?, ?, ?, ?)",
		role.Record, role.Name, role.Sequence, role.Memo, role.Creator)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateRole(role *Role) error {
	_, err := db.Exec("UPDATE `role` SET `name` = ?, `sequence` = ?, `memo` = ?, `creator` = ? WHERE `record_id` = ?",
		role.Name, role.Sequence, role.Memo, role.Creator, role.Record)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func DeleteRole(record string) error {
	_, err := db.Exec("DELETE FROM `role` WHERE `record_id` = ?", record)
	if err == sql.ErrNoRows {
		return nil
	}
	return nil
}

func QueryRole() {}
