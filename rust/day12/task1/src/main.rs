use std::collections::HashMap;
use std::io::{self, BufRead};

#[derive(Hash, Eq, PartialEq)]
struct CacheKey {
    lava: String,
    springs: Vec<i32>,
}

fn recurse(lava: &str, springs: &[i32], cache: &mut HashMap<CacheKey, i32>) -> i32 {
    let key = CacheKey {
        lava: lava.to_string(),
        springs: springs.to_vec(),
    };
    if let Some(&result) = cache.get(&key) {
        return result;
    }

    let result = if springs.is_empty() {
        if lava.contains('#') { 0 } else { 1 }
    } else {
        let (&current, rest_springs) = springs.split_first().unwrap();
        let mut result = 0;
        for i in 0..=lava.len() - rest_springs.len() - rest_springs.iter().sum::<i32>() as usize - current as usize {
            if lava[..i].contains('#') {
                break;
            }
            let next = i + current as usize;
            if next > lava.len() {
                break;
            }
            if lava[i..next].contains('.') {
                continue;
            }
            if next == lava.len() {
                result += 1;
            } else if &lava[next..next+1] != "#" {
                result += recurse(&lava[next+1..], rest_springs, cache);
            }
        }
        result
    };

    cache.insert(key, result);
    result
}

fn main() {
    let stdin = io::stdin();
    let mut total_ways = 0;
    for line in stdin.lock().lines() {
        let line = line.unwrap();
        let data: Vec<&str> = line.split(' ').collect();
        let lava = data[0];
        let springs: Vec<i32> = data[1].split(',')
            .map(|s| s.parse().unwrap())
            .collect();
        let mut cache = HashMap::new();
        let combinations = recurse(lava, &springs, &mut cache);
        total_ways += combinations;
    }
    println!("{}", total_ways);
}
