package create

import (
	"log/slog"
	"sns_backend/pkg/common/random"
	"sns_backend/pkg/db"
)

var sql_stm []string = []string{
	`create table IF NOT EXISTS "user" (user_id bigint PRIMARY KEY, nickname text NOT NULL, username text UNIQUE, password text)`,
	`create table IF NOT EXISTS "group" (group_id bigint PRIMARY KEY, group_name text)`,
	`create table IF NOT EXISTS "user_group" (user_id bigint, group_id bigint, PRIMARY KEY(user_id, group_id), FOREIGN KEY(user_id) REFERENCES "user"(user_id), FOREIGN KEY(group_id) REFERENCES "group"(group_id))`,
}

func CreateDefaultTable() {
	db := db.Connect()
	defer db.Close()

	for _, sql := range sql_stm {
		_, err := db.Exec(sql)
		if err != nil {
			slog.Error("Error creating table: " + err.Error())
			return
		}
	}
	slog.Info("Tables created successfully")
}

func CreateUser(user_id int, nickname string, username string, hashed_password string) error {
	db := db.Connect()
	defer db.Close()

	sql := `insert into "user" (user_id, nickname, username, password) values ($1, $2, $3, $4)`
	_, err := db.Exec(sql, user_id, nickname, username, hashed_password)
	if err != nil {
		slog.Error("Error creating user: " + err.Error())
		return err
	}
	slog.Info("User created successfully")
	return nil
}

// グループを作成
func CreateGroup(group_name string) (int, error) {
	db := db.Connect()
	defer db.Close()

	// ランダムなグループIDを生成
	group_id := random.GenerateRandomInt()

	sql := `insert into "group" (group_id, group_name) values ($1, $2)`
	_, err := db.Exec(sql, group_id, group_name)
	if err != nil {
		slog.Error("Error creating group: " + err.Error())
		return -1, err
	}
	slog.Info("Group created successfully")
	return group_id, nil
}

// ユーザをグループに追加
func AddUserToGroup(user_id []int, group_id int) error {
	db := db.Connect()
	defer db.Close()

	for _, id := range user_id {
		sql := `insert into "user_group" (user_id, group_id) values ($1, $2)`
		_, err := db.Exec(sql, id, group_id)
		if err != nil {
			slog.Error("Error adding user to group: " + err.Error())
			return err
		}
	}
	slog.Info("User added to group successfully")
	return nil
}

// 写真を登録
func CreatePicture(picture_id []string, group_id int) error {
	db := db.Connect()
	defer db.Close()

	for _, id := range picture_id {
		sql := `insert into "picture" (picture_id, group_id) values ($1, $2)`
		_, err := db.Exec(sql, id, group_id)
		if err != nil {
			slog.Error("Error creating picture: " + err.Error())
			return err
		}
	}
	slog.Info("Picture created successfully")
	return nil
}
