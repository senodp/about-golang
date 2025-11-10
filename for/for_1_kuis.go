package main

import "fmt"

func main(){
	//menampilkan data yang indexnya angka genal saja
	namaAnak := "Ali Miran Hafidzuan"
	for index, next := range namaAnak {
		if index % 2 == 0{
			fmt.Println("index :", index, " next :", string(next))
		}
	}
	
	fmt.Println("==============================")

	//menampilkan data yang huruf vocal saja
	namaKu := "alfiani Syafitri"
	for indexing, lettering := range namaKu{
		letterString := string(lettering)

		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o"{
			fmt.Println("indexing :", indexing, " lettering :", string(lettering))
		} 
	}
}