package main

import (
	blueserver "NetExchangeBlue/blue_server/src"

	"log"
)

func main() {
	blueserver.Run()
	log.Println("Esperando que el servidor esté en línea...")

}
