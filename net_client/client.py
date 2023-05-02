import socket

class ExchangeConnector:

    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.connection = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def connect_server(self):
        self.connection.connect((self.host, self.port))
        return self.connection

    def disconnect(self):
        self.connection.close()

    def send_message(self, message):
        self.connection.sendall(message.encode())

    def receive_message(self):
        response = self.connection.recv(1024).decode()
        print('Server response:', response)
        return response
