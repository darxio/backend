package fruits

import (
	"backend/internal/database/connection"
	"backend/internal/models"

	"github.com/jackc/pgx"
	"github.com/lib/pq"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func GetFruitsInfo(name string, f *models.Fruit) (code int, message string) {
	err := database.QueryRow(`
	SELECT
		id,
		COALESCE(name, 'NULL') AS name,
		COALESCE(name_ru, 'NULL') AS name_ru,
		COALESCE(image, 'NULL') AS image,
		groups,
		COALESCE(description, 'NULL') AS description,
		nutrition_labels,
		nutrition,
		vitamins_labels,
		vitamins
		FROM fruits_info
		WHERE name = $1;
		`, name).Scan(&f.ID, &f.Name, &f.NameRu, &f.Image, pq.Array(&f.Groups), &f.Description, pq.Array(&f.NutritionLabels),
		pq.Array(&f.Nutrition), pq.Array(&f.VitaminsLabels), pq.Array(&f.Vitamins))

	if err == pgx.ErrNoRows {
		return 404, "Fruit not found."
	} else if err != nil {
		return 500, err.Error()
	}
	return 200, "Successful."
}
