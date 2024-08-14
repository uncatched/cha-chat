use core::str;
use std::io::Write;
use std::str::FromStr;
use std::sync::mpsc;
use std::thread;
use std::{io::Read, net::TcpListener, net::TcpStream};
use uuid::Uuid;

fn main() {
    start_listener();
}

fn start_listener() {
    let listener = TcpListener::bind("192.168.0.111:6969").expect("Failed to bind listener");
    listener
        .set_nonblocking(true)
        .expect("Failed to sten non-blocking listener");

    let (sender, receiver) = mpsc::channel::<String>();
    loop {
        if let Ok((mut stream, addr)) = listener.accept() {
            println!("Client connected: {}", addr);

            thread::spawn(move || loop {
                let mut buf = vec![0; 32];
                match stream.read(&mut buf) {
                    Ok(_) => {
                        let msg = buf.into_iter().take_while(|&x| x != 0).collect::<Vec<_>>();
                        let msg = String::from_utf8(msg).expect("Invalid UTF8 message");
                        let components: Vec<&str> = msg.trim().split(":").collect();
                        let name = components[0];
                        let message = components[1].trim();

                        match Command::new(message) {
                            Some(cmd) => {
                                handle(&cmd);
                            }
                            None => {
                                println!("{name}: {message}");
                            }
                        }
                    }
                    Err(e) => {
                        // println!("Failed to read: {e}");
                    }
                }
            });
        }
    }
}

fn handle(cmd: &Command) {
    match cmd {
        Command::Help => println!("----help----"),
    }
}

enum Command {
    Help,
}

impl Command {
    fn new(raw_value: &str) -> Option<Command> {
        match raw_value {
            "/help" => Some(Command::Help),
            _ => None,
        }
    }
}
