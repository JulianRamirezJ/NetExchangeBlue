use std::collections::HashMap;
use std::io;
use std::net::{TcpListener, TcpStream};
use std::sync::{Arc, Mutex};
use std::thread;
use std::sync::mpsc::{Sender, Receiver, channel};

mod receiver;
mod sender;
use receiver::receive_loop;
use sender::send_loop;

fn handle_connection(
    stream: TcpStream, 
    tx: Sender<(String, String)>, 
    streams: Arc<Mutex<HashMap<String, TcpStream>>>,
) 
{
    let peer_addr = stream.peer_addr().unwrap().to_string();
    println!("New client connected: {}", peer_addr);
    streams.lock().unwrap().insert(peer_addr.clone(), stream.try_clone().unwrap());
    let s_stream = stream.try_clone().unwrap();
    let tx_c = tx.clone();
    let _handle = thread::spawn(|| receive_loop(s_stream, tx_c));
}

fn main() -> io::Result<()>
{

    let (tx, rx): (Sender<(String, String)>, Receiver<(String, String)>) = channel();
    let listener: TcpListener = TcpListener::bind("127.0.0.1:12345")?;
    let streams: Arc<Mutex<HashMap<String, TcpStream>>> = Arc::new(Mutex::new(HashMap::new()));

    let streams_clone = streams.clone(); 
    let send_handle = thread::spawn(move ||  send_loop(rx, streams_clone));

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                handle_connection(stream, tx.clone(), streams.clone());
            }
            Err(e) => {
                println!("Error while accepting connection: {}", e);
            }
        }
    }
    send_handle.join().unwrap();
    Ok(())
}