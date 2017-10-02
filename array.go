package main

import (
	"fmt"
);


var array []string;

func arr() {
	
	array = appendValue(array, "string");
	array = appendValue(array, "second");
	
	fmt.Println(array)

	
	
}

func appendValue (param []string , value string) []string {
	param = append(param, value);	
	return param;
}	

func sum([]int ) int {

	
	return 0;
}
