
// 139. 单词拆分
//
// 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，
// 判定s是否可以被空格拆分为一个或多个在字典中出现的单词。
//
// 说明：
//
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。

#include <string>
#include <unordered_set>
#include <vector>

class Solution {
public:
    // 动态规划
    bool wordBreak(std::string s, std::vector<std::string>& wordDict) {
        // 将单词添加到set去重, 同时提供了O(1)的查询时间复杂度
        std::unordered_set<std::string> word_set(wordDict.begin(), wordDict.end());

        // dp数组 -- 记录前i个字母能否满足单词拆分
        std::vector<bool> dp(s.size() + 1, false);
        dp[0] = true;

        // 依次遍历数组, 查看前i个能够满足单词拆分 -- 类似剪绳子
        for (int i = 1; i <= s.size(); i++) {
            for (int j = 0; j < i; j++) {
                if (dp[j] && word_set.find(std::string(s, j, i - j)) != word_set.end()) {
                    dp[i] = true;
                    break;
                }
            }
        }

        return dp[s.size()];
    }
};