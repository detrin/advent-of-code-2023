use std::io::{self, BufRead};

fn match_digit<'a>(s: &'a str, digit_map: &'a [(&'a str, &'a str)]) -> Option<&'a str> {
    for (word, digit) in digit_map {
        if s.starts_with(word) {
            return Some(digit);
        }
    }
    None
}

fn find_digits<'a>(s: &'a str, digit_map: &'a [(&'a str, &'a str)]) -> (Option<&'a str>, Option<&'a str>) {
    let mut first_digit = None;
    let mut last_digit = None;
    let mut pos = 0;
    while pos < s.len() {
        let ch = &s[pos..pos+1];
        if ch.chars().next().unwrap().is_alphabetic() {
            if let Some(digit) = match_digit(&s[pos..], digit_map) {
                if first_digit.is_none() {
                    first_digit = Some(digit);
                }
                last_digit = Some(digit);
            }
        } else if ch.chars().next().unwrap().is_numeric() {
            if first_digit.is_none() {
                first_digit = Some(ch);
            }
            last_digit = Some(ch);
        }
        pos += 1;
    }
    (first_digit, last_digit)
}

fn main() -> io::Result<()> {
    let stdin = io::stdin();
    let mut sum = 0;
    let digit_map: [(&str, &str); 9] = [
        ("one", "1"),
        ("two", "2"),
        ("three", "3"),
        ("four", "4"),
        ("five", "5"),
        ("six", "6"),
        ("seven", "7"),
        ("eight", "8"),
        ("nine", "9"),
    ];

    for line in stdin.lock().lines() {
        let line = line?.to_lowercase();
        let (first_digit, last_digit) = find_digits(&line, &digit_map);
        if let (Some(first_digit), Some(last_digit)) = (first_digit, last_digit) {
            let number = format!("{}{}", first_digit, last_digit).parse::<i32>().unwrap_or(0);
            sum += number;
        }
    }
    println!("{}", sum);
    Ok(())
}
