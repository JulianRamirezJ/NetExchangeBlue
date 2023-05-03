import socket
import threading

HOST = '127.0.0.1'
PORT = 12345

def receive_loop(conn):
    while True:
        data = conn.recv(1024)
        if not data:
            break
        message = data.decode()
        print(f'Received message: {message}')
        reply = f'Received message: {message}\n'
        conn.sendall(reply.encode())

def send_loop(conn):
    while True:
        message = input('Enter a message: ')
        conn.sendall(message.encode())

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.bind((HOST, PORT))
    s.listen()
    conn, addr = s.accept()
    print('Connected by', addr)

    receive_thread = threading.Thread(target=receive_loop, args=(conn,), daemon=True)
    send_thread = threading.Thread(target=send_loop, args=(conn,), daemon=True)
    receive_thread.start()
    send_thread.start()

    receive_thread.join()
    send_thread.join()
