package main

import (
	"fmt"
	"reflect"
)

func printTypeInfo(value interface{}) {
  typeOf := reflect.TypeOf(value)
  kind := typeOf.Kind()
  fmt.Println("Type:", typeOf)
  fmt.Println("Kind:", kind)
}

func main() {
  message := "Hello, world!"
  number := 42
  isAwesome := true

  printTypeInfo(message) // Pass string argument
  printTypeInfo(number)   // Pass integer argument
  printTypeInfo(isAwesome) // Pass boolean argument
}

/* 
Type: string
Kind: string
Type: int
Kind: int
Type: bool
Kind: bool
*/
