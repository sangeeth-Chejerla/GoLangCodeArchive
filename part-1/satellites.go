package main

import "fmt"

func main() {
	//string
	name :="Neptune"
	description := "Planet"
	// integer
	radius:= 24764
	// float
	mass := 1.024e26
	//boolean
	active := true
	//array  
	satellites :=[]string{
		"Naiad", "Thalassa", "Despina", "Galatea", "Larissa", "S/2004 N 1", "Proteus",
	}



fmt.Println(name)
fmt.Println(description)
fmt.Println(radius)
fmt.Println(mass)
fmt.Println(active)
fmt.Println(satellites)
}
