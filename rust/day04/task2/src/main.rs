use std::collections::{HashSet, VecDeque};
use std::io::{self, BufRead, BufReader};

struct Cards {
    count: VecDeque<i32>,
}

impl Cards {
    fn assure_capacity(&mut self, capacity: usize) {
        while self.count.len() < capacity {
            self.count.push_back(1);
        }
    }

    fn add_cards(&mut self, pos: usize, count: i32) {
        self.assure_capacity(pos + count as usize + 1);
        for i in pos + 1..pos + count as usize + 1 {
            *self.count.get_mut(i).unwrap() += self.count[pos];
        }
    }
}

fn main() -> io::Result<()> {
    let stdin = BufReader::new(io::stdin());
    let mut generated_cards = Cards { count: VecDeque::with_capacity(25)}; 
    let mut card_num = 0;

    let mut winning_nums = HashSet::with_capacity(25); 

    for line in stdin.lines() {
        let line = line?;
        let parts: Vec<&str> = line.split(": ").collect();
        let num_parts: Vec<&str> = parts[1].split(" | ").collect();

        winning_nums.clear();
        winning_nums.extend(
            num_parts[0]
            .split_whitespace()
            .map(|num| num.parse::<i32>().unwrap())
        );

        generated_cards.assure_capacity(card_num + 1);
        let mut matched = 0;

        for num in num_parts[1].split_whitespace() {
            let actual_num: i32 = num.parse().unwrap();
            if winning_nums.contains(&actual_num) {
                matched += 1;
            }
        }
        if matched > 0 {
            generated_cards.add_cards(card_num, matched);
        }
        card_num += 1;
    }

    let generated_cards_total: i32 = generated_cards.count.iter().sum();
    println!("{}", generated_cards_total);

    Ok(())
}
