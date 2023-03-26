
// 459. 重复的子字符串

// 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。
// 给定的字符串只含有小写英文字母，并且长度不超过10000。

#include <string>

class Solution {
public:
    // 如果 s 中没有循环节，那么 ss 中必然有且只有两个 s，
    // 此时从 ss[1] 处开始寻找 s ，必然只能找到第二个，所以此时返回值为 s.size()。
    // 如果s中有循环节，查找的返回值将是第一个s中的一个循环节下标，必然小于 s.size()
    bool repeatedSubstringPattern(std::string s) {
        return (s + s).find(s, 1) != s.size();
    }
};
