use std::io::{self, BufRead};
use std::str::FromStr;

#[derive(Default)]
struct Cube {
    red: i32,
    green: i32,
    blue: i32,
}

impl Cube {
    fn from_str(s: &str) -> Cube {
        let mut cube = Cube::default();
        for part in s.split(", ") {
            let vals: Vec<&str> = part.split(" ").collect();
            let num = i32::from_str(vals[0]).unwrap_or_default();
            match vals[1] {
                "red" => cube.red += num,
                "green" => cube.green += num,
                "blue" => cube.blue += num,
                _ => {},
            }
        }
        cube
    }
}

fn main() {
    let stdin = io::stdin();
    let mut game_mult_sum = 0;

    for line in stdin.lock().lines() {
        let line = line.unwrap();
        let parts: Vec<&str> = line.split(": ").collect();
        let sets: Vec<&str> = parts[1].split("; ").collect();
        let mut red_needed = 0;
        let mut green_needed = 0;
        let mut blue_needed = 0;
        for set in sets {
            let cube = Cube::from_str(set);
            red_needed = red_needed.max(cube.red);
            green_needed = green_needed.max(cube.green);
            blue_needed = blue_needed.max(cube.blue);
        }
        game_mult_sum += red_needed * green_needed * blue_needed;
    }

    println!("{}", game_mult_sum);
}
