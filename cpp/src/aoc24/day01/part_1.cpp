#include "day01.h"

namespace aoc24::day01
{
    int Part1(std::istream& in)
    {
        std::vector<int> left;
        std::vector<int> right;
        std::string line;
        parseInput(in, left, right);
        std::sort(left.begin(), left.end());
        std::sort(right.begin(), right.end());

        int sum = 0;
        for (std::size_t i = 0; i < left.size(); ++i)
        {
            sum += std::abs(left[i] - right[i]);
        }

        return sum;
    }
}  // namespace aoc24::day01
