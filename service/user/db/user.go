package db

import (
	"database/sql"
	"time"
)

// User 用户表
type User struct {
	UserId     string    `json:"user_id" db:"user_id"`
	RecordId   string    `json:"record_id" db:"record_id"`
	UserName   string    `json:"user_name" db:"user_name"`
	Password   string    `json:"password" db:"password"`
	Email      string    `json:"email" db:"email"`
	Phone      string    `json:"phone" db:"phone"`
	Profile    string    `json:"profile" db:"profile"`
	CreateAt   time.Time `json:"create_at" db:"create_at"`
	LastActive time.Time `json:"last_active" db:"last_active"`
	Status     string    `json:"status" db:"status"`
}

func InsertUser(username string, password string, record string) error {
	stmt, err := db.Prepare(
		"insert into user (`record_id`,`user_name`,`password`,`status`) values (?, ?, ?, 1)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(record, username, password)
	if err != nil {
		return err
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return nil
	}
	return err
}

func SelectUserByUserName(userName string) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM user where `user_name` = ?", userName)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func SelectUserByUserEmail(userEmail string) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM user where `email` = ?", userEmail)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func SelectUserByUserPhone(userPhone string) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM user where `phone` = ?", userPhone)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func SelectUserByUsernamePassword(userName, password string) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM user WHERE `user_name` = ? AND `password` = ? AND `status` = 1 LIMIT 1", userName, password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func UpdateUserPassword(userName, password, record string) error {
	_, err := db.Exec(
		"UPDATE `user` SET `password` = ? WHERE `user_name` = ? AND `record_id` = ? AND `status` = 1",
		password, userName, record)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateUserNameProfile(userName, record string) error {
	_, err := db.Exec("UPDATE `user` SET `user_name` = ? WHERE `record_id` = ? AND `status` = 1", userName, record)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateUserEmailProfile(email, record string) error {
	_, err := db.Exec("UPDATE `user` SET `email` = ? WHERE `record_id` = ? AND `status` = 1", email, record)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func UpdateUserPhoneProfile(phone, record string) error {
	_, err := db.Exec("UPDATE `user` SET `phone` = ? WHERE `record_id` = ? AND `status` = 1", phone, record)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
