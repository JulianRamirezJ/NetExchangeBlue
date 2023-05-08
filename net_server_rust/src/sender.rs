use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use std::io::{self, Write};
use std::sync::mpsc::Receiver;
use std::net::TcpStream;
use rayon::prelude::*;

pub fn send_loop(
    rx: Receiver<(String,String)>,
    streams: Arc<Mutex<HashMap<String, TcpStream>>>
) -> io::Result<()> 
{
    loop {
        match rx.recv() {
            Ok((addr, message)) => {
                let mut streams = streams.lock().unwrap();
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