package blueserver

import (
	"fmt"
	"net"

	"github.com/tarm/serial"
)

func Run() {
	// Configura el puerto serie Bluetooth
	config := &serial.Config{Name: "/dev/pts/2", Baud: 9600}
	bluetooth, err := serial.OpenPort(config)
	if err != nil {
		fmt.Println("Error al abrir el puerto Bluetooth:", err)
		return
	}

	// Crear un canal para nuevas conexiones entrantes
	newConnections := make(chan net.Conn)

	// Crear un canal para mensajes entrantes de cada conexión
	messages := make(chan []byte)

	// Escucha los datos entrantes y los envía a todos los clientes conectados
	go func() {
		for {
			buffer := make([]byte, 128)
			numBytes, err := bluetooth.Read(buffer)
			if err != nil {
				fmt.Println("Error al leer los datos del Bluetooth:", err)
				continue
			}
			// Enviar el mensaje a todos los clientes conectados
			fmt.Println("(Server): " + string(buffer[:numBytes]) + " recibido.")
			messages <- append(buffer[:numBytes], '\n')
		}
	}()

	// Aceptar nuevas conexiones
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor iniciado. Esperando conexiones...")

	// Mantener un mapa de clientes conectados
	clients := make(map[net.Conn]bool)

	// Escuchar nuevas conexiones entrantes y procesar mensajes entrantes
	for {
		select {
		case conn := <-newConnections:
			// Agregar nuevo cliente al mapa
			clients[conn] = true
			fmt.Println("Nuevo cliente conectado:", conn.RemoteAddr())
			// Leer los mensajes entrantes del cliente
			go func(conn net.Conn) {
				defer func() {
					fmt.Println("Cerrando conexión:", conn.RemoteAddr())
					conn.Close()
					delete(clients, conn)
				}()
				for {
					buffer := make([]byte, 128)
					numBytes, err := conn.Read(buffer)
					if err != nil {
						fmt.Println("Error al leer los datos del cliente:", err)
						return
					}
					messages <- append(buffer[:numBytes], '\n')
				}
			}(conn)

		case message := <-messages:
			// Enviar el mensaje a todos los clientes conectados
			for conn := range clients {
				_, err := conn.Write(message)
				if err != nil {
					fmt.Println("Error al enviar mensaje al cliente:", err)
					delete(clients, conn)
				}
			}
		default:
			// Aceptar nuevas conexiones entrantes
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error al aceptar la conexión:", err)
				continue
			}
			newConnections <- conn
		}
	}
}
