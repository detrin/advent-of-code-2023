use std::io::{self, BufRead};
use std::collections::HashMap;

#[derive(Debug)]
struct Converter {
    from: String,
    to: String,
    map: Vec<Vec<u64>>,
}

fn parse_seeds(seeds_line: &str) -> Vec<u64> {
    seeds_line.trim_start_matches("seeds: ")
        .split_whitespace()
        .map(|s| s.parse::<u64>().unwrap())
        .collect()
}

fn parse_converter(input: &str) -> Converter {
    let lines: Vec<&str> = input.split('\n').collect();
    let from_to: Vec<&str> = lines[0].split("-to-").collect();
    let from = from_to[0].trim().to_string();
    let to = from_to[1].split_whitespace().next().unwrap().to_string();

    let mut maps = vec![];
    for line in &lines[1..] {
        let nums: Vec<u64> = line.split_whitespace()
            .map(|num| num.parse().unwrap())
            .collect();
        maps.push(nums);
    }

    Converter {
        from: from,
        to: to,
        map: maps,
    }
}

fn parse_input(input: &str) -> (Vec<u64>, Vec<Converter>) {
    let lines: Vec<&str> = input.split('\n').collect();
    let seeds = parse_seeds(lines[0]);
    let binding = lines[2..].join("\n");
    let converter_strings: Vec<&str> = binding.split("\n\n").collect();
    let mut converters = vec![];
    for cs in converter_strings {
        converters.push(parse_converter(cs));
    }
    (seeds, converters)
}

fn convert(n: u64, converter: &Converter) -> u64 {
    for mapping in &converter.map {
        let start_dst = mapping[0];
        let start_src = mapping[1];
        let length = mapping[2];
        if n >= start_src && n < start_src + length {
            return start_dst + (n - start_src);
        }
    }
    n
}

fn main() {
    let stdin = io::stdin();
    let mut text = vec![];

    for line in stdin.lock().lines() {
        text.push(line.unwrap());
    }

    let input = text.join("\n");
    let (mut seeds_store1, converters) = parse_input(&input);
    let mut seeds_store2 = vec![];

    let mut converter_map = HashMap::new();
    for converter in &converters {
        converter_map.insert(&*converter.from, converter);
    }

    let mut state = "seed";
    let mut is_converter_available = true;
    while is_converter_available {
        if let Some(converter) = converter_map.get(state) {
            seeds_store2 = vec![];
            for seed in &seeds_store1 {
                seeds_store2.push(convert(*seed, converter));
            }
            std::mem::swap(&mut seeds_store1, &mut seeds_store2);
            state = &converter.to;  
        } else {
            is_converter_available = false;
        }
    }

    let seed_min = *seeds_store1.iter().min().unwrap();
    println!("{}", seed_min);
}
