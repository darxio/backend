package groups

import (
	"backend/internal/database/connection"
	"backend/internal/models"
	"log"
	"strconv"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func All(groups *models.GroupArr) (code int, message string) {
	rows, err := database.Query(`SELECT id, name, about FROM groups;`)

	if err != nil {
		log.Println("database/groups.go: 500, " + err.Error())
		return 500, err.Error()
	}

	for rows.Next() {
		group := models.Group{}
		rows.Scan(&group.ID, &group.Name, &group.About)
		*groups = append(*groups, &group)
	}
	rows.Close()

	return 200, "Successful."
}

func About(groupName string, groupID int32, group *models.Group) (code int, message string) {
	var err error
	if groupID != 0 {
		err = database.QueryRow(`SELECT id, name, about FROM groups WHERE id = $1;`, groupID).Scan(&group.ID, &group.Name, &group.About)
	} else {
		err = database.QueryRow(`SELECT id, name, about FROM groups WHERE name = $1;`, groupName).Scan(&group.ID, &group.Name, &group.About)
	}

	if err == pgx.ErrNoRows {
		log.Println("database/groups.go: 404, " + err.Error())
		return 404, "Group not found."
	} else if err != nil {
		log.Println("database/groups.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return 200, "Successful."
}

func Ingredients(groupID int, count int, offset int, ingredients *models.IngredientArr) (code int, message string) {
	rows, err := database.Query(`
		SELECT id, name, danger, description, wiki_link 
			FROM ingredients WHERE id IN (
				SELECT id FROM ing_groups WHERE (NOT ($1 = ANY (groups))) AND (array_length(groups, 1) > 0)
			)
				ORDER BY frequency DESC, danger DESC LIMIT $2 OFFSET $3
				`, strconv.Itoa(groupID), count, offset)

	if err == pgx.ErrNoRows {
		return 404, "Group not found."
	} else if err != nil {
		return 500, err.Error()
	}

	for rows.Next() {
		curIng := models.Ingredient{}
		rows.Scan(
			&curIng.ID, &curIng.Name, &curIng.Danger,
			&curIng.Description, &curIng.WikiLink)
		*ingredients = append(*ingredients, &curIng)
	}

	return 200, "Successful."
}

func Search(groupsName string, groups *models.GroupArr) (code int, message string) {
	rows, err := database.Query(`
		SELECT id, name, about 
			FROM groups WHERE name LIKE '%' || $1 || '%' 
				ORDER BY id
				`, groupsName)

	if err == pgx.ErrNoRows {
		return 404, "Ingredient not found."
	} else if err != nil {
		return 500, err.Error()
	}

	for rows.Next() {
		curGroup := models.Group{}
		rows.Scan(&curGroup.ID, &curGroup.Name, &curGroup.About)
		*groups = append(*groups, &curGroup)
	}

	return 200, "Successful."
}
