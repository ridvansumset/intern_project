package main

import (
	"github.com/KamillKAPLAN/intern_project/server/controllers"

	"github.com/labstack/echo"
)

func main() {
	router := echo.New()
	// router.GET("/catgry", )
	router.GET("/", controllers.Init)
	router.GET("/category/:id", controllers.GetCategory)
	// e.POST("/catgry", saveUser)
	// e.PUT("/catgry/:id", updateUser)
	// e.DELETE("/catgry/:id", deleteUser)

	router.Logger.Fatal(router.Start(":1323"))
}
