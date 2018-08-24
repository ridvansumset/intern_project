package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type databaseQuery map[string]interface{}

var CategorySlice Categories

var ProductSlice Products

func Init() {
	os.Chdir("server")

	jsonFile, err := os.Open("categories.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &CategorySlice)

	jsonFile2, err2 := os.Open("products.json")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer jsonFile2.Close()
	byteValue2, _ := ioutil.ReadAll(jsonFile2)
	json.Unmarshal(byteValue2, &ProductSlice)
}
