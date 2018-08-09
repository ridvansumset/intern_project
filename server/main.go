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
			AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
			AllowOrigins:     []string{"*"},
			AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}))

	router.GET("/categories", controllers.ListCategory)
	router.GET("/categories/:category_id", controllers.GetCategory)
	// router.POST("/category", controllers.PostCategory)
	// router.PUT("/category/:id", controllersUpdateUser)
	// router.DELETE("/category/:id", controllers.DeleteUser)

	router.Logger.Fatal(router.Start(":1323"))
}
