use std::io;
use std::thread;
use std::io::prelude::*;
use std::net::{TcpListener, TcpStream, Shutdown};

fn send_loop(mut stream: TcpStream) -> std::io::Result<()> {
    loop{
        let mut name = String::new();
        io::stdin().read_line(&mut name)
            .expect("Failed to read line");
        stream.write_all(name.as_bytes())?;
    }
}

fn receive_loop(mut stream: TcpStream) -> std::io::Result<()> {
    loop {
        let mut buffer = [0; 1024];
        match stream.read(&mut buffer) {
            Ok(bytes_read) if bytes_read > 0 => {
                let message = String::from_utf8_lossy(&buffer[..bytes_read]).trim().to_string();
                println!("Received message: {}\n", message);
            }
            Ok(_) => continue,
            Err(_e) => break,
        }
    }
    Ok(())
}

fn main() -> std::io::Result<()>{
    let listener: TcpListener = TcpListener::bind("127.0.0.1:12345")?;

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                println!("New client connected: {}", stream.peer_addr()?);
                let s_stream = stream.try_clone().unwrap();
                let handle = thread::spawn(|| receive_loop(stream));
                let send = thread::spawn(|| send_loop(s_stream));
            }
            Err(e) => {
                println!("Error while accepting connection: {}", e);
            }
        }
    }

    Ok(())
}
