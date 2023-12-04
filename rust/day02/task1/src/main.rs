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
    let red_limit = 12;
    let green_limit = 13;
    let blue_limit = 14;
    let mut game_id_sum = 0;

    for line in stdin.lock().lines() {
        let line = line.unwrap();
        let parts: Vec<&str> = line.split(": ").collect();
        let game_id = i32::from_str(parts[0].split(" ").nth(1).unwrap_or_default()).unwrap_or_default();
        let sets: Vec<&str> = parts[1].split("; ").collect();
        let mut all_games_possible = true;
        for set in sets {
            let cube = Cube::from_str(set);
            all_games_possible &= cube.red <= red_limit && cube.green <= green_limit && cube.blue <= blue_limit;
        }
        if all_games_possible {
            game_id_sum += game_id;
        }
    }

    println!("{}", game_id_sum);
}
