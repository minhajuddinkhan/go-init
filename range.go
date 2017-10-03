package main

import "fmt"

func main() {

	var kv = map[string]string{"foo": "123", "bar": "456"}
	for k, v := range kv {
		fmt.Println(k, v)
	}

	var str = "go"
	for i, c := range str {
		fmt.Println(i, string(c))

	}
}
