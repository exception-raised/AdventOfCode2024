local function generateFrequencyMap(inputTable) 
    local frequencyMap = {}

    for _, value in ipairs(inputTable) do
        frequencyMap[value] = (frequencyMap[value] or 0) + 1
    end

    return frequencyMap
end

local function calculateSimilarityScore(table1, table2)
    local frequencyMap = generateFrequencyMap(table2)
    
    local similarityScore = 0
    for _, value in ipairs(table1) do
        local countInTable2 = frequencyMap[value] or 0 
        similarityScore = similarityScore + (value * countInTable2)
    end

    return similarityScore
end


return {
    parseColumnsIntoTables = parseColumnsIntoTables,
    calculateSimilarityScore = calculateSimilarityScore,
}