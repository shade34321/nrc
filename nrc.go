package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseURL = "https://vpic.nhtsa.dot.gov/api/vehicles/"

func main() {
	fmt.Println("Start NRC")
	GetAllMakes()
}

func GetAllMakes() {
	url := baseURL + "GetAllMakes?format=json"
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Get recall failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
