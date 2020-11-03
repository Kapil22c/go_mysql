package main

import (
	"fmt"

	"github.com/Kapil22c/Go_MySQL/config"
	"github.com/Kapil22c/Go_MySQL/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// fmt.Println("Go MySQL basic")

	// db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	// fmt.Println("Successfully connected to MySQL server")
	Demo1_CallFindAll()
}

func Demo1_CallFindAll() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err := productModel.FindAll()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(products)
			fmt.Println("Product List")
			for _, product := range products {
				fmt.Println("Id:", product.Id)
				fmt.Println("Name:", product.Name)
				fmt.Println("Price:", product.Price)
				fmt.Println("Quantity:", product.Quantity)
				fmt.Println("Status:", product.Status)
				fmt.Println("-------------------------")

			}
		}
	}
}
