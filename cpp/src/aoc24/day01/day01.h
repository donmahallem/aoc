#ifndef AOC24_DAY01_H
#define AOC24_DAY01_H

#include <string>
#include <istream>
#include <vector>
#include <algorithm>

namespace Aoc24Day01
{
    void parseInput(std::istream &in, std::vector<int> &left, std::vector<int> &right);
    int Part1(std::istream &in);
    int Part2(std::istream &in);
}

#endif // AOC24_DAY01_H
