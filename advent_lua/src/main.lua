local day11 = require('src.day11.day11')

local function main()
    local part1 = day11.Day11('./src/day11/inputs/input.txt', 2)
    local part2 = day11.Day11('./src/day11/inputs/input.txt', 1000000)
    print('Day 11, Part 1:', part1)
    print('Day 11, Part 2:', part2)
end

main()
