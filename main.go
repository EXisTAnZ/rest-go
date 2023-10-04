package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	connectionString := "localhost:5000"
	router := gin.Default()
	router.Static("/images", "./assets/images/")
	router.GET("/products", getProducts)
	router.POST("/products", addProduct)
	router.GET("/products/:id", getProduct)
	router.PATCH("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)
	router.Run(connectionString)
	fmt.Println("Server started at " + connectionString + " port")
}

type product struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Count       int     `json:"count"`
}

var products = []product{
	{Id: "1", Title: "Очарованье осени", Description: "Данный шедевр навеян маминым борщом", Price: 1500.00, Count: 3},
	{Id: "2", Title: "Нежный шепот", Description: "Модному явлению asmr посвящается", Price: 2500.00, Count: 2},
	{Id: "3", Title: "Хронический синусит", Description: "Этот аромат пробудит спящие рецепторы в вашем носу", Price: 1300.00, Count: 1},
}

func getProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, products)
}

func getProduct(context *gin.Context) {
	id := context.Param("id")
	if product, err := getProductById(id); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		context.IndentedJSON(http.StatusOK, product)
	}
}

func getProductById(id string) (*product, error) {
	for idx, product := range products {
		if product.Id == id {
			return &products[idx], nil
		}
	}
	return nil, errors.New("product not found")
}

func addProduct(context *gin.Context) {
	var newProduct product
	if err := context.BindJSON(&newProduct); err != nil {
		return
	}
	products = append(products, newProduct)
	context.IndentedJSON(http.StatusCreated, newProduct)
}

func updateProduct(context *gin.Context) {
	id := context.Param("id")
	if product, err := getProductById(id); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		if err := context.BindJSON(product); err != nil {
			return
		}
		context.IndentedJSON(http.StatusOK, product)
	}
}

func deleteProduct(context *gin.Context) {
	id := context.Param("id")
	if product, err := getProductById(id); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		deleteProductById(id)
		context.IndentedJSON(http.StatusNoContent, product)
	}
}

func deleteProductById(id string) {
	for idx, product := range products {
		if product.Id == id {
			products[idx] = products[len(products)-1]
			products = products[:len(products)-1]
		}
	}
}
