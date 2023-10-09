package main

import (
	"errors"
	"flag"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "5000", "http port to start server")
	host := flag.String("host", "localhost", " dns or ip adress of server")
	flag.Parse()
	connectionString := *host + ":" + *port
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	router.Use(cors.New(corsConfig))
	router.Static("/images", "./assets/images/")
	router.GET("/products", getProducts)
	router.POST("/products", addProduct)
	router.GET("/products/:id", getProduct)
	router.PATCH("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)
	router.Run(connectionString)
}

type product struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Count       int     `json:"count"`
	Image       string  `json:"image"`
}

var products = []product{
	{
		Id:          "1",
		Title:       "Очарованье осени",
		Description: "Шедевр, навеянный маминым борщом.",
		Price:       1500.00,
		Count:       3,
		Image:       "http://localhost:5000/images/flower_1.jpg",
	},
	{
		Id:          "2",
		Title:       "Нежный шепот",
		Description: "Модному явлению asmr посвящается.",
		Price:       2500.00,
		Count:       2,
		Image:       "http://localhost:5000/images/flower_2.jpg",
	},
	{
		Id:          "3",
		Title:       "Хронический синусит",
		Description: "Этот аромат пробудит спящие рецепторы в вашем носу.",
		Price:       1300.00,
		Count:       1,
		Image:       "http://localhost:5000/images/flower_3.jpg",
	},
	{
		Id:          "4",
		Title:       "Дон дракон вилсон",
		Description: "Зелёный дракон навеет спокойствие и умер от варения.",
		Price:       900.00,
		Count:       20,
		Image:       "http://localhost:5000/images/flower_4.jpg",
	},
	{
		Id:          "5",
		Title:       "Небесная стая",
		Description: "Стая зеленых драконов может очень долго радовать ваше изменненное сознание.",
		Price:       18000.00,
		Count:       1,
		Image:       "http://localhost:5000/images/flower_5.jpg",
	},
	{
		Id:          "6",
		Title:       "Экзист о нот экзист",
		Description: "Здесь просто пустой текст о красоте цветов для статистики.",
		Price:       10300.00,
		Count:       1,
		Image:       "http://localhost:5000/images/flower_6.jpg",
	},
	{
		Id:          "7",
		Title:       "Гвоздики",
		Description: "Это не те гвоздики, которые молоточком в стенку забивают.",
		Price:       13000.00,
		Count:       2,
		Image:       "http://localhost:5000/images/flower_7.jpg",
	},
	{
		Id:          "8",
		Title:       "Зефирные облака",
		Description: "Кирпичи для того, чтобы построить воздушные замки любви",
		Price:       3400.00,
		Count:       1,
		Image:       "http://localhost:5000/images/flower_8.jpg",
	},
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
