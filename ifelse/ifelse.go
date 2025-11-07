package main

import "fmt"

func main(){
	nilai := 30
	var rank string

	if nilai >= 90{
		rank = "A"
	}else if nilai >=80{
		rank = "B"
	}else if nilai >=70{
		rank = "C"
	}else if nilai >=60{
		rank = "D"
	}else if nilai >=40{
		rank = "E"
	}else {
		rank = "DANGER!!!"
	}

	fmt.Println(rank)
}