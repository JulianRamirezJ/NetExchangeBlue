package main

import (
	blueclient "NetExchangeBlue/blue_client/src"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese nombre de usuario: ")
	if scanner.Scan() {
		name := scanner.Text()
		blueclient.Run(name)
	}
}
