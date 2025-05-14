#include "day02.h"

namespace Aoc24Day02
{

    bool checkLine(const std::vector<int> &line)
    {
        int diff;
        bool upwardDir = line[0] < line[1];
        for (size_t i = 1; i < line.size(); i++)
        {
            diff = line[i] - line[i - 1];
            if (diff == 0 || abs(diff) > 3)
            {
                return false;
            }
            if ((diff < 0) == upwardDir)
            {
                return false;
            }
        }
        return true;
    }
    int Part1(std::istream &in)
    {
        std::string line;
        int sum = 0;
        parseInput(in, [&sum](const std::vector<int> &numbers)
                   {
                    if (checkLine(numbers)){
                        sum+=1;
                    } });

        return sum;
    }
}
