package main

import (
	"fmt"
	"github.com/senodp/about-golang/calculation"
)
func main(){
	fmt.Println("Belajar Golang Pertamaku")

	result := calculation.Add(10, 10)
	fmt.Println(result)
}