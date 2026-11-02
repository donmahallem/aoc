#include "day01.h"

namespace aoc23::day01
{

    long long Part2(std::istream& in)
    {
        std::vector<LineData> lines;
        parseInput(in, lines);
        long long total = 0;
        for (const auto& ld : lines)
        {
            total += ld.firstWordDigit * 10 + ld.lastWordDigit;
        }
        return total;
    }
}  // namespace aoc23::day01
