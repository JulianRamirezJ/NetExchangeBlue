# NetExchangeBlue

Este proyecto consiste en desarrollar un chat por línea de comandos que permita la comunicación entre diferentes clientes a través de dos protocolos de comunicación: TCP con sockets y Bluetooth. En esencia, consta de dos servidores, uno para la comunicación a través de sockets hecho en Rust y desplegado en una instancia de Compute Engine de GCP, y otro para la comunicación Bluetooth hecho en Go que se conecta con el servidor de sockets a través de sockets. Además, se han desarrollado clientes que pueden conectarse a ambos servidores, permitiendo la comunicación entre múltiples clientes de ambos tipos. Los clientes de sockets se han implementado en Python, mientras que los de Bluetooth se han hecho en Go. El resultado final es un sistema que permite el envío sincrónico de mensajes entre diferentes clientes conectados a cualquiera de los servidores.


A continuación se presenta la arquitectura para la solución implmentada en el proyecto:

![NetExchangeBlue drawio](https://user-images.githubusercontent.com/57159295/236983173-2c0738eb-4000-47ed-b3ee-2120adf8165c.png)

## Como se lamza y usa el proyecto


Por supuesto, aquí te dejo los pasos detallados para lanzar el proyecto:

Clona el repositorio en tu máquina local usando Git:
bash
Copy code
git clone https://github.com/tu_usuario/proyecto_chat.git
Reemplaza "tu_usuario" con tu nombre de usuario en GitHub.

Ejecuta el script de instalación ubicado en la raíz del proyecto para instalar las dependencias necesarias:
bash
Copy code
./install.sh
En una terminal, navega hasta la carpeta "net_rust_server" y lanza el servidor de sockets en Rust con el siguiente comando:
arduino
Copy code
cargo run --release
Esto iniciará el servidor de sockets y lo dejará escuchando en un puerto en tu máquina.

En otra terminal, navega hasta la carpeta "go_bluetooth_server" y lanza el servidor de Bluetooth en Go con el siguiente comando:
go
Copy code
go run server.go
Esto iniciará el servidor de Bluetooth y lo dejará escuchando en un puerto en tu máquina.

Ahora debes verificar la IP y puerto de tu máquina donde se está ejecutando el servidor de sockets de Rust y lanzar tantos clientes web como desees usando Python 3. Abre una nueva terminal y navega hasta la carpeta "python_clients". Luego, ejecuta el siguiente comando:
Copy code
python3 user.py IP_DEL_SERVIDOR PUERTO_DEL_SERVIDOR
Reemplaza "IP_DEL_SERVIDOR" y "PUERTO_DEL_SERVIDOR" con la dirección IP y puerto del servidor de sockets de Rust que se muestra en la terminal donde iniciaste el servidor en el paso 3.

Finalmente, para lanzar los clientes Bluetooth con Go, abre una nueva terminal y navega hasta la carpeta "go_bluetooth_client". Luego, ejecuta el siguiente comando:
go
Copy code
go run client.go
Esto iniciará el cliente Bluetooth y lo conectará al servidor de Bluetooth.

Ahora puedes escribir mensajes en cualquiera de los clientes conectados y todos los demás clientes deberían recibirlos de forma sincronizada.
¡Listo! Ahora puedes disfrutar del chat por línea de comandos que permite la comunicación entre diferentes clientes mediante dos opciones de protocolo de comunicación: TCP con sockets y Bluetooth.

