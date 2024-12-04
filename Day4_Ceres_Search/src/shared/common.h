#include <iostream>
#include <fstream>
#include <vector>
#include <string>

std::vector<std::string> parseGridFromFile(const std::string& filename) {
    std::vector<std::string> grid;
    std::ifstream inputFile(filename);

    if (!inputFile.is_open()) {
        std::cerr << "Error: Could not open file " << filename << std::endl;
        return grid;
    }

    std::string line;
    while (getline(inputFile, line)) {
        grid.push_back(line);
    }

    inputFile.close();
    return grid;
}