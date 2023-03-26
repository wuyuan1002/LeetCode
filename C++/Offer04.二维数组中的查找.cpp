
// 剑指 Offer 04. 二维数组中的查找
//
// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，
// 每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，
// 输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

#include <iostream>
#include <vector>

class Solution {
public:
    // 240
    // 从左下角开始找
    bool findNumberIn2DArray(std::vector<std::vector<int>>& matrix, int target) {
        if (matrix.size() == 0 || matrix[0].size() == 0) {
            return false;
        }

        int i = matrix.size() - 1, j = 0;  // i,j指向左下角元素，由左下开始移动
        while (i >= 0 && j <= matrix[0].size() - 1) {
            if (matrix[i][j] > target) {
                i--;
            } else if (matrix[i][j] < target) {
                j++;
            } else {
                return true;
            }
        }

        return false;
    }
};

int main() {
    std::vector<std::vector<int>> list = {{}};
    bool res = Solution().findNumberIn2DArray(list, 1);
    std::cout << res << std::endl;
    return 0;
}
