package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var CategorySlice Categories

func Init() {
	os.Chdir("server")
	jsonFile, err := os.Open("categories.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &CategorySlice)
}
