package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/KamillKAPLAN/intern_project/server/controllers"
	"github.com/KamillKAPLAN/intern_project/server/models"

	"github.com/labstack/echo"
)

func init() {
	fmt.Println("Started...")
	os.Chdir("server")
	jsonFile, err := os.Open("categories.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cats models.Categories

	json.Unmarshal(byteValue, &cats)

	for i := 0; i < len(cats); i++ {
		fmt.Println("Category Id: 			 " + cats[i].ID)
		fmt.Println("Category Name: 		 " + cats[i].Name)
		fmt.Println("Category ListOrder: " + strconv.FormatInt(cats[i].ListOrder, 10))
	}
}

func main() {
	router := echo.New()
	// router.GET("/catgry", )
	router.GET("/category/:id", controllers.GetMyCategory)
	// e.POST("/catgry", saveUser)
	// e.PUT("/catgry/:id", updateUser)
	// e.DELETE("/catgry/:id", deleteUser)

	router.Logger.Fatal(router.Start(":1323"))
}
