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
	port_config := &serial.Config{Name: BLUEPORT, Baud: 9600}
	bluetooth, err := serial.OpenPort(port_config)
	if err != nil {
		fmt.Println("Error al abrir el puerto Bluetooth:", err)
		return
	}

	go func() {
		for {
			buffer := make([]byte, 128)
			numBytes, err := bluetooth.Read(buffer)
			if err != nil {
				fmt.Println("Error al leer los datos del Bluetooth:", err)
				continue
			}

			// Enviar el mensaje a todos los clientes conectados
			fmt.Println("\n(Server): " + string(buffer[:numBytes]) + " recibido.")
		}
	}()

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
