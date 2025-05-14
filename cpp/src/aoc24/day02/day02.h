#ifndef AOC24_DAY02_H
#define AOC24_DAY02_H

#include <string>
#include <istream>
#include <vector>
#include <algorithm>
#include <functional>

namespace Aoc24Day02
{
    using lineCallback = std::function<void(std::vector<int> *)>;
    void parseInput(std::istream &in, lineCallback callback);
    int Part1(std::istream &in);
    int Part2(std::istream &in);
    bool checkLine(const std::vector<int> &line);
    bool isLineFixable(const std::vector<int> &line);
}

#endif // AOC24_DAY02_H
