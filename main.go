package main

import (
	blueclient "NetExchangeBlue/blue_client"
	blueserver "NetExchangeBlue/blue_server"

	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	go blueserver.Run()
	log.Println("Esperando que el servidor esté en línea...")
	<-time.After(5 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese nombre de usuario: ")
	if scanner.Scan() {
		name := scanner.Text()
		blueclient.Run(name)
	}

}
