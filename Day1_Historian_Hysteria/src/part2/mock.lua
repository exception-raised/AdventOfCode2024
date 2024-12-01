local functions = require("functions")

local mockAdapter = {
    shouldFailDistanceCalc = false
}
mockAdapter.__index = mockAdapter


function mockAdapter:calculateSimilarityScore(table1, table2)
    if self.shouldFailDistanceCalc then
        error("Mocked calculateSimilarityScore failure.")
        return {}, {}
    end

    if self.invalidTypeInput then 
        return functions.calculateSimilarityScore(1, '')
    end

    return functions.calculateSimilarityScore(table1, table2)
end

function mockAdapter:new(shouldFailDistanceCalc, invalidTypeInput)
    local self = setmetatable({}, mockAdapter)
    self.shouldFailDistanceCalc = shouldFailDistanceCalc
    self.invalidTypeInput = invalidTypeInput

    return self
end

return mockAdapter