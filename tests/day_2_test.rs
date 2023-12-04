use advent_of_code_2023::days::day_2::{day_2_part_1, day_2_part_2};

#[test]
fn day_2_part_1_example_test() {
    let red_max: i32 = 12;
    let green_max: i32 = 13;
    let blue_max: i32 = 14;
    assert_eq!(
        day_2_part_1(
            "inputs/day_2_test.txt",
            red_max, green_max, blue_max), 8);
}

#[test]
fn day_2_part_1_input_test() {
    let red_max: i32 = 12;
    let green_max: i32 = 13;
    let blue_max: i32 = 14;
    assert_eq!(
        day_2_part_1(
            "inputs/day_2_input.txt",
            red_max, green_max, blue_max), 2278);
}

#[test]
fn day_1_part_2_example_test() {
    assert_eq!(day_2_part_2("inputs/day_2_test.txt"), 2286);
}

#[test]
fn day_1_part_2_input_test() {
    assert_eq!(day_2_part_2("inputs/day_2_input.txt"), 67953);
}
