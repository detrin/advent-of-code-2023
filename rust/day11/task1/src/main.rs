use std::io::{self, BufRead};

#[derive(Copy, Clone)]
struct Point {
    r: i64,
    c: i64,
}

impl Point {
    fn manhattan_distance(self, other: Point) -> i64 {
        (self.r - other.r).abs() + (self.c - other.c).abs()
    }
}

fn main() {
    let stdin = io::stdin();
    let buffer: Vec<String> = stdin.lock().lines().map(|l| l.unwrap()).collect();

    let max_r = buffer.len();
    let max_c = buffer[0].len();

    let mut space_warping_r = vec![0; max_r];
    let mut space_warping_c = vec![0; max_c];
    let expanding_factor = 2;

    for r in 0..max_r {
        let empty = buffer[r].chars().all(|c| c != '#');
        if r > 0 {
            space_warping_r[r] = space_warping_r[r - 1];
        }
        if empty {
            space_warping_r[r] += expanding_factor - 1;
        }
    }

    for c in 0..max_c {
        let empty = (0..max_r).all(|r| buffer[r].chars().nth(c).unwrap() != '#');
        if c > 0 {
            space_warping_c[c] = space_warping_c[c - 1];
        }
        if empty {
            space_warping_c[c] += expanding_factor - 1;
        }
    }

    let mut galaxies = Vec::new();
    for r in 0..max_r {
        for c in 0..max_c {
            if buffer[r].chars().nth(c).unwrap() == '#' {
                galaxies.push(Point {
                    r: r as i64 + space_warping_r[r],
                    c: c as i64 + space_warping_c[c],
                });
            }
        }
    }

    let total_dist = galaxies.iter()
        .enumerate()
        .map(|(i, &p1)| galaxies[i + 1..].iter().map(move |&p2| p1.manhattan_distance(p2)).sum::<i64>())
        .sum::<i64>();

    println!("{}", total_dist);
}
