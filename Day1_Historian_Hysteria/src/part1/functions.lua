local function calculateDistance(table1, table2)
    if type(table1) ~= "table" or type(table2) ~= "table" then
        error("inputs must be tables.")
        return {}, {}
    end

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