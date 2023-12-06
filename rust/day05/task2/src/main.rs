use std::str::FromStr;
use std::collections::HashMap;
use std::cmp::{min, max};
use std::io::{self, BufRead}; 

#[derive(Clone, PartialEq, Eq, Hash, Copy)]
struct SeedRange {
    start: u64,
    range: u64,
}

#[derive(Clone)]
struct Converter {
    from: String,
    to: String,
    map: Vec<Vec<u64>>,
}

fn parse_seeds(seeds_line: &str) -> Option<Vec<SeedRange>> {
    let seed_strings: Vec<&str> = seeds_line.trim_start_matches("seeds:").split_whitespace().collect();
    if seed_strings.len() % 2 != 0 {
        return None;
    }

    let mut seeds = Vec::new();
    for i in (0..seed_strings.len()).step_by(2) {
        let start = u64::from_str(seed_strings[i]).unwrap();
        let range = u64::from_str(seed_strings[i+1]).unwrap();
        seeds.push(SeedRange { start, range });
    }
    Some(seeds)
}

fn parse_converter(input: &str) -> Converter {
    let lines: Vec<&str> = input.split('\n').collect();
    let from_to: Vec<&str> = lines[0].split(" map:").collect::<Vec<&str>>()[0].split("-to-").collect();

    let from = from_to[0].trim().to_string();
    let to = from_to[1].trim().to_string();

    let mut map = Vec::new();
    for line in &lines[1..] {
        let nums: Vec<u64> = line.split_whitespace().map(|num| u64::from_str(num).unwrap()).collect();
        map.push(nums);
    }

    Converter { from, to, map }
}

fn parse_input(input: &str) -> (Option<Vec<SeedRange>>, Vec<Converter>) {  
    let lines: Vec<&str> = input.split('\n').collect();  
    let seeds = parse_seeds(lines[0]);  
    let binding = lines[2..].join("\n");
    let converter_strings: Vec<&str> = binding.split("\n\n").collect();  
    let mut converters = Vec::new();  
    for cs in converter_strings {  
        converters.push(parse_converter(cs));  
    }  
    (seeds, converters)  
} 

fn convert(seed_range: SeedRange, converter: Converter) -> Vec<SeedRange> {
    let mut new_seed_ranges = Vec::new();
    let mut remaining_range = SeedRange { start: seed_range.start, range: seed_range.range };

    for mapping in converter.map {
        let (start_dst, start_src, length) = (mapping[0], mapping[1], mapping[2]);
        let end_src = start_src + length;

        if remaining_range.start < end_src && remaining_range.start+remaining_range.range > start_src {
            let intersect_start = max(remaining_range.start, start_src);
            let intersect_end = min(remaining_range.start + remaining_range.range, end_src);

            let new_start = start_dst + (intersect_start - start_src);
            let new_end = start_dst + (intersect_end - start_src);
            if new_start < new_end {
                new_seed_ranges.push(SeedRange { start: new_start, range: new_end - new_start });
            }

            if remaining_range.start == intersect_start {
                remaining_range.start = intersect_end;
                remaining_range.range -= intersect_end - intersect_start;
            } else if remaining_range.start + remaining_range.range == intersect_end {
                remaining_range.range = intersect_start - remaining_range.start;
            }
        }
    }

    if remaining_range.range > 0 {
        new_seed_ranges.push(remaining_range);
    }

    new_seed_ranges
}

fn merge_seed_ranges(seed_ranges: Vec<SeedRange>) -> Vec<SeedRange> {
    let mut seen = HashMap::new();
    let mut unique_seed_ranges = Vec::new();
    for seed_range in seed_ranges {
        if !seen.contains_key(&seed_range) {
            seen.insert(seed_range, true);
            unique_seed_ranges.push(seed_range);
        }
    }

    unique_seed_ranges.sort_by(|a, b| a.start.cmp(&b.start));

    let mut merged: Vec<SeedRange> = Vec::new();
    for seed_range in unique_seed_ranges {
        if merged.is_empty() || merged.last().unwrap().start + merged.last().unwrap().range <= seed_range.start {
            merged.push(seed_range);
        } else {
            merged.last_mut().unwrap().range = max(merged.last().unwrap().range, seed_range.start + seed_range.range - merged.last().unwrap().start);
        }
    }

    merged
}

fn main() {
    let stdin = io::stdin();
    let input = stdin.lock().lines().map(|l| l.unwrap()).collect::<Vec<String>>().join("\n");

    let (seed_store1_option, converters) = parse_input(&input);
    let mut seed_store1 = seed_store1_option.expect("Failed to parse seed store1");
    let mut seed_store2 = Vec::new();
    let (p_seed_store1, p_seed_store2) = (&mut seed_store1, &mut seed_store2);


    let mut converter_map = HashMap::new();
    for converter in converters {
        converter_map.insert(converter.from.clone(), converter);
    }

    let mut state = "seed".to_string();
    let mut is_converter_available = true;
    while is_converter_available {
        if let Some(converter) = converter_map.get(&state) {
            p_seed_store2.clear();
            for seed_range in p_seed_store1.iter() {
                let seed_ranges = convert(*seed_range, converter.clone());
                p_seed_store2.extend(seed_ranges);
            }
            std::mem::swap(p_seed_store1, p_seed_store2);
            *p_seed_store1 = merge_seed_ranges(p_seed_store1.clone());
            state = converter.to.clone();
        } else {
            is_converter_available = false;
        }
    }

    let min_seed = p_seed_store1.iter().min_by_key(|s| s.start).unwrap().start;
    println!("{}", min_seed);
}
