package products

import (
	"backend/internal/database/connection"
	"backend/internal/models"

	"log"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func All(products *models.ProductExtendedArr) (code int, message string) {
	rows, err := database.Query(`
	SELECT
	barcode,
	name,
	description,
    contents,
    category_url,
    mass,
    bestbefore,
    nutrition,
    manufacturer,
	image
	FROM products_extended`)

	if err != nil {
		log.Println("database/products.go: 500, " + err.Error())
		return 500, err.Error()
	}

	for rows.Next() {
		product := models.ProductExtended{}
		rows.Scan(&product.Barcode, &product.Name, &product.Description, &product.Contents, &product.CategoryURL, &product.Mass, &product.BestBefore, &product.Nutrition, &product.Manufacturer, &product.Image)
		*products = append(*products, &product)
	}
	rows.Close()

	return 200, "Successful."
}

func GetOneBarcode(barcode int64, productExt *models.ProductExtended, productShr *models.ProductShrinked, shrinked bool) (code int, message string) {
	err := database.QueryRow(`
	SELECT
	barcode,
	name,
	description,
    contents,
    category_url,
    mass,
    bestbefore,
    nutrition,
    manufacturer,
	image
	FROM products_extended
	WHERE barcode = $1;`, barcode).Scan(&productExt.Barcode, &productExt.Name, &productExt.Description, &productExt.Contents, &productExt.CategoryURL, &productExt.Mass, &productExt.BestBefore, &productExt.Nutrition, &productExt.Manufacturer, &productExt.Image)

	log.Println(err)
	if err == pgx.ErrNoRows {
		errSelect := database.QueryRow(`SELECT barcode, name FROM products WHERE barcode = $1;`, barcode).Scan(&productShr.Barcode, &productShr.Name)

		if errSelect != nil {
			if errSelect == pgx.ErrNoRows {
				log.Println("database/products.go: 404, " + err.Error())
				return 404, "Product not found."
			}
			log.Println("database/products.go: 500, " + err.Error())
			return 500, err.Error()
		}

		shrinked = true
		return 200, "Successful."

	} else if err != nil {
		log.Println("database/products.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return 200, "Successful."
}
