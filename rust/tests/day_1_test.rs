use advent_of_code_2023::days::day_1::{day_1_part_1, day_1_part_2};

#[test]
fn day_1_part_1_example_test() {
    assert_eq!(day_1_part_1("inputs/day_1_problem_1_test.txt"), 142);
}

#[test]
fn day_1_part_1_input_test() {
    assert_eq!(day_1_part_1("inputs/day_1_input.txt"), 54597);
}

#[test]
fn day_1_part_2_example_test() {
    assert_eq!(day_1_part_2("inputs/day_1_problem_2_test.txt"), 281);
}

#[test]
fn day_1_part_2_input_test() {
    assert_eq!(day_1_part_2("inputs/day_1_input.txt"), 54504);
}
