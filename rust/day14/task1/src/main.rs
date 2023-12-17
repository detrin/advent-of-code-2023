use std::io::{self, BufRead};

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

    roll(&mut g);
    println!("{}", score(&g));
}
