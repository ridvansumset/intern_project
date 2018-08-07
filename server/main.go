package main

import (
	"controllers"

	"github.com/labstack/echo"
)

func main() {
	router := echo.New()
	// router.GET("/users", )
	router.GET("/category/:id", controllers.GetMyCategory)
	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	router.Logger.Fatal(router.Start(":1323"))
}
