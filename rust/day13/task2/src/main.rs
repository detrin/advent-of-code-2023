use std::io::stdin;
use std::str::FromStr;
use std::io::Read;

fn diff_horizontal(grid: &Vec<String>, i: usize) -> usize {
    let dist = std::cmp::min(i, grid.len() - i);
    let mut sum = 0;
    for k in 0..grid[0].len() {
        for j in 0..dist {
            if grid[i + j].chars().nth(k).unwrap() != grid[i - j - 1].chars().nth(k).unwrap() {
                sum += 1;
            }
        }
    }
    sum
}

fn total_horizontal_diffs(grid: &Vec<String>, flips: usize) -> usize {
    (1..grid.len()).filter(|&i| diff_horizontal(grid, i) == flips).sum()
}

fn check_mirror(input: &str, flips: usize) -> usize {
    let chunks: Vec<&str> = input.split("\n\n").collect();
    let mut h = 0;
    let mut v = 0;
    for chunk in chunks {
        let grid: Vec<String> = chunk.lines().map(String::from_str).collect::<Result<_, _>>().unwrap();
        h += total_horizontal_diffs(&grid, flips);
        let transposed = transpose(&grid);
        v += total_horizontal_diffs(&transposed, flips);
    }
    h * 100 + v
}

fn transpose(grid: &Vec<String>) -> Vec<String> {
    let mut transposed = vec![String::new(); grid[0].len()];
    for i in 0..grid[0].len() {
        for row in grid {
            transposed[i].push(row.chars().nth(i).unwrap());
        }
    }
    transposed
}

fn main() {
    let mut input = String::new();
    stdin().read_to_string(&mut input).unwrap();
    let total = check_mirror(&input, 1);
    println!("{}", total);
}
