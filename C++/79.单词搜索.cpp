
// 79. 单词搜索
//
// 给定一个m x n 二维字符网格board 和一个字符串单词word 。
// 如果word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，
// 其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

#include <string>
#include <vector>

class Solution {
public:
    // 回溯法
    bool exist(std::vector<std::vector<char>>& board, std::string word)
    {
        std::vector<char> res;

        for (int i = 0; i < board.size(); ++i) {
            for (int j = 0; j < board[0].size(); ++j) {
                if (dfs(board, word, i, j, 0, res)) {
                    return true;
                }
            }
        }

        return false;
    }

    bool dfs(std::vector<std::vector<char>>& board, std::string& word, int i, int j, int index, std::vector<char>& res)
    {
        if (res.size() == word.size()) {
            return true;
        }

        bool exist = false;
        if (i >= 0 && i < board.size() && j >= 0 && j < board[0].size() && board[i][j] == word[index] && board[i][j] != '*') {

            // 标记当前位置已被访问
            char c = board[i][j];
            board[i][j] = '*';

            // 将当前字符加入到路径中
            res.push_back(c);

            // 前后左右
            exist = dfs(board, word, i + 1, j, index + 1, res) || dfs(board, word, i - 1, j, index + 1, res) || dfs(board, word, i, j + 1, index + 1, res) || dfs(board, word, i, j - 1, index + 1, res);

            // 将当前字符移出路径，并标记当前位置为未访问
            res.pop_back();
            board[i][j] = c;
        }

        return exist;
    }
};
