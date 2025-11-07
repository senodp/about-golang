package main

import "fmt"

func main(){
	number := "Ahad"

	switch number {
	case "Senin":
		fmt.Println("Ini adalah hari senin")
	case "Selasa":
		fmt.Println("Ini adalah hari selasa")
	case "Rabu":
		fmt.Println("Ini adalah hari rabu")
	case "Kamis":
		fmt.Println("Ini adalah hari kamis") 
	default:
		fmt.Println("Tidak ada di daftar hari")
	}
}