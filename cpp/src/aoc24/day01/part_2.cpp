#include "day01.h"
#include "parse_input.h"

namespace Aoc24Day01
{
    int Part2(std::istream &in)
    {
        std::vector<int> left;
        std::vector<int> right;
        std::string line;
        Aoc24Day01::parseInput(in, left, right);
        std::sort(left.begin(), left.end());
        std::sort(right.begin(), right.end());

        int sum = 0;
        for (std::size_t i = 0; i < left.size(); ++i)
        {
            sum += std::abs(left[i] - right[i]);
        }

        return sum;
    }
}
