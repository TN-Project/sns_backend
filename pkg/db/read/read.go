package read

import (
	"log/slog"
	"sns_backend/pkg/common/model"
	"sns_backend/pkg/db"
)

// ユーザ名からユーザ情報を取得
func GetUser(username string) (model.User, error) {
	db := db.Connect()
	defer db.Close()

	sql := `select user_id, nickname, username, password from "user" where username = $1`
	var user model.User
	err := db.QueryRow(sql, username).Scan(&user.User_id, &user.Nickname, &user.Username, &user.Password)
	if err != nil {
		slog.Error("Error getting user: " + err.Error())
		return model.User{}, err
	}

	return user, nil
}

// ユーザ名が存在するか確認
func ExistUser(username string) (bool, error) {
	db := db.Connect()
	defer db.Close()

	sql := `select count(*) from "user" where username = $1`
	var count int
	err := db.QueryRow(sql, username).Scan(&count)
	if err != nil {
		slog.Error("Error checking user existence: " + err.Error())
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// ユーザが所属するグループを取得
func GetUsersGroup(username string) ([]model.Group, error) {
	db := db.Connect()
	defer db.Close()

	// usernameからユーザ情報を取得
	userdata, err := GetUser(username)
	if err != nil {
		return nil, err
	}

	sql := `SELECT "group"."group_id", "group"."group_name"
			FROM "group"
			JOIN "user_group" ON "group"."group_id" = "user_group"."group_id"
			WHERE "user_group"."user_id" = $1`

	rows, err := db.Query(sql, userdata.User_id)
	if err != nil {
		slog.Error("Error getting user's group: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var groups []model.Group
	for rows.Next() {
		var group model.Group
		err := rows.Scan(&group.Group_id, &group.Group_name)
		if err != nil {
			slog.Error("Error scanning user's group: " + err.Error())
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

// グループ名からグループ情報を取得
func GetGroup(group_name string) (model.Group, error) {
	db := db.Connect()
	defer db.Close()

	sql := `select group_id, group_name from "group" where group_name = $1`
	var group model.Group
	err := db.QueryRow(sql, group_name).Scan(&group.Group_id, &group.Group_name)
	if err != nil {
		slog.Error("Error getting group: " + err.Error())
		return model.Group{}, err
	}

	return group, nil
}

// グループに所属するユーザを取得
func GetGroupsUser(group_name string) ([]model.User, error) {
	db := db.Connect()
	defer db.Close()

	// group_nameからグループ情報を取得
	groupdata, err := GetGroup(group_name)
	if err != nil {
		return nil, err
	}

	sql := `SELECT "user"."user_id", "user"."nickname", "user"."username", "user"."password"
			FROM "user"
			JOIN "user_group" ON "user"."user_id" = "user_group"."user_id"
			WHERE "user_group"."group_id" = $1`

	rows, err := db.Query(sql, groupdata.Group_id)
	if err != nil {
		slog.Error("Error getting group's user: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.User_id, &user.Nickname, &user.Username, &user.Password)
		if err != nil {
			slog.Error("Error scanning group's user: " + err.Error())
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// グループに登録されている写真を取得
func GetPicture(group_id int) ([]string, error) {
	db := db.Connect()
	defer db.Close()

	sql := `select picture_id from "picture" where group_id = $1`
	rows, err := db.Query(sql, group_id)
	if err != nil {
		slog.Error("Error getting picture: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var picture_id []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			slog.Error("Error getting picture: " + err.Error())
			return nil, err
		}
		picture_id = append(picture_id, id)
	}
	slog.Info("Picture got successfully")
	return picture_id, nil
}
