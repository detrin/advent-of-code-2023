use std::io::{self, BufRead};

fn is_symbol(c: char) -> bool {
    !c.is_digit(10) && c != '.'
}

fn main() {
    let stdin = io::stdin();
    let grid: Vec<Vec<char>> = stdin.lock().lines()
        .map(|line| line.unwrap().chars().collect())
        .collect();

    let width = grid[0].len();
    let height = grid.len();

    let mut within_symbol = vec![vec![false; width]; height];

    for (y, row) in grid.iter().enumerate() {
        for (x, &cell) in row.iter().enumerate() {
            if is_symbol(cell) {
                for dy in -1..=1 {
                    for dx in -1..=1 {
                        let nx = x as i32 + dx;
                        let ny = y as i32 + dy;
                        if nx >= 0 && nx < width as i32 && ny >= 0 && ny < height as i32 {
                            within_symbol[ny as usize][nx as usize] = true;
                        }
                    }
                }
            }
        }
    }

    let mut numbers_within_symbol = 0;
    for (y, row) in grid.iter().enumerate() {
        let mut number = 0;
        let mut is_within_symbol = false;
        for (x, &cell) in row.iter().enumerate() {
            if cell.is_digit(10) {
                number = number * 10 + cell.to_digit(10).unwrap();
                is_within_symbol = is_within_symbol || within_symbol[y][x];
            } else {
                if is_within_symbol {
                    numbers_within_symbol += number;
                }
                number = 0;
                is_within_symbol = false;
            }
        }
        if is_within_symbol {
            numbers_within_symbol += number;
        }
    }
    println!("{}", numbers_within_symbol);
}
