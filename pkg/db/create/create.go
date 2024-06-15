package create

import (
	"log/slog"
	"sns_backend/pkg/db"
)

var sql_stm []string = []string{
	`create table IF NOT EXISTS "user" (user_id bigint PRIMARY KEY, nickname text NOT NULL, username text UNIQUE, password text)`,
	`create table IF NOT EXISTS "group" (group_id int PRIMARY KEY, group_name text)`,
	`create table IF NOT EXISTS "user_group" (user_id int, group_id int, PRIMARY KEY(user_id, group_id), FOREIGN KEY(user_id) REFERENCES "user"(user_id), FOREIGN KEY(group_id) REFERENCES "group"(group_id))`,
}

func CreateDefaultTable() {
	db := db.Connect()
	defer db.Close()

	for _, sql := range sql_stm {
		_, err := db.Exec(sql)
		if err != nil {
			slog.Error("Error creating table: ", err)
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
		slog.Error("Error creating user: ", err)
		return err
	}
	slog.Info("User created successfully")
	return nil
}
