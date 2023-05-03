import config
import client

if __name__ == '__main__':
    connector = client.ExchangeConnector(config.HOST, config.PORT)
    connector.connect_server()
    connector.start_receive_thread()

    while True:
        message = input("Type your message: ")
        connector.send_message(message)
