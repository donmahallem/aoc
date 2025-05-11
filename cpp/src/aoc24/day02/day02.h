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
}

#endif // AOC24_DAY02_H
