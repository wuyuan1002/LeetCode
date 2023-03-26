
// 344. 反转字符串

#include <algorithm>
#include <vector>

class Solution {
public:
    void reverseString(std::vector<char>& s) {
        for (int l = 0, r = s.size() - 1; l < r; ++l, --r) {
            std::swap(s[l], s[r]);
        }
    }
};
