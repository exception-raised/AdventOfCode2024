local functions = require("functions")
local common    = dofile("../common/common.lua")

local function main()
    local x, y = common.parseColumnsIntoTables("../../input.txt")
    local score = functions.calculateDistance(x, y)
    print("Distance Score:", score)
end

main()