package main

import "fmt"

func main(){
	//dengan cara for biasa
	for i := 1; i <= 10; i++ {
		fmt.Println("Seno Bermain Bola :", i)
	}

	//dengan cara seperti while
	a := 1;
	for a <= 5{
		fmt.Println("Urutan angka :", a)
		a++
	}

	//dengan cara seperti foreach
	nama := "Seno Dwi Prasetyo"
	// for indexing, next := range nama{
	// 	fmt.Println("index :", indexing, " letter :", string(next))
	// }
	//gunakan tanda _ untuk membuat variabel tidak dibaca atau digunakan
	for _, next := range nama{
		fmt.Println("letter :", string(next))
	}
}