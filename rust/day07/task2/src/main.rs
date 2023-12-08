use std::collections::HashMap;
use std::cmp::Ordering;
use std::io::{self, BufRead};
use lazy_static::lazy_static;  

lazy_static! {  
    static ref CARD_RANKS: HashMap<char, u8> = {  
        let mut m = HashMap::new();  
        m.insert('A', 14);
        m.insert('K', 13); 
        m.insert('Q', 12);
        m.insert('T', 11); 
        for i in 2..=9 {  
            m.insert(std::char::from_digit(i, 10).unwrap(), (i+1).try_into().unwrap());
        }  
        m.insert('J', 2); 
        m  
    };  
}  

#[derive(Debug, PartialEq, PartialOrd, Eq, Ord)]
enum HandRanking {
    HighCard,
    OnePair,
    TwoPair,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind,
}

struct PokerHand {
    ranking: HandRanking,
    hand: String,
    bid: i32,
}

impl PokerHand {
    fn new(hand_str: String, bid: i32) -> PokerHand {
        let mut hand = PokerHand {
            ranking: HandRanking::HighCard,
            hand: "".to_string(),
            bid,
        };
        hand.set_hand(hand_str, bid);
        hand
    }

    fn compare(&self, other: &PokerHand) -> Ordering {
        if self.ranking < other.ranking {
            return Ordering::Less;
        } else if self.ranking > other.ranking {
            return Ordering::Greater;
        } else {
            for (s, o) in self.hand.chars().zip(other.hand.chars()) {
                let s_rank = CARD_RANKS[&s];
                let o_rank = CARD_RANKS[&o];
                if s_rank < o_rank {
                    return Ordering::Less;
                } else if s_rank > o_rank {
                    return Ordering::Greater;
                }
            }
            return Ordering::Equal;
        }
    }

    fn set_hand(&mut self, hand_str: String, bid: i32) {
        let mut ranks = HashMap::new();
        let mut joker_cnt = 0;
        for card in hand_str.chars() {
            if card == 'J' {
                joker_cnt += 1;
            } else {
                *ranks.entry(card).or_insert(0) += 1;
            }
        }

        let mut counts: Vec<_> = ranks.values().cloned().collect();
        counts.sort();
        counts.reverse();
        if counts.len() > 0 {
            counts[0] += joker_cnt;
        } else {
            counts.push(joker_cnt);
        }

        let ranking = match counts.len() {
            1 => HandRanking::FiveOfAKind,
            2 => {
                if counts[0] == 4 {
                    HandRanking::FourOfAKind
                } else {
                    HandRanking::FullHouse
                }
            }
            3 => {
                if counts[0] == 3 {
                    HandRanking::ThreeOfAKind
                } else {
                    HandRanking::TwoPair
                }
            }
            4 => HandRanking::OnePair,
            _ => HandRanking::HighCard,
        };

        self.ranking = ranking;
        self.hand = hand_str;
        self.bid = bid;
    }
}


fn main() {
    let stdin = io::stdin();
    let mut hands = Vec::new();

    for line in stdin.lock().lines() {
        let line = line.unwrap();
        let mut parts = line.split_whitespace();
        let hand = parts.next().unwrap().to_string();
        let bid = parts.next().unwrap().parse().unwrap();
        hands.push(PokerHand::new(hand, bid));
    }

    hands.sort_unstable_by(|a, b| a.compare(b));

    let total_winnings: i32 = hands.iter().enumerate().map(|(rank, hand)| (rank as i32 + 1) * hand.bid).sum();

    println!("{}", total_winnings);
}
