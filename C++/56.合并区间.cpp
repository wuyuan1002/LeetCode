
// 56. 合并区间
//
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi].
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

#include <algorithm>
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> merge(std::vector<std::vector<int>>& intervals) {
        if (intervals.empty()) {
            return {};
        }

        // 按照左端点升序排序
        std::sort(intervals.begin(), intervals.end(), [](const std::vector<int>& a, const std::vector<int>& b) {
            return a[0] < b[0];
        });

        std::vector<std::vector<int>> result;     // 结果集
        std::vector<int> current = intervals[0];  // 当前总区间
        for (int i = 1; i < intervals.size(); ++i) {
            std::vector<int> interval = intervals[i];
            if (interval[0] <= current[1]) {
                // 有交集
                current[1] = std::max(current[1], interval[1]);
            } else {
                // 没有交集
                result.push_back(current);
                current = interval;
            }
        }
        result.push_back(current);

        return result;
    }
};
