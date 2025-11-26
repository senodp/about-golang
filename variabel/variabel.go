package main

import "fmt"

func main(){
	//deklarasi variabel
	// var -> perlu menulis tipe datanya
	// := -> tidak perlu menulis tipe datanya
	// uint -> hanya support bilangan positif
	var number int64 = 12345
	number_2 := 12345

	fmt.Println(number_2)

	number_2 = 123

	fmt.Println(number)
	fmt.Println(number_2)

	var yes bool = true
	no := false

	no = true
	fmt.Println(no) 
	fmt.Println(yes)
}