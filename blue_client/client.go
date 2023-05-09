package blueclient

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/tarm/serial"
)

func Run(name string) {
	// Configura el puerto serie Bluetooth
	client_name := name
	port_config := &serial.Config{Name: "/dev/pts/1", Baud: 9600}
	bluetooth, err := serial.OpenPort(port_config)
	if err != nil {
		fmt.Println("Error al abrir el puerto Bluetooth:", err)
		return
	}

	// Envia un mensaje al servidor
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("(%s) Escriba su mensaje: ", client_name)
		if scanner.Scan() {
			input := scanner.Text()
			message := client_name + "::" + input
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
