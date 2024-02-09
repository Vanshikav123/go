package main

import (
	"encoding/json"
	"fmt"
)

// javascript object notation
// json.Marshal to get back data in json form
func main() {
	var wrapData map[string]interface{} //using interface allows it to take any type of data

	data := `{"name": "vanshika","designation":"creator"}`
	err := json.Unmarshal([]byte(data), &wrapData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(wrapData)
	//fmt.Println(wrapData["name"])

	val, ok := wrapData["designation"]

	if ok {
		fmt.Println(val)
	}
	/*bytesData,err:=json.Marshal(&wrapData)
	if err!=nil {
		fmt.Println(err)
		}
		fmt.Println(string(bytesData))*/
}
