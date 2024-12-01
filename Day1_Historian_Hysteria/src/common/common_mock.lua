local common = require("common")

local mockAdapter = {
    fileDoesNotExist,
    invalidTypeInput,
}
mockAdapter.__index = mockAdapter


function mockAdapter:parseColumnsIntoTables(table1, table2)
    if self.fileDoesNotExist then
        return nil
    end

    if self.invalidTypeInput then 
        return common.parseColumnsIntoTables(1)
    end

    return common.parseColumnsIntoTables(table1, table2)
end

function mockAdapter:new(fileDoesNotExist, invalidTypeInput)
    local self = setmetatable({}, mockAdapter)
    self.fileDoesNotExist = fileDoesNotExist
    self.invalidTypeInput = invalidTypeInput
    return self
end

return mockAdapter