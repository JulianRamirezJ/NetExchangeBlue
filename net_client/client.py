import socket
import threading
import traceback


class ExchangeConnector:
    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.connection = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def connect_server(self):
        try:
            self.connection.connect((self.host, self.port))
        except Exception as e:
            print(f"Error: {e}")
            traceback.print_exc()
    

    def disconnect(self):
        self.connection.close()

    def send_message(self, message):
        self.connection.sendall(message.encode())

    def receive_loop(self):
        while True:
            response = self.connection.recv(1024).decode()
            print('Server response:', response)

    def start_receive_thread(self):
        t = threading.Thread(target=self.receive_loop, daemon=True)
        t.start()