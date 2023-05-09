# NetExchangeBlue

Este proyecto consiste en desarrollar un chat por línea de comandos que permita la comunicación entre diferentes clientes a través de dos protocolos de comunicación: TCP con sockets y Bluetooth. En esencia, consta de dos servidores, uno para la comunicación a través de sockets hecho en Rust y desplegado en una instancia de Compute Engine de GCP, y otro para la comunicación Bluetooth hecho en Go que se conecta con el servidor de sockets a través de sockets. Además, se han desarrollado clientes que pueden conectarse a ambos servidores, permitiendo la comunicación entre múltiples clientes de ambos tipos. Los clientes de sockets se han implementado en Python, mientras que los de Bluetooth se han hecho en Go. El resultado final es un sistema que permite el envío sincrónico de mensajes entre diferentes clientes conectados a cualquiera de los servidores.


A continuación se presenta la arquitectura para la solución implmentada en el proyecto:

![NetExchangeBlue drawio](https://user-images.githubusercontent.com/57159295/236983173-2c0738eb-4000-47ed-b3ee-2120adf8165c.png)

