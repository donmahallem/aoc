#ifndef AOC24_DAY01_H
#define AOC24_DAY01_H

#include <algorithm>
#include <istream>
#include <string>
#include <vector>

namespace aoc24::day01
{
    void parseInput(std::istream& in, std::vector<int>& left, std::vector<int>& right);
    int Part1(std::istream& in);
    int Part2(std::istream& in);
}  // namespace aoc24::day01

#endif  // AOC24_DAY01_H
