use std::collections::HashMap;
use std::io::{self, Read, Write};
use std::net::{TcpListener, TcpStream};
use std::sync::{Arc, Mutex};
use std::thread;
use rayon::prelude::*;
use std::sync::mpsc::{Sender, Receiver, channel};

fn send_loop(rx: Receiver<(String,String)>, streams: Arc<Mutex<HashMap<String, TcpStream>>>) -> io::Result<()> {
    loop {
        match rx.recv() {
            Ok((addr, message)) => {
                let mut streams = streams.lock().unwrap();
                /**for (address, stream) in streams.iter_mut() {
                    if *address != addr {
                        stream.write_all(message.as_bytes())?;
                    }
                }**/
                streams.par_iter_mut().for_each(|(address, stream)| {
                    if *address != addr {
                        stream.write_all(message.as_bytes()).unwrap();
                    }
                });  
            }
            Err(_) => continue,
        }
    }
}

fn receive_loop(mut stream: TcpStream, tx: Sender<(String,String)>) -> io::Result<()> 
{
    loop {
        let mut buffer = [0; 1024];
        match stream.read(&mut buffer) {
            Ok(bytes_read) if bytes_read > 0 => {
                let message = String::from_utf8_lossy(&buffer[..bytes_read]).trim().to_string();
                tx.send((stream.peer_addr()?.to_string(), message.clone())).unwrap();
                println!("Received message: {}\n", message);
            }
            Ok(_) => continue,
            Err(_e) => break,
        }
    }
    Ok(())
}

fn main() -> io::Result<()>{
    let (tx, rx): (Sender<(String, String)>, Receiver<(String, String)>) = channel();
    let listener: TcpListener = TcpListener::bind("127.0.0.1:12345")?;
    let streams: Arc<Mutex<HashMap<String, TcpStream>>> = Arc::new(Mutex::new(HashMap::new()));

    let streams_clone = streams.clone(); 
    let send_handle = thread::spawn(move ||  send_loop(rx, streams_clone));

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                let peer_addr = stream.peer_addr()?.to_string();
                println!("New client connected: {}", peer_addr);
                streams.lock().unwrap().insert(peer_addr.clone(), stream.try_clone().unwrap());
                let s_stream = stream.try_clone().unwrap();
                let tx_c = tx.clone();
                let _handle = thread::spawn(|| receive_loop(s_stream, tx_c));
            }
            Err(e) => {
                println!("Error while accepting connection: {}", e);
            }
        }
    }
    send_handle.join().unwrap();
    Ok(())
}
