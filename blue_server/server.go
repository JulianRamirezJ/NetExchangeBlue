package blueserver

import (
	"fmt"

	"github.com/tarm/serial"
)

func Run() {
	// Configura el puerto serie Bluetooth
	config := &serial.Config{Name: "/dev/pts/4", Baud: 9600}
	bluetooth, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println("Error al abrir el puerto Bluetooth:", err)
		return
	}

	// Escucha los datos entrantes y los muestra en la consola
	for {
		buffer := make([]byte, 128)
		numBytes, err := bluetooth.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer los datos del Bluetooth:", err)
			return
		}
		fmt.Println(string(buffer[:numBytes]) + " recibido.")
	}
}
