local function getFileContents(filename)
    local f = io.open(filename, 'r')
    if not f then
        print("File not found: " .. filename)
        return nil
    end
    local contents = f:read("*a")
    f:close()
    return contents
end

local function printTable(t, print_func)
    if not print_func then
        print_func = print
    end

    for _, v in pairs(t) do
        print_func(v)
    end
end

local function getCoordinates(t)
    local coordinates = {}
    for y, row in ipairs(t) do
        for x = 1, #row do
            local c = row:sub(x, x)
            if c == '#' then
                table.insert(coordinates, {x = x - 1, y = y - 1})
            end
        end
    end
    return coordinates
end

local function printCoord(coord)
    print(string.format("(%d, %d)", coord.x, coord.y))
end

local function getDistance(coord1, coord2)
    return math.abs(coord1.x - coord2.x) + math.abs(coord1.y - coord2.y)
end

local function getEmptiness(grid)
    local emptiness = {
        rows = {},
        cols = {}
    }

    for y, row in ipairs(grid) do
        local row_empty = true
        for x = 1, #row do
            local c = row:sub(x, x)
            if c == '#' then
                row_empty = false
                break
            end
        end
        if row_empty then
            table.insert(emptiness.rows, y - 1)
        end
    end

    for x = 1, #grid[1] do
        local col_empty = true
        for _, row in ipairs(grid) do
            local c = row:sub(x, x)
            if c == '#' then
                col_empty = false
                break
            end
        end
        if col_empty then
            table.insert(emptiness.cols, x - 1)
        end
    end

    return emptiness
end

local function expand(coordinates, expansion_factor, emptiness)
    local new_coordinates = {}
    for _, coord in ipairs(coordinates) do
        local x_expand = 0
        local y_expand = 0
        for _, e in ipairs(emptiness.rows) do
            if coord.y > e then
                y_expand = y_expand + 1
            end
        end
        for _, e in ipairs(emptiness.cols) do
            if coord.x > e then
                x_expand = x_expand + 1
            end
        end
        table.insert(new_coordinates, {
            x = coord.x + x_expand * (expansion_factor - 1),
            y = coord.y + y_expand * (expansion_factor - 1),
        })
    end
    return new_coordinates
end

local function solve(coordinates, expansion_factor, emptiness)
    local new_coordinates = expand(coordinates, expansion_factor, emptiness)
    local sum = 0
    for i, coord in ipairs(new_coordinates) do
        for j = i + 1, #new_coordinates do
            sum = sum + getDistance(coord, new_coordinates[j])
        end
    end
    return sum
end

local M = {}

M.Day11 = function(filename, expansion_factor)
    local contents = getFileContents(filename)
    if not contents then
        return
    end

    -- split contents into lines
    local lines = {}
    for line in contents:gmatch("[^\r\n]+") do
        table.insert(lines, line)
    end
    local emptiness = getEmptiness(lines)
    local coordinates = getCoordinates(lines)
    local sum = solve(coordinates, expansion_factor, emptiness)
    return sum
end

return M
