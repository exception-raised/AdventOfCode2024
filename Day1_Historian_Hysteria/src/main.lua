local input = {
    {1, 3},
    {2, 3},
    {3, 3},
    {3, 4},
    {3, 5},
    {4, 9}
}

function main()

end

main()

local filename = "../input.txt"

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

table.sort(table1)
table.sort(table2)

local dist = 0

for i = 1, #table1 do
    dist = dist + math.abs(table1[i] - table2[i])
end

print(dist)