
// 93. 复原 IP 地址

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，
// 但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，
// 这些地址可以通过在 s 中插入 '.' 来形成。
// 你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

#include <algorithm>
#include <string>
#include <vector>

class Solution {
public:
    std::vector<std::string> restoreIpAddresses(std::string s)
    {
        if (s == "") {
            return {};
        }

        std::vector<std::string> res;
        std::vector<std::string> result;
        dfs(s, 0, res, result);
        return result;
    }

    void dfs(std::string& s, int start, std::vector<std::string>& res, std::vector<std::string>& result)
    {
        if (res.size() == 4) {
            if (start == s.size()) {
                std::string ip;
                for (int i = 0; i < res.size(); ++i) {
                    ip.append(res[i]);
                    if (i < res.size() - 1) {
                        ip.push_back('.');
                    }
                }
                result.push_back(ip);
            }
            return;
        }

        for (int i = start; i < s.size(); ++i) {
            std::string num = s.substr(start, i - start + 1);
            // 剪枝
            if (num != "0" && num[0] == '0') {
                return;
            } else if (std::stoi(num) > 255) {
                return;
            }

            res.push_back(num);
            dfs(s, i + 1, res, result);
            res.pop_back();
        }
    }
};
