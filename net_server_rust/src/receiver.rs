use std::io::{self, Read};
use std::sync::mpsc::Sender;
use std::net::TcpStream;    

pub fn receive_loop(
    mut stream: TcpStream, 
    tx: Sender<(String,String)>
) -> io::Result<()> 
{
    loop {
        let mut buffer = [0; 1024];
        match stream.read(&mut buffer) {
            Ok(bytes_read) if bytes_read > 0 => {
                let message = String::from_utf8_lossy(&buffer[..bytes_read]).trim().to_string();
                tx.send((stream.peer_addr()?.to_string(), message.clone())).unwrap();
                println!("Received message: {}", message);
            }
            Ok(_) => continue,
            Err(_e) => break,
        }
    }
    Ok(())
}