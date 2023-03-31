package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type DatabaseController struct {
	File     string
	Products []Product
}

func (db *DatabaseController) loadData() error {
	var data []Product
	file, err := os.Open(db.File)

	defer file.Close()

	if err != nil {
		return err
	}

	err = json.NewDecoder(file).Decode(&data)

	if err != nil {
		return err
	}

	db.Products = data

	return nil
}

func (db *DatabaseController) getProducts() []Product {
	return db.Products
}

var Database = DatabaseController{File: "/Users/totorres/Desktop/bootcamp/go/Go web/practica_1/web-server/products.json"}

func main() {

	err := Database.loadData()

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/ping", PingHandler)
	router.GET("/products", GetProductsHandler)
	router.GET("/products/:id", GetProductsByIdHandler)
	router.GET("/products/search", GetProductsByParamHandler)

	router.Run()
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

func GetProductsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"products": Database.getProducts(),
	})
}

func GetProductsByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, product := range Database.getProducts() {
		if product.ID == id {
			c.JSON(200, product)
		}
	}

	c.JSON(404, gin.H{
		"error": "product doesnt found",
	})

}

func GetProductsByParamHandler(c *gin.Context) {
	priceGt := c.Query("priceGt")
	priceGtFloat, _ := strconv.ParseFloat(priceGt, 64)

	for _, product := range Database.getProducts() {
		if product.Price >= priceGtFloat {
			c.JSON(200, product)
		}
	}

	c.JSON(404, gin.H{
		"error": "product doesnt found",
	})
}
