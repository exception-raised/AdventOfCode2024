#include "../shared/common.h"


int countWordOccurrences(const std::vector<std::string>& grid, const std::string& word) {
    int rows = grid.size();
    int cols = grid[0].size();
    int wordLength = word.size();
    int count = 0;

    std::vector<std::pair<int, int>> directions = {
        {0, 1},  // Horizontal right
        {0, -1}, // Horizontal left
        {1, 0},  // Vertical down
        {-1, 0}, // Vertical up
        {1, 1},  // Diagonal down-right
        {1, -1}, // Diagonal down-left
        {-1, 1}, // Diagonal up-right
        {-1, -1} // Diagonal up-left
    };

    auto isValid = [&](int x, int y) {
        return x >= 0 && x < rows && y >= 0 && y < cols;
    };

    for (int r = 0; r < rows; ++r) {
        for (int c = 0; c < cols; ++c) {
            for (auto [dr, dc] : directions) {
                bool found = true;
                for (int i = 0; i < wordLength; ++i) {
                    int nr = r + dr * i;
                    int nc = c + dc * i;
                    if (!isValid(nr, nc) || grid[nr][nc] != word[i]) {
                        found = false;
                        break;
                    }
                }
                if (found) {
                    ++count;
                }
            }
        }
    }

    return count;
}

int main() {
    std::vector<std::string> grid = parseGridFromFile("../../input.txt");

    std::string word = "XMAS";

    int result = countWordOccurrences(grid, word);

    std::cout << "Total occurrences of '" << word << "': " << result << std::endl;

    return 0;
}
