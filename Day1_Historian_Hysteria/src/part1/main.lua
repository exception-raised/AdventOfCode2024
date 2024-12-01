local functions = require("functions")

local function main()
    local x, y = functions.parseColumnsIntoTables("../../input.txt")
    local score = functions.calculateDistance(x, y)
    print("Distance Score:", score)
end

main()