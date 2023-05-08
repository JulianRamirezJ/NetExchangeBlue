package main

import (
	blueclient "NetExchangeBlue/blue_client"
	blueserver "NetExchangeBlue/blue_server"
	"log"
	"time"
)

func main() {
	go blueserver.Run()
	log.Println("Esperando que el servidor esté en línea...")
	<-time.After(5 * time.Second)

	blueclient.Run()
}
