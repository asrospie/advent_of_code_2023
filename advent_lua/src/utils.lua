local M = {}

M.read_file = function(path)
    local file = io.open(path, 'r')
    if not file then
        print('File not found: ' .. path)
        return nil
    end
    local content = file:read('*a')
    file:close()

    return content
end

M.print_table = function(t, f)
    for _, v in ipairs(t) do
        if f then
            print(f(v))
        else
            print(v)
        end
    end
end

return M
