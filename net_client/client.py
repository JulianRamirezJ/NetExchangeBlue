import socket
import config

client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client_socket.connect((config.HOST, config.PORT))

message = 'Hello, server!'
client_socket.sendall(message.encode())

response = client_socket.recv(1024).decode()
print('Server response:', response)

client_socket.close()
