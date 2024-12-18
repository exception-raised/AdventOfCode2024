local functions = require("functions")

local mockAdapter = {
    shouldFailDistanceCalc = false
}
mockAdapter.__index = mockAdapter


function mockAdapter:calculateDistance(table1, table2)
    if self.shouldFailDistanceCalc then
        error("Mocked calculateDistance failure.")
        return {}, {}
    end

    if self.invalidTypeInput then 
        return functions.calculateDistance(1, '')
    end

    return functions.calculateDistance(table1, table2)
end

function mockAdapter:new(shouldFailDistanceCalc, invalidTypeInput)
    local self = setmetatable({}, mockAdapter)
    self.shouldFailDistanceCalc = shouldFailDistanceCalc
    self.invalidTypeInput = invalidTypeInput

    return self
end

return mockAdapter