package main

import (
	"fmt"
	"reflect"
)

func main(){
	const pi = 3.141592
	fmt.Println("type:",reflect.TypeOf(pi))
}