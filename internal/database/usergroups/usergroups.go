package groups

import (
	"backend/internal/database/connection"
	"backend/internal/models"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func All(username string, groups *models.GroupArr) (code int, message string) {
	rows, err := database.Query(`SELECT id, name, about
								FROM groups
								JOIN user_groups
								ON groups.id = user_groups.group_id
								WHERE id = (SELECT id FROM users WHERE username = $1);`, username)

	if err != nil {
		return 500, "Something went wrong.."
	}

	for rows.Next() {
		group := models.Group{}
		rows.Scan(&group.ID, &group.Name, &group.About)
		*groups = append(*groups, &group)
	}
	rows.Close()

	return 200, "Successful."
}

func About(username string, groupName string, groupID int32, group *models.Group) (code int, message string) {
	var err error
	if groupID != 0 {
		err = database.QueryRow(`SELECT id, name, about
								FROM groups
								JOIN user_groups
								ON groups.id = user_groups.group_id
								WHERE user_id = (SELECT id FROM users WHERE username = $1)
								AND group_id = $2;`, username, groupID).Scan(&group.ID, &group.Name, &group.About)
	} else {
		err = database.QueryRow(`SELECT id, name, about
								FROM groups
								JOIN user_groups
								ON groups.id = user_groups.group_id
								WHERE user_id = (SELECT id FROM users WHERE username = $1)
								AND name = $2;`, username, groupName).Scan(&group.ID, &group.Name, &group.About)
	}

	if err == pgx.ErrNoRows {
		return 404, "Group not found."
	} else if err != nil {
		return 500, "Something went wrong.."
	}

	return 200, "Successful."
}

func Add(username string, groupName string, groupID int32, groups *models.GroupArr) (code int, message string) {
	var err error
	if groupID != 0 {
		_, err = database.Exec("INSERT INTO user_groups(user_id, group_id) VALUES ((SELECT id FROM users WHERE username = $1), $2);", username, groupID)
	} else {
		_, err = database.Exec("INSERT INTO user_groups(user_id, group_id) VALUES ((SELECT id FROM users WHERE username = $1), (SELECT id FROM groups WHERE name = $2));", username, groupName)
	}

	if err != nil {
		pgErr := err.(pgx.PgError)
		if pgErr.Code == "23505" {
			return 409, "This user is already in this group."
		}
		if pgErr.Code == "23503" {
			return 404, "This group doesn't exist."
		}
		return 500, "Something went wrong.."
	}

	return All(username, groups)
}

func Delete(username string, groupName string, groupID int32, groups *models.GroupArr) (code int, message string) {
	var err error
	if groupID != 0 {
		_, err = database.Exec("DELETE FROM user_groups WHERE user_id = (SELECT id FROM users WHERE username = $1) AND group_id = $2;", username, groupID)
	} else {
		_, err = database.Exec("DELETE FROM user_groups WHERE user_id = (SELECT id FROM users WHERE username = $1) AND group_id = (SELECT id FROM groups WHERE name = $2);", username, groupName)
	}

	if err != nil {
		return 500, "Something went wrong.."
	}

	return All(username, groups)
}
