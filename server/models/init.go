package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func init() {
	os.Chdir("server")
	jsonFile, err := os.Open("categories.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var CategorySlice Categories
	json.Unmarshal(byteValue, &CategorySlice)

	for i := 0; i < len(CategorySlice); i++ {
		fmt.Println("Category Id: 			 " + CategorySlice[i].ID)
		fmt.Println("Category Name: 		 " + CategorySlice[i].Name)
		fmt.Println("Category ListOrder: " + strconv.FormatInt(CategorySlice[i].ListOrder, 10))
	}
}
