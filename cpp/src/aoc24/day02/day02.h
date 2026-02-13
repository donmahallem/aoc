#ifndef AOC24_DAY02_H
#define AOC24_DAY02_H

#include <algorithm>
#include <functional>
#include <istream>
#include <string>
#include <vector>

namespace aoc24::day02
{
    using lineCallback = std::function<void(const std::vector<int>&)>;
    void parseInput(std::istream& in, lineCallback callback);
    int Part1(std::istream& in);
    int Part2(std::istream& in);
    bool checkLine(const std::vector<int>& line);
    bool isLineFixable(const std::vector<int>& line);
}  // namespace aoc24::day02

#endif  // AOC24_DAY02_H
