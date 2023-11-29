package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// create map struct and get datas from user
	var name, addres string
	fmt.Scan(&name)
	fmt.Scan(&addres)

	idMap := map[string]string{
		"name":    name,
		"address": addres,
	}

	jsonMap, _ := json.Marshal(idMap)

	fmt.Println(string(jsonMap))

}
