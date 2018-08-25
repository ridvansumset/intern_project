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

var OptionSlice Options

var ChoiceSlice Choices

func Init() {
	os.Chdir("server")
	// categories.json
	jsonFile, err := os.Open("categories.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &CategorySlice)
	// products.json
	jsonFile2, err2 := os.Open("products.json")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer jsonFile2.Close()
	byteValue2, _ := ioutil.ReadAll(jsonFile2)
	json.Unmarshal(byteValue2, &ProductSlice)
	// option.json
	jsonFile3, err3 := os.Open("option.json")
	if err3 != nil {
		fmt.Println(err3)
	}
	defer jsonFile3.Close()
	byteValue3, _ := ioutil.ReadAll(jsonFile3)
	json.Unmarshal(byteValue3, &OptionSlice)
	// choice.json
	jsonFile4, err4 := os.Open("choice.json")
	if err4 != nil {
		fmt.Println(err4)
	}
	defer jsonFile4.Close()
	byteValue4, _ := ioutil.ReadAll(jsonFile4)
	json.Unmarshal(byteValue4, &ChoiceSlice)
}
