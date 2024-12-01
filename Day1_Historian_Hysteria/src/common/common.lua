local function parseColumnsIntoTables(filename) 
    if type(filename) ~= "string" then
        error("Filename must be a string.")
        return nil
    end

    local table1 = {}
    local table2 = {}

    local file = io.open(filename, "r")
    if not file then
        error("Could not open file: " .. filename)
        return nil
    end

    for line in file:lines() do
        local num1, num2 = line:match("(%d+)%s+(%d+)")
        if num1 and num2 then
            table.insert(table1, tonumber(num1))
            table.insert(table2, tonumber(num2))
        end
    end

    file:close()

    return table1, table2
end


return {
    parseColumnsIntoTables = parseColumnsIntoTables
}