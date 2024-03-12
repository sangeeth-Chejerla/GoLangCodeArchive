package main

import "fmt"

var steps int = 21
type Message string

func (msg Message) Say() {fmt.Print(msg)}

func main(){
	msg:=Message("I will master go")
	msg.Say()
	fmt.Println(steps,"steps.")
}