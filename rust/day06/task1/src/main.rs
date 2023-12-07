use std::io;
use std::process;

fn get_ways_to_win(time: u64, record: u64) -> u64 {
    let a = -1.0;
    let b = time as f64;
    let c = -1.0 * (record as f64);
    let discriminant = (b*b - 4.0*a*c).sqrt();
    if discriminant < 0.0 {
        return 0;
    }

    let mut lower = (-b + discriminant) / (2.0 * a);
    let mut upper = (-b - discriminant) / (2.0 * a);
    if lower.ceil() != lower {
        lower = lower.ceil();
    } else {
        lower += 1.0;
    }
    if upper.floor() != upper {
        upper = upper.floor();
    } else {
        upper -= 1.0;
    }

    let num_ways = (upper - lower + 1.0) as u64;
    if num_ways > 0 {
        num_ways
    } else {
        0
    }
}

fn read_ints(label: &str) -> Vec<u64> {
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    input = input.trim().to_string();
    input = input.replace(label, "");
    input.split_whitespace()  
        .filter_map(|x| {  
            if x.is_empty() {  
                None  
            } else {  
                Some(x.parse::<u64>().unwrap_or_else(|_| {  
                    println!("Invalid input: {}", x);  
                    process::exit(1);  
                }))  
            }  
        })  
        .collect::<Vec<u64>>()  
}

fn main() {
    let times = read_ints("Time:");
    let records = read_ints("Distance:");
    let mut total_ways = 1;
    for i in 0..times.len() {
        let ways = get_ways_to_win(times[i], records[i]);
        total_ways *= ways;
    }
    println!("{}", total_ways);
}
