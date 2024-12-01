local function mockParseColumnsIntoTables(filename)
    return {1, 2, 3}, {3, 2, 1}
end

local function mockCalculateDistance(table1, table2)
    return 5
end

local mockAdapter = {
    shouldFailDistanceCalc = false
}

function mockAdapter:new(shouldFailDistanceCalc)
    local self = {}
    self.shouldFailDistanceCalc = shouldFailDistanceCalc

    return self
end
