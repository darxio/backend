package products

import (
	"backend/internal/database/connection"
	"backend/internal/models"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func All(products *models.ProductArr) (code int, message string) {
	rows, err := database.Query("SELECT id, name, barcode FROM products")

	if err != nil {
		return 500, "Something went wrong.."
	}

	for rows.Next() {
		product := models.Product{}
		rows.Scan(&product.ID, &product.Name, &product.Barcode)
		*products = append(*products, &product)
	}
	rows.Close()

	return 200, "Successful."
}

func GetOneBarcode(barcode int, product *models.Product) (code int, message string) {

	err := database.QueryRow("SELECT id, name, barcode FROM products WHERE barcode = $1;", barcode).Scan(&product.ID, &product.Name, &product.Barcode)

	if err == pgx.ErrNoRows {
		return 404, "Product not found."
	} else if err != nil {
		return 500, "Something went wrong.."
	}

	return 200, "Successful."
}
