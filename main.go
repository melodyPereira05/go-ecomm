package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/melodyPereira05/go-ecomm/controllers"
	"github.com/melodyPereira05/go-ecomm/database"
	"github.com/melodyPereira05/go-ecomm/middleware"
	"github.com/melodyPereira05/go-ecomm/routes"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.ButFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
