package main

import (
	"fmt"
	"time"
)

func SendWelcomeNotification(username string){
	fmt.Printf("Hi %s, Welcome to the App\n", username)

	time.Sleep(3 * time.Second) //3 detik

	fmt.Println("[Logging] Notification successfully sent")
}

func main(){
	SendWelcomeNotification("Seno Dwi Prasetyo")
	SendWelcomeNotification("Alfiani Syafitri")
}