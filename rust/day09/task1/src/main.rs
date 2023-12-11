use std::io::{self, BufRead};

fn get_next(arr: &[i32]) -> i32 {
    let n = arr.len();
    let mut all_zero = true;
    let mut diffs = vec![0; n - 1];
    for i in 0..n - 1 {
        diffs[i] = arr[i + 1] - arr[i];
        if diffs[i] != 0 {
            all_zero = false;
        }
    }
    if !all_zero {
        return arr[n - 1] + get_next(&diffs);
    }
    arr[n - 1]
}

fn main() {
    let stdin = io::stdin();
    let mut histories = Vec::new();
    for line in stdin.lock().lines() {
        let nums: Vec<i32> = line
            .unwrap()
            .split_whitespace()
            .map(|x| x.parse().unwrap())
            .collect();
        if nums.is_empty() {
            break;
        }
        histories.push(nums);
    }
    let mut sum = 0;
    for history in &histories {
        let next = get_next(history);
        sum += next;
    }
    println!("{}", sum);
}
