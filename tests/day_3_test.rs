use advent_of_code_2023::days::day_3::{day_3_part_1, day_3_part_2};

#[test]
fn day_3_part_1_example_test() {
    assert_eq!(
        day_3_part_1("inputs/day_3_test.txt"), 4361);
}


#[test]
fn day_3_part_1_input_test() {
    assert_eq!(
        day_3_part_1("inputs/day_3_input.txt"), 550064);
}

#[test]
fn day_3_part_2_example_test() {
    assert_eq!(day_3_part_2("inputs/day_3_test.txt"), 467835);
}

#[test]
fn day_3_part_2_input_test() {
    assert_eq!(day_3_part_2("inputs/day_3_input.txt"), 85010461);
}
