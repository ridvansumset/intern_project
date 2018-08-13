package main

import (
	"github.com/KamillKAPLAN/intern_project/server/controllers"

	"github.com/labstack/echo"
	echomiddleware "github.com/labstack/echo/middleware"
)

func main() {
	router := echo.New()
	controllers.Init()

	router.Pre(echomiddleware.RemoveTrailingSlash())
	router.Use(echomiddleware.Logger())
	router.Use(echomiddleware.Recover())
	router.Use(echomiddleware.Gzip())

	router.Use(
		echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
			AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", echo.DELETE, "OPTIONS"},
			AllowOrigins:     []string{"*"},
			AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}))

	router.GET("/categories", controllers.ListCategories)
	router.GET("/categories/:category_id", controllers.GetCategory)
	router.POST("/categories", controllers.CreateCategory)
	router.DELETE("/categories/:category_id", controllers.DeleteCategory)
	// router.PUT("/categories/:category_id", controllers.UpdateCategory)

	router.Logger.Fatal(router.Start(":1323"))
}
