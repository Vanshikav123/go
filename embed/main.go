package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//var t []byte
//go embed test.json

var d []byte

//deals with static files

type Data struct {
	Name  string
	Phone string
}

func main() {
	fmt.Println(d)
	var dd Data
	json.Unmarshal(d, &dd)
	fmt.Println(dd.Phone)
}

//these data are taken from the files and there binaries are created to maintain the confidentiality
