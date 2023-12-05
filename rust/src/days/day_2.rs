use crate::aoc_utils::read_file_lines;
use regex::Regex;


fn get_game_id(line: &str) -> i32 {
    let re = Regex::new(r"Game (?<game_id>\d+):").unwrap();
    let caps = re.captures(&line).unwrap();
    let game_id = caps.name("game_id").unwrap().as_str().parse::<i32>().unwrap();
    game_id
}

pub fn day_2_part_1(filename: &str, red_max: i32, green_max: i32, blue_max: i32) -> i32 {
    let contents = match read_file_lines(filename.to_string()) {
        Ok(contents) => contents,
        Err(_) => panic!("Error reading file {}", filename),
    };

    // let mut game_ids: Vec<i32> = vec![];
    let mut sum: i32 = 0;
    'outer: for line in contents {
        let game_id = get_game_id(&line);

        let mut sub_lines: Vec<&str> = line.split(" ").collect();

        // remove the game id
        sub_lines.remove(0);
        sub_lines.remove(0);

        for i in 0..sub_lines.len() {
            let amount = match sub_lines[i].parse::<i32>() {
                Ok(amount) => amount,
                Err(_) => continue,
            };
            let color = sub_lines[i + 1].trim_end_matches([',', ';']);
            if color == "red" && amount > red_max {
                continue 'outer;
            } else if color == "green" && amount > green_max {
                continue 'outer;
            } else if color == "blue" && amount > blue_max {
                continue 'outer;
            }
        }
        sum += game_id;
    }

    sum
}

pub fn day_2_part_2(filename: &str) -> i32 {
    let contents = match read_file_lines(filename.to_string()) {
        Ok(contents) => contents,
        Err(_) => panic!("Error reading file {}", filename),
    };

    // let mut game_ids: Vec<i32> = vec![];
    let mut sum: i32 = 0;
    for line in contents {
        let mut sub_lines: Vec<&str> = line.split(" ").collect();
        // remove the game id
        sub_lines.remove(0);
        sub_lines.remove(0);

        let mut red_min: i32 = i32::MIN;
        let mut green_min: i32 = i32::MIN;
        let mut blue_min: i32 = i32::MIN;

        for i in 0..sub_lines.len() {
            let amount = match sub_lines[i].parse::<i32>() {
                Ok(amount) => amount,
                Err(_) => continue,
            };
            let color = sub_lines[i + 1].trim_end_matches([',', ';']);
            if color == "red" && amount >= red_min {
                red_min = amount;
            } else if color == "green" && amount >= green_min {
                green_min = amount;
            } else if color == "blue" && amount >= blue_min {
                blue_min = amount;
            }
        }
        println!("{} {} {}", red_min, green_min, blue_min);
        sum += red_min * green_min * blue_min;
    }

    sum
}
