use std::collections::HashMap;
use std::io::{self, BufRead};

fn rotate(g: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let r = g.len();
    let c = g[0].len();
    let mut ng: Vec<Vec<char>> = vec![vec!['.'; r]; c];
    for (r_i, row) in g.iter().enumerate() {
        for (c_i, &val) in row.iter().enumerate() {
            ng[c_i][r-1-r_i] = val;
        }
    }
    ng
}

fn roll(g: &mut Vec<Vec<char>>) {
    let r = g.len();
    let c = g[0].len();
    for _ in 0..r {
        for c_i in 0..c {
            for r_i in 0..r {
                if r_i > 0 && g[r_i][c_i] == 'O' && g[r_i-1][c_i] == '.' {
                    g[r_i][c_i] = '.';
                    g[r_i-1][c_i] = 'O';
                }
            }
        }
    }
}

fn score(g: &Vec<Vec<char>>) -> usize {
    let mut ans = 0;
    let r = g.len();
    let c = g[0].len();
    for r_i in 0..r {
        for c_i in 0..c {
            if g[r_i][c_i] == 'O' {
                ans += r - r_i;
            }
        }
    }
    ans
}

fn main() {
    let stdin = io::stdin();
    let mut g: Vec<Vec<char>> = stdin.lock().lines()
        .map(|line| line.unwrap().chars().collect())
        .collect();

    let mut by_grid = HashMap::new();
    let target = 1_000_000_000;
    let mut t = 0;
    while t < target {
        t += 1;
        for _ in 0..4 {
            roll(&mut g);
            g = rotate(&g);
        }
        let gh: String = g.iter().map(|row| row.iter().collect::<String>()).collect::<Vec<String>>().join("\n");
        if let Some(val) = by_grid.get(&gh) {
            let cycle_length = t - val;
            let amt = (target - t) / cycle_length;
            t += amt * cycle_length;
        }
        by_grid.insert(gh, t);
    }
    println!("{}", score(&g));
}
