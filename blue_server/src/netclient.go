package blueserver

import (
	"fmt"
	"net"
)

type ExchangeConnector struct {
	host       string
	port       int
	name       string
	connection net.Conn
}

func (e *ExchangeConnector) connectServer() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", e.host, e.port))
	if err != nil {
		fmt.Println("Error al conectarse al servidor:", err)
		return
	}
	e.connection = conn
}

func (e *ExchangeConnector) disconnect() {
	e.connection.Close()
}

func (e *ExchangeConnector) sendMessage(message string) {
	msg := e.name + ":" + message
	_, err := e.connection.Write([]byte(msg))
	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		return
	}
}

func (e *ExchangeConnector) receiveLoop() {
	for {
		buffer := make([]byte, 128)
		numBytes, err := e.connection.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer los datos del servidor:", err)
			e.disconnect()
			return
		}
		fmt.Println("(Net): " + string(buffer[:numBytes]) + " recibido.")
	}
}

func (e *ExchangeConnector) startReceiveThread() {
	go e.receiveLoop()
}
