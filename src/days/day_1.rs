use crate::aoc_utils::read_file_string;

pub fn day_1_part_1(filename: &str) -> i32 {
    let contents = match read_file_string(filename.to_string()) {
        Ok(contents) => contents,
        Err(e) => panic!("Error reading file: {}", e),
    };

    let mut front: char = ' ';
    let mut back: char = ' ';
    let mut sum: i32 = 0;
    for c in contents.chars() {
        match c {
            c if c.is_digit(10) => {
                if front.is_whitespace() {
                    front = c;
                }
                back = c;
            },
            c if c == '\n' => {
                let num_str: String = front.to_string() + back.to_string().as_str();  
                let num: i32 = match num_str.parse() {
                    Ok(num) => num,
                    Err(e) => panic!("Error parsing number: {}", e),
                };
                sum += num;
                front = ' ';
                back = ' ';
            },
            _ => continue,
        }
    }
    return sum;
}

pub fn day_1_part_2(filename: &str) -> i32 {
    let contents = match read_file_string(filename.to_string()) {
        Ok(contents) => contents,
        Err(e) => panic!("Error reading file: {}", e),
    };

    let chars: Vec<char> = contents.chars().collect();

    let mut builder: Vec<char> = vec![];
    let mut sum: i32 = 0;
    
    for i in 0..chars.len() { 
        let c = chars[i];
        match c {
            c if c.is_digit(10) => {
                builder.push(c);
            },
            'o' => {
                if chars[i+1] == 'n' && chars[i+2] == 'e' {
                    builder.push('1');
                } 
            },
            't' => {
                if chars[i+1] == 'w' && chars[i+2] == 'o' {
                    builder.push('2');
                } else if chars[i+1] == 'h' && chars[i+2] == 'r' && chars[i+3] == 'e' && chars[i+4] == 'e' {
                    builder.push('3');
                }
            },
            'f' => {
                if chars[i+1] == 'o' && chars[i+2] == 'u' && chars[i+3] == 'r' {
                    builder.push('4');
                } else if chars[i+1] == 'i' && chars[i+2] == 'v' && chars[i+3] == 'e' {
                    builder.push('5');
                }
            },
            's' => {
                if chars[i+1] == 'i' && chars[i+2] == 'x' {
                    builder.push('6');
                } else if chars[i+1] == 'e' && chars[i+2] == 'v' && chars[i+3] == 'e' && chars[i+4] == 'n' {
                    builder.push('7');
                }
            },
            'e' => {
                if chars[i+1] == 'i' && chars[i+2] == 'g' && chars[i+3] == 'h' && chars[i+4] == 't' {
                    builder.push('8');
                }
            },
            'n' => {
                if chars[i+1] == 'i' && chars[i+2] == 'n' && chars[i+3] == 'e' {
                    builder.push('9');
                }
            },
            '\n' => {
                let num_str: String = builder[0].to_string() 
                    + builder[builder.len()-1].to_string().as_str();
                let num: i32 = match num_str.parse() {
                    Ok(num) => num,
                    Err(e) => panic!("Error parsing number: {}", e),
                };
                sum += num;
                builder = vec![];
            },
            _ => continue,
        }
    }

    sum
}

