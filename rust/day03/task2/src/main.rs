use std::io::{self, BufRead};

fn process_gear(gear_pointer: &mut Vec<usize>, number: i32, numbers_per_gear: &mut Vec<i32>, gear_ratio: &mut Vec<i32>) {
    for &mut gear in gear_pointer {
        numbers_per_gear[gear - 1] += 1;
        gear_ratio[gear - 1] *= number;
    }
}

fn contains(slice: &[usize], val: usize) -> bool {
    slice.iter().any(|&x| x == val)
}

fn main() {
    let stdin = io::stdin();
    let mut grid: Vec<Vec<char>> = Vec::new();

    for line in stdin.lock().lines() {
        let line = line.unwrap();
        grid.push(line.chars().collect());
    }

    let width = grid[0].len();
    let height = grid.len();

    let mut gears_total = 0;
    let mut gears_pointers = vec![vec![0; width]; height];

    for y in 0..height {
        for x in 0..width {
            if grid[y][x] == '*' {
                gears_total += 1;
                for dy in -1..=1 {
                    for dx in -1..=1 {
                        let nx = x as i32 + dx;
                        let ny = y as i32 + dy;
                        if nx >= 0 && nx < width as i32 && ny >= 0 && ny < height as i32 {
                            gears_pointers[ny as usize][nx as usize] = gears_total;
                        }
                    }
                }
            }
        }
    }

    let mut numbers_per_gear = vec![0; gears_total];
    let mut gear_ratio = vec![1; gears_total];

    for y in 0..height {
        let mut number = 0;
        let mut gear_pointer = Vec::new();
        for x in 0..width {
            if grid[y][x].is_digit(10) {
                number = number * 10 + grid[y][x].to_digit(10).unwrap() as i32;
                if gears_pointers[y][x] != 0 && !contains(&gear_pointer, gears_pointers[y][x]) {
                    gear_pointer.push(gears_pointers[y][x]);
                }
            } else {
                process_gear(&mut gear_pointer, number, &mut numbers_per_gear, &mut gear_ratio);
                gear_pointer.clear();
                number = 0;
            }
        }
        process_gear(&mut gear_pointer, number, &mut numbers_per_gear, &mut gear_ratio);
    }

    let gear_ratio_sum: i32 = gear_ratio.iter().enumerate()
        .filter(|&(i, _)| numbers_per_gear[i] == 2)
        .map(|(_, &value)| value)
        .sum();

    println!("{}", gear_ratio_sum);
}
