use std::io::{self, BufRead};
use std::str::FromStr;

#[derive(Clone)]
struct Lens {
    label: String,
    power: i32,
}

fn hash(s: &str) -> usize {
    let mut current = 0;
    for ch in s.chars() {
        current = ((current + ch as u32) * 17) % 256;
    }
    current as usize
}

fn main() {
    let stdin = io::stdin();
    let mut line = String::new();
    stdin.lock().read_line(&mut line).unwrap();
    line = line.trim().to_string();

    let steps: Vec<&str> = line.split(",").collect();
    let mut boxes: Vec<Vec<Lens>> = vec![vec![]; 256];

    for step in steps {
        let parts: Vec<&str> = step.split("=").collect();
        let label = parts[0].replace("-", "");
        let box_index = hash(&label);

        if parts.len() == 1 {
            boxes[box_index].retain(|lens| lens.label != label);
        } else {
            let power = i32::from_str(parts[1]).unwrap();
            let new_lens = Lens {
                label: label.clone(),
                power,
            };

            if let Some(lens) = boxes[box_index].iter_mut().find(|lens| lens.label == label) {
                lens.power = power;
            } else {
                boxes[box_index].push(new_lens);
            }
        }
    }

    let total_power: i32 = boxes.iter().enumerate().map(|(i, box_lenses)| {
        box_lenses.iter().enumerate().map(|(j, lens)| {
            (i + 1) as i32 * (j + 1) as i32 * lens.power
        }).sum::<i32>()
    }).sum();

    println!("{}", total_power);
}
