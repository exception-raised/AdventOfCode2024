local functions = require("functions")
local common    = dofile("../common/common.lua")

local function main()
    local x, y = common.parseColumnsIntoTables("../../input.txt")
    local score = functions.calculateSimilarityScore(x, y)
    print("Similarity Score:", score)
end

main()