use crate::aoc_utils::read_file_chars;

fn search_around(grid: &Vec<Vec<char>>, i: usize, j: usize) -> bool {
    // check for out of bounds and set new out of bounds value
    let i_min = if i == 0 { 0 } else { i - 1 };
    let i_max = if i == grid.len() - 1 { grid.len() - 1 } else { i + 1 };
    let j_min = if j == 0 { 0 } else { j - 1 };
    let j_max = if j == grid[i].len() - 1 { grid[i].len() - 1 } else { j + 1 };

    println!("{} i: {}, j: {}, i_min: {}, i_max: {}, j_min: {}, j_max: {}", grid[i][j], i, j, i_min, i_max, j_min, j_max);

    for x in i_min..=i_max {
        for y in j_min..=j_max {
            let c = grid[x][y];
            println!("{}\tChecking c: {}", grid[i][j], c);
            if !c.is_digit(10) && c != '.' {
                return true;
            }
        }
    }
    false
}

pub fn day_3_part_1(filename: &str) -> i32 {
    let contents: Vec<char> = match read_file_chars(filename.to_string()) {
        Ok(contents) => contents,
        Err(_) => panic!("Error reading file {}", filename),
    };

    let grid: Vec<Vec<char>> = contents
        .split(|c| *c == '\n')
        .map(|line| line.to_vec())
        .filter(|line| line.len() > 0)
        .collect();

    let mut tagged_nums: Vec<i32> = Vec::new();
    let mut sum: i32 = 0;
    for i in 0..grid.len() {
        let mut num_builder: String = String::new();
        let mut tagged = false;
        for j in 0..grid[i].len() {
            let c = grid[i][j];
            if c.is_digit(10) {
                num_builder.push(c);
            }
            if (!c.is_digit(10) || j == grid[i].len() - 1) && tagged {
                println!("num_builder: {}", num_builder);
                let num: i32 = num_builder.parse().unwrap();
                println!("num: {}", num);
                sum += num;
                tagged_nums.push(num);
                num_builder = String::new();
                tagged = false;
                continue;
            } else if !c.is_digit(10) || j == grid[i].len() - 1 {
                num_builder = String::new();
                tagged = false;
                continue;
            }
            if !tagged {
                tagged = search_around(&grid, i, j);
            }
        }
    }

    println!("tagged_nums: {:?}", tagged_nums);
    sum
}

struct Gear {
    i: usize,
    j: usize,
    part_numbers: Vec<i32>,
}

struct Gears {
    gears: Vec<Gear>
}

impl Gears {
    pub fn new() -> Gears {
        Gears { gears: vec![] }
    }

    fn contains(&self, i: usize, j: usize) -> bool {
        for gear in self.gears.iter() {
            if gear.i == i && gear.j == j {
                return true;
            }
        } 
        false
    }

    pub fn add(&mut self, i: usize, j: usize) -> &Gears {
        if self.contains(i, j) {
            return self;
        }
        self.gears.push(Gear{ i, j, part_numbers: vec![] });
        self
    } 

    pub fn add_part_number(&mut self, i: usize, j: usize, part_number: i32) -> &Gears {
        for idx in 0..self.gears.len() {
            if self.gears[idx].i == i && self.gears[idx].j == j {
                self.gears[idx].part_numbers.push(part_number);
            }
        }
        self
    }

    pub fn sum_gear_ratios(&self) -> i32 {
        let mut sum: i32 = 0;
        self.gears.iter()
            .filter(|gear| gear.part_numbers.len() == 2)
            .for_each(|gear| sum += gear.part_numbers[0] * gear.part_numbers[1]);
        sum
    }
}

fn search_around_gear(grid: &Vec<Vec<char>>, i: usize, j: usize) -> (bool, usize, usize) {
    // check for out of bounds and set new out of bounds value
    let i_min = if i == 0 { 0 } else { i - 1 };
    let i_max = if i == grid.len() - 1 { grid.len() - 1 } else { i + 1 };
    let j_min = if j == 0 { 0 } else { j - 1 };
    let j_max = if j == grid[i].len() - 1 { grid[i].len() - 1 } else { j + 1 };

    println!("{} i: {}, j: {}, i_min: {}, i_max: {}, j_min: {}, j_max: {}", grid[i][j], i, j, i_min, i_max, j_min, j_max);

    for x in i_min..=i_max {
        for y in j_min..=j_max {
            let c = grid[x][y];
            println!("{}\tChecking c: {}", grid[i][j], c);
            if c == '*' {
                return (true, x, y);
            }
        }
    }
    (false, 0, 0)
}
pub fn day_3_part_2(filename: &str) -> i32 {
    let contents: Vec<char> = match read_file_chars(filename.to_string()) {
        Ok(contents) => contents,
        Err(_) => panic!("Error reading file {}", filename),
    };

    let grid: Vec<Vec<char>> = contents
        .split(|c| *c == '\n')
        .map(|line| line.to_vec())
        .filter(|line| line.len() > 0)
        .collect();

    let mut gears = Gears::new();

    // build gears array
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            let c = grid[i][j];
            if c == '*' {
                gears.add(i, j);
            }
        }
    }

    let mut tagged_nums: Vec<i32> = Vec::new();
    for i in 0..grid.len() {
        let mut num_builder: String = String::new();
        let mut tagged = false;
        let mut cur_gear_i: usize = 0;
        let mut cur_gear_j: usize = 0;
        for j in 0..grid[i].len() {
            let c = grid[i][j];
            if c.is_digit(10) {
                num_builder.push(c);
            }
            if (!c.is_digit(10) || j == grid[i].len() - 1) && tagged {
                println!("num_builder: {}", num_builder);
                let num: i32 = num_builder.parse().unwrap();
                gears.add_part_number(cur_gear_i, cur_gear_j, num);
                println!("num: {}", num);
                tagged_nums.push(num);
                num_builder = String::new();
                tagged = false;
                continue;
            } else if !c.is_digit(10) || j == grid[i].len() - 1 {
                num_builder = String::new();
                tagged = false;
                continue;
            }
            if !tagged {
                (tagged, cur_gear_i, cur_gear_j) = search_around_gear(&grid, i, j);
            }
        }
    }

    println!("tagged_nums: {:?}", tagged_nums);
    gears.sum_gear_ratios()
}
