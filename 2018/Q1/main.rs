use std::fs::File;
use std::io::prelude::*;

fn main() {
    println!("Hello!");

    // Read a file
    let mut fp = File::open("ip.txt").expect("File not found!");
    let mut contents = String::new();
    fp.read_to_string(&mut contents)
        .expect("something went wrong reading the file");
    println!("{}", contents);
}
