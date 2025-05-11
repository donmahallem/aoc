#include "day02.h"

namespace Aoc24Day02
{
    int Part1(std::istream &in)
    {
        std::vector<int> left;
        std::vector<int> right;
        std::string line;
        int sum = 0;
        parseInput(in, [&sum](std::vector<int> *numbers)
                   {
        for (auto num : *numbers) {
            sum+=num;  // 'factor' is captured by value
        } });

        return sum;
    }
}
