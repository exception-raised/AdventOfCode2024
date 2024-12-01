local function parseColumnsIntoTables(filename) 
    if type(filename) ~= "string" then
        error("Filename must be a string.")
        return {}, {}
    end

    local table1 = {}
    local table2 = {}

    local file = io.open(filename, "r")
    if not file then
        error("Could not open file: " .. filename)
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

local function calculateDistance(table1, table2) 
    table.sort(table1)
    table.sort(table2)

    local dist = 0
    for i = 1, #table1 do
        dist = dist + math.abs(table1[i] - table2[i])
    end

    return dist
end


return {
    calculateDistance = calculateDistance,
    parseColumnsIntoTables = parseColumnsIntoTables
}