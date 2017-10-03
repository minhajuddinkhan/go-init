package main

import (
	"fmt"
)

func main() {

	var maps = make(map[string]map[string]int)

	var key = make(map[string]int)
	key["myKey"] = 1
	maps["key"] = key
	maps = deleteKey(maps, "key")
	printDict(maps)

}

func printDict(dict map[string]map[string]int) {

	fmt.Println(dict)
}

func deleteKey(dict map[string]map[string]int, key string) map[string]map[string]int {
	delete(dict, key)
	return dict
}
