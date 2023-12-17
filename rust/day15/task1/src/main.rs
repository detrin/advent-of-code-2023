use std::io::{self, BufRead};

fn hash(s: &str) -> u8 {
    let mut current = 0;
    for ch in s.chars() {
        current = ((current as u16 + ch as u16) * 17) % 256;
    }
    current as u8
}

fn main() {
    let stdin = io::stdin();
    let mut line = String::new();
    stdin.lock().read_line(&mut line).unwrap();
    line = line.trim().to_string();

    let steps: Vec<&str> = line.split(",").collect();
    let mut total = 0;
    for step in steps {
        total += hash(step) as u32;
    }

    println!("{}", total);
}
