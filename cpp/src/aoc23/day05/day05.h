#ifndef AOC23_DAY05_H
#define AOC23_DAY05_H

#include <algorithm>
#include <array>
#include <istream>
#include <limits>
#include <string>
#include <vector>

namespace aoc23::day05
{
    class InputData
    {
       public:
        std::vector<long long> seeds;
        std::vector<std::array<long long, 3>> mapData[7];
    };
    int parseInput(std::istream& in, InputData& data);
    long long Part1(std::istream& in);
    long long Part2(std::istream& in);
}  // namespace aoc23::day05

#endif  // AOC23_DAY05_H
