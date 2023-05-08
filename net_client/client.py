import socket
import threading
import traceback


class ExchangeConnector:
    def __init__(self, host, port, name):
        self.host = host
        self.port = port
        self.name = name
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
        msg = self.name + ":" + message
        self.connection.sendall(msg.encode())

    def receive_loop(self):
        while True:
            response = self.connection.recv(1024).decode()
            print('\n', response)

    def start_receive_thread(self):
        t = threading.Thread(target=self.receive_loop, daemon=True)
        t.start()