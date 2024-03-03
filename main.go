package main

// import (
// 	"github.com/fellgar246/go-ecommerce-yt/controllers"
// 	"github.com/fellgar246/go-ecommerce-yt/database"
// 	"github.com/fellgar246/go-ecommerce-yt/middleware"
// 	"github.com/fellgar246/go-ecommerce-yt/routes"
// 	"github.com/gin-gonic/gin"
// )

import (
	"log"
	"os"

	"github.com/fellgar246/go-ecommerce-yt/controllers"
	"github.com/fellgar246/go-ecommerce-yt/database"
	"github.com/fellgar246/go-ecommerce-yt/middleware"
	"github.com/fellgar246/go-ecommerce-yt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
