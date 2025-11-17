package main

import "fmt"

func main(){
	//deklarasi 1
	var arr [4]int

	//deklarasi 2 inisiasi langsung
	arr2 := [4]string{"Seno", "Dwi", "Prasetyo", "Lorem"}

	//mengisi array
	arr[0] = 190
	arr[3] = 80

	fmt.Println(arr)
	//akses ke array tertentu
	fmt.Println(arr[3] + arr[0])

	//akses deklarasi 2
	fmt.Println(arr2)

	//pemanggilan dengan for
	for i := 0; i < len(arr2); i++{ //length
		fmt.Println(arr2[i])
	} 
}