local common    = require("common")
local mock      = require("common_mock")

local testCases = {
    {
        description = "should succeed if file exists and is in correct format.",
        expected = type({}),
        error = false,
        adapter = mock:new()
    },
    {
        description = "should fail if the file does not exist.",
        expected = nil,
        error = true,
        adapter = mock:new(true, false)
    },
    {
        description = "should fail if non string types are provided.",
        expected = nil,
        error = true,
        adapter = mock:new(false, true)
    }
}

local function parseColumnsIntoTablesTest() 
    for _, testCase in pairs(testCases) do
        local success, result = pcall(function()
            return testCase.adapter:parseColumnsIntoTables('test_input.txt')
        end)

        if testCase.error and type(result) == testCase.expected then 
            error(("Test FAILED: %s - Expected failure, but calculation succeeded."):format(testCase.description))
        elseif not testCase.error and not success then
            error(("Test FAILED: %s - Expected success, but calculation failed."):format(testCase.description))
        elseif not testCase.error and type(result) ~= testCase.expected then
            error(("Test FAILED: %s - Got %s, expected %s"):format(testCase.description, type(result), testCase.expected))
        else
            print(("Test PASSED: %s"):format(testCase.description))
        end
    end

end



local function runTests()
    parseColumnsIntoTablesTest()
end

runTests()