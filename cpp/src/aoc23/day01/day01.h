#ifndef AOC23_DAY01_H
#define AOC23_DAY01_H

#include <istream>
#include <string>
#include <vector>

namespace aoc23::day01
{
    struct LineData
    {
        int firstDigit;
        int lastDigit;
        int firstWordDigit;
        int lastWordDigit;
    };
    int parseInput(std::istream& in, std::vector<LineData>& lines);
    long long Part1(std::istream& in);
    long long Part2(std::istream& in);
}  // namespace aoc23::day01

#endif  // AOC23_DAY01_H