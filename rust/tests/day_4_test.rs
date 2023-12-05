use advent_of_code_2023::days::day_4::{day_4_part_1, day_4_part_2};

#[test]
fn day_4_part_1_example_test() {
    assert_eq!(day_4_part_1("inputs/day_4_test.txt"), 13);
}


#[test]
fn day_4_part_1_input_test() {
    assert_eq!(
        day_4_part_1("inputs/day_4_input.txt"), 33950);
}

#[test]
fn day_4_part_2_example_test() {
    assert_eq!(day_4_part_2("inputs/day_4_test.txt"), 30);
}

#[test]
fn day_4_part_2_input_test() {
    assert_eq!(day_4_part_2("inputs/day_4_input.txt"), 14814534);
}
