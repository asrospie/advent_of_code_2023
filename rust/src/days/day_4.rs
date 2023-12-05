use crate::aoc_utils::read_file_string;
use std::collections::HashMap;

pub fn day_4_part_1(filename: &str) -> i32 {
    let contents = match read_file_string(filename.to_string()) {
        Ok(c) => c,
        Err(e) => panic!("Error reading file: {}", e),
    };

    let contents = contents.split_whitespace().filter(|x| !x.contains(":"));
    let mut winning_numbers: Vec<i32> = Vec::new();
    let mut is_testing = false;
    let mut card_score: i32 = 0;
    let mut sum: i32 = 0;
    for el in contents {
        if el == "Card" {
            is_testing = false;
            winning_numbers.clear();
            sum += card_score;
            card_score = 0;
            println!("Card score: {}", card_score);
            continue;
        }
        if el == "|" {
            is_testing = true;
            println!("Testing: {:?}", winning_numbers);
            continue;
        }
        let num = el.parse::<i32>().unwrap_or(-1);
        if is_testing && winning_numbers.contains(&num) {
            card_score = if card_score == 0 { 1 } else { card_score * 2 };
        } else if !is_testing {
            winning_numbers.push(num);
        }
    }
    println!("Card score: {}", card_score);
    sum += card_score;

    sum
}

#[derive(Debug, Clone)]
struct Card {
    winning_numbers: Vec<i32>,
    play_numbers: Vec<i32>,
    card_num: i32,
    copies: i32,
}

impl Card {
    fn new(num: i32) -> Card {
        Card {
            winning_numbers: Vec::new(),
            play_numbers: Vec::new(),
            card_num: num,
            copies: 0,
        }
    }

    pub fn add_winning_number(&mut self, num: i32) {
        self.winning_numbers.push(num);
    }

    pub fn add_play_number(&mut self, num: i32) {
        self.play_numbers.push(num);
    }

    fn score(&self) -> i32 {
        let mut score = 0;
        for num in self.play_numbers.iter() {
            if self.winning_numbers.contains(num) {
                score += 1;
            }
        }
        score
    }
}

pub fn day_4_part_2(filename: &str) -> i32 {
    let contents = match read_file_string(filename.to_string()) {
        Ok(c) => c,
        Err(e) => panic!("Error reading file: {}", e),
    };
    let contents = contents.split_whitespace()
        .filter(|x| !x.contains(":"));
    let mut cards: Vec<Card> = Vec::new();
    let mut is_testing = false;
    let mut card = Card::new(1);
    for line in contents {
        if line == "Card" {
            is_testing = false;
            cards.push(card);
            card = Card::new(cards.len() as i32);
            continue;
        }
        if line == "|" {
            is_testing = true;
            continue;
        }
        let num = line.parse::<i32>().unwrap_or(-1);
        if is_testing {
            card.add_play_number(num);
            continue;
        }
        card.add_winning_number(num);
    }
    cards.push(card);
    cards = cards.iter()
        .filter(|x| x.winning_numbers.len() > 0)
        .map(|x| x.clone()).collect();

    let mut idx: usize = 0;
    while idx < cards.len() {
        let card = cards[idx].clone();
        let card_score = card.score();

        for s in 0..card_score {
            cards[idx + s as usize + 1].copies += 1 * (card.copies + 1);
        }
        idx += 1;
    }

    // num of copies
    let mut sum: i32 = 0;
    for card in cards.iter() {
        sum += card.copies + 1;
    }
    sum
}
