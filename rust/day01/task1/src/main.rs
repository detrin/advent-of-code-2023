use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    let stdin = io::stdin();
    let mut sum: i32 = 0;

    for line in stdin.lock().lines() {
        let line = line?;
        let digits: Vec<char> = line.chars().filter(|c| c.is_digit(10)).collect();
        if let Some((&first, &last)) = digits.first().zip(digits.last()) {
            let number = first.to_digit(10).unwrap() * 10 + last.to_digit(10).unwrap();
            sum += number as i32;
        }
    }
    println!("{}", sum);

    Ok(())
}
