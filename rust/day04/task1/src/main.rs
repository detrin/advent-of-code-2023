use std::collections::HashSet;
use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    let stdin = io::stdin();
    let mut winning_nums: HashSet<i32> = HashSet::new();

    let cards_pts_sum: i32 = stdin.lock().lines()
        .filter_map(|line| line.ok())
        .map(|line| {
            winning_nums.clear();

            let parts: Vec<&str> = line.split(": ").collect();
            let num_parts: Vec<&str> = parts[1].split(" | ").collect();
            num_parts[0]
                .split_whitespace()
                .filter_map(|num| num.parse().ok())
                .for_each(|num| { winning_nums.insert(num); });

            num_parts[1]
                .split_whitespace()
                .filter_map(|num| num.parse().ok())
                .fold(0, |pts, actual_num| {
                    if winning_nums.contains(&actual_num) {
                        if pts == 0 {
                            1
                        } else {
                            pts * 2
                        }
                    } else {
                        pts
                    }
                })
        })
        .sum();
    
    println!("{}", cards_pts_sum);

    Ok(())
}
