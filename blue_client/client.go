package blueclient

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/tarm/serial"
)

func Run() {
	// Configura el puerto serie Bluetooth
	config := &serial.Config{Name: "/dev/pts/3", Baud: 9600}
	bluetooth, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println("Error al abrir el puerto Bluetooth:", err)
		return
	}

	// Envia un mensaje al servidor
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Escriba su mensaje: ")
		if scanner.Scan() {
			message := scanner.Text()
			_, err = bluetooth.Write([]byte(message))
			if err != nil {
				fmt.Println("Error al enviar los datos al Bluetooth:", err)
				return
			}
		} else {
			fmt.Println("Error al leer la entrada del usuario")
		}
		<-time.After(1 * time.Second)
	}

}
