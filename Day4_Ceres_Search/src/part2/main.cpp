#include "../shared/common.h"

bool isValidXMAS(const std::vector<std::string>& grid, int r, int c, 
                 const std::string& topLeftBottomRight, const std::string& topRightBottomLeft) {
    int rows = grid.size();
    int cols = grid[0].size();

    if (r - 1 < 0 || r + 1 >= rows || c - 1 < 0 || c + 1 >= cols) {
        return false;
    }

    return grid[r - 1][c - 1] == topLeftBottomRight[0] && // Top-left
           grid[r][c] == topLeftBottomRight[1] &&        // Center
           grid[r + 1][c + 1] == topLeftBottomRight[2] && // Bottom-right
           grid[r - 1][c + 1] == topRightBottomLeft[0] && // Top-right
           grid[r + 1][c - 1] == topRightBottomLeft[2];   // Bottom-left
}

int countXMASPatterns(const std::vector<std::string>& grid) {
    int rows = grid.size();
    int cols = grid[0].size();
    int count = 0;

    for (int r = 0; r < rows; ++r) {
        for (int c = 0; c < cols; ++c) {
            if (isValidXMAS(grid, r, c, "MAS", "MAS") || // Both forward
                isValidXMAS(grid, r, c, "MAS", "SAM") || // Top-left forward, top-right backward
                isValidXMAS(grid, r, c, "SAM", "MAS") || // Top-left backward, top-right forward
                isValidXMAS(grid, r, c, "SAM", "SAM")) { // Both backward
                ++count;
            }
        }
    }

    return count;
}

int main() {
    std::string filename = "../../input.txt";

    std::vector<std::string> grid = parseGridFromFile(filename);

    if (grid.empty()) {
        std::cerr << "Error: Grid is empty or could not be read." << std::endl;
        return 1;
    }

    int result = countXMASPatterns(grid);

    std::cout << "Total X-MAS patterns: " << result << std::endl;

    return 0;
}
