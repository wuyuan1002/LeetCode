
// 54. 螺旋矩阵
//
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

#include <vector>

class Solution {
public:
    std::vector<int> spiralOrder(std::vector<std::vector<int>>& matrix) {
        if (matrix.size() == 0 || matrix[0].size() == 0) {
            return {};
        }

        // 总结果
        std::vector<int> result;
        result.reserve(matrix.size() * matrix[0].size());

        // 要打印的圈数
        int order_num = std::min(matrix.size() + 1, matrix[0].size() + 1) / 2;

        // 打印每一圈
        for (int i = 0; i < order_num; i++) {
            printOrder(matrix, i, result);
        }

        return result;
    }

    // 打印一圈
    void printOrder(std::vector<std::vector<int>>& matrix, int start, std::vector<int>& result) {
        int endX = matrix[0].size() - start - 1;  // 该圈最大横坐标
        int endY = matrix.size() - start - 1;     // 该圈最大纵坐标

        // 打印上面一行
        for (int i = start; i <= endX; i++) {
            result.emplace_back(matrix[start][i]);
        }

        // 打印右面一列
        for (int j = start + 1; j <= endY; j++) {
            result.emplace_back(matrix[j][endX]);
        }

        // 打印下面一行
        if (endY > start) {
            for (int i = endX - 1; i >= start; i--) {
                result.emplace_back(matrix[endY][i]);
            }
        }

        // 打印左面一列
        if (endX > start) {
            for (int j = endY - 1; j > start; j--) {
                result.emplace_back(matrix[j][start]);
            }
        }
    }
};
