package blueserver

import (
	"fmt"

	"github.com/tarm/serial"
)

func Run() {

	config := &serial.Config{Name: BLUEPORT, Baud: 9600}
	bluetooth, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println("Error al abrir el puerto Bluetooth:", err)
		return
	}

	messages := make(chan []byte)

	exchange := &ExchangeConnector{host: HOST, port: PORT}
	exchange.connectServer()
	exchange.startReceiveThread()

	// Escucha los datos entrantes
	go func() {
		for {
			buffer := make([]byte, 128)
			numBytes, err := bluetooth.Read(buffer)
			if err != nil {
				fmt.Println("Error al leer los datos del Bluetooth:", err)
				continue
			}

			fmt.Println("(Server): " + string(buffer[:numBytes]) + " recibido.")
			exchange.sendMessage(string(buffer[:numBytes]))
			messages <- append(buffer[:numBytes])
		}
	}()

	fmt.Println("Servidor iniciado. Esperando conexiones...")

	for {
		message := <-messages
		_, err = bluetooth.Write([]byte(message))
		if err != nil {
			fmt.Println("Error al enviar los datos al Bluetooth:", err)
			return
		}
	}
}
