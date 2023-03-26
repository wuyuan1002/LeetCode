
// 5. 最长回文子串
//
// 给你一个字符串 s，找到 s 中最长的回文子串。

#include <algorithm>
#include <iostream>
#include <string>
#include <unordered_map>

class Solution {
public:
    // 类似647、516
    // 中心扩散法
    std::string longestPalindrome(std::string s) {
        std::string longest = "";  // 最长回文串
        for (int i = 0; i < s.size(); ++i) {
            int l = i - 1, r = i + 1;
            for (; l >= 0 && s[l] == s[i]; --l) {
            }
            for (; r <= s.size() - 1 && s[r] == s[i]; ++r) {
            }
            for (; l >= 0 && r <= s.size() - 1 && s[l] == s[r]; --l, ++r) {
            }

            if (r - l - 1 > longest.size()) {
                longest = s.substr(l + 1, r - l - 1);
            }
        }

        return longest;
    }
};

int main() {
    std::string a = "zcsaasddsaadsa";
    std::cout << Solution().longestPalindrome(a) << std::endl;
}