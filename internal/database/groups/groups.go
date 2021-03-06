package groups

import (
	"backend/internal/database/connection"
	"backend/internal/models"
	"log"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/lib/pq"
)

var database *pgx.ConnPool
var hostURL = "https://static.foodwise.rasseki.org/"

func init() {
	database = connection.Connect()
}

func All(groups *models.GroupArr) (code int, message string) {
	rows, err := database.Query(`SELECT id, name, about, image_link FROM groups;`)

	if err != nil {
		log.Println("database/groups.go: 500, " + err.Error())
		return 500, err.Error()
	}

	for rows.Next() {
		group := models.Group{}
		rows.Scan(&group.ID, &group.Name, &group.About, &group.ImageLink)
		group.ImageLink = hostURL + group.ImageLink
		*groups = append(*groups, &group)
	}
	rows.Close()

	return 200, "Successful."
}

func About(groupName string, groupID int32, group *models.Group) (code int, message string) {
	var err error
	if groupID != 0 {
		err = database.QueryRow(`
		SELECT id, name, about, image_link
			FROM groups WHERE id = $1;`, groupID).Scan(
			&group.ID, &group.Name, &group.About, &group.ImageLink)
	} else {
		err = database.QueryRow(`
		SELECT id, name, about, image_link
			FROM groups WHERE name = $1;`, groupName).Scan(
			&group.ID, &group.Name, &group.About, &group.ImageLink)
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

func Search_Ing(groupID int, query string, count int, offset int, ingredients *models.IngredientArr) (code int, message string) {
	rows, err := database.Query(`
		SELECT
			i.id,
			i.name,
			i.danger,
			i.description,
			i.wiki_link,
			coalesce(ig.groups, '{}') AS groups
		FROM ingredients AS i
		JOIN ing_groups AS ig ON i.id = ig.id
		WHERE i.id IN (
			SELECT id FROM ing_groups WHERE  ($1 = ANY (groups))
		) AND i.name LIKE '%' || $2 || '%'
			ORDER BY i.frequency DESC, i.danger DESC LIMIT $3 OFFSET $4
	`, strconv.Itoa(groupID), query, count, offset)

	println(groupID)
	if err == pgx.ErrNoRows {
		return 404, "Group not found."
	} else if err != nil {
		return 500, err.Error()
	}

	for rows.Next() {
		curIng := models.Ingredient{}
		rows.Scan(
			&curIng.ID, &curIng.Name, &curIng.Danger,
			&curIng.Description, &curIng.WikiLink, pq.Array(&curIng.Groups))
		*ingredients = append(*ingredients, &curIng)
	}

	return 200, "Successful."
}

func Ingredients(groupID int, count int, offset int, ingredients *models.IngredientArr) (code int, message string) {
	rows, err := database.Query(`
		SELECT i.id, i.name, i.danger, i.description, i.wiki_link, coalesce(ig.groups, '{}')
			FROM ingredients AS i
			JOIN ing_groups AS ig ON i.id = ig.id
			WHERE i.id IN (
				SELECT id FROM ing_groups WHERE  ($1 = ANY (groups))
			)
				ORDER BY i.frequency DESC, i.danger DESC LIMIT $2 OFFSET $3
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
			&curIng.Description, &curIng.WikiLink, pq.Array(&curIng.Groups))
		*ingredients = append(*ingredients, &curIng)
	}

	return 200, "Successful."
}

func Search(groupsName string, groups *models.GroupArr) (code int, message string) {
	rows, err := database.Query(`
		SELECT id, name, about, image_link
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
		rows.Scan(&curGroup.ID, &curGroup.Name, &curGroup.About, &curGroup.ImageLink)
		curGroup.ImageLink = hostURL + curGroup.ImageLink
		*groups = append(*groups, &curGroup)
	}

	return 200, "Successful."
}
