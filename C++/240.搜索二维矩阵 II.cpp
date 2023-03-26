
// 240. 搜索二维矩阵 II
//
// 编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。

#include <vector>

class Solution {
public:
    // offer 04
    // 从左下角开始找
    bool searchMatrix(std::vector<std::vector<int>>& matrix, int target) {
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
