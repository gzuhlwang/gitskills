package main

import "fmt"

func main(){
	fmt.Print(add(1,2))
}

func add(a , b int)(c int){
	return a+b
}