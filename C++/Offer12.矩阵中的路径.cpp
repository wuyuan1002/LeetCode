
// 剑指 Offer 12. 矩阵中的路径
//
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
// 如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，
// 其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

#include <string>
#include <vector>

class Solution {
public:
    bool exist(std::vector<std::vector<char>>& board, std::string word)
    {
        std::vector<std::vector<bool>> visited(board.size(), std::vector<bool>(board[0].size(), false));
        for (int i = 0; i < board.size(); ++i) {
            for (int j = 0; j < board[0].size(); ++j) {
                if (dfs(board, i, j, 0, word, visited)) {
                    return true;
                }
            }
        }
        return false;
    }

    bool dfs(std::vector<std::vector<char>>& board, int i, int j, int index, std::string& word, std::vector<std::vector<bool>>& visited)
    {
        if (index >= word.size()) {
            return true;
        }

        bool exist = false; // 从当前位置出发, 是否存在完整路径
        if (i >= 0 && i < board.size() && j >= 0 && j < board[0].size() && word[index] == board[i][j] && !visited[i][j]) {
            // 标记当前字母已被使用
            visited[i][j] = true;

            // 查找前后左右
            exist = dfs(board, i + 1, j, index + 1, word, visited) || dfs(board, i - 1, j, index + 1, word, visited) || dfs(board, i, j + 1, index + 1, word, visited) || dfs(board, i, j - 1, index + 1, word, visited);

            // 若不存在完整路径, 标记当前位置未被使用
            if (!exist) {
                visited[i][j] = false;
            }
        }

        return exist;
    }
};
