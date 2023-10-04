package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	connectionString := "localhost:5000"
	router := gin.Default()
	router.GET("/products", getProducts)
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
