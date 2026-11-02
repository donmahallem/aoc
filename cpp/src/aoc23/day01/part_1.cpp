#include "day01.h"

namespace aoc23::day01
{

    long long Part1(std::istream& in)
    {
        std::vector<LineData> lines;
        parseInput(in, lines);
        long long total = 0;
        for (const auto& ld : lines)
        {
            total += ld.firstDigit * 10 + ld.lastDigit;
        }
        return total;
    }
}  // namespace aoc23::day01
