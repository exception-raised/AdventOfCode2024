local functions = require("functions")
local mock      = require("mock")

local testCases = {
    {
        description = "should succeed if expected result matches received result.",
        expected = 11,
        error = false,
        adapter = mock:new(false, false)
    },
    {
        description = "should fail if expected result does not match received result.",
        expected = math.huge,
        error = true,
        adapter = mock:new(false, false)
    },
    {
        description = "should fail if non table types are provided.",
        expected = 11,
        error = true,
        adapter = mock:new(false, true)
    }
}

local function calculateDistanceTest() 
    local testInputs = {
        {
            4,
            3,
            5,
            3,
            9,
            3,
        },
        {
            3,
            4,
            2,
            1,
            3,
            3
        }
    }
    for _, testCase in pairs(testCases) do
        local success, result = pcall(function()
            return testCase.adapter:calculateDistance(testInputs[1], testInputs[2])
        end)

        if testCase.error and result == testCase.expected then 
            error(("Test FAILED: %s - Expected failure, but calculation succeeded."):format(testCase.description))
        elseif not testCase.error and not success then
            error(("Test FAILED: %s - Expected success, but calculation failed."):format(testCase.description))
        elseif not testCase.error and result ~= testCase.expected then
            error(("Test FAILED: %s - Got %s, expected %s"):format(testCase.description, result, testCase.expected))
        else
            print(("Test PASSED: %s"):format(testCase.description))
        end
    end

end



local function runTests()
    calculateDistanceTest()
end

runTest()