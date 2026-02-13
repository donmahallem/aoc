#include "day02.h"

namespace aoc24::day02
{

    bool isLineFixable(const std::vector<int>& line)
    {
        bool upwardDir = line[0] < line[1];
        int lIndex = 0;
        int rIndex = 1;
        bool validLine = false;
        for (int skipIdx = -1; skipIdx < static_cast<int>(line.size()); skipIdx++)
        {
            if (skipIdx == 0)
            {
                upwardDir = line[1] < line[2];
                lIndex = 1;
                rIndex = 2;
            }
            else if (skipIdx == 1)
            {
                upwardDir = line[0] < line[2];
                lIndex = 0;
                rIndex = 2;
            }
            else
            {
                upwardDir = line[0] < line[1];
                lIndex = 0;
                rIndex = 1;
            }
            validLine = true;
            while (rIndex < line.size())
            {
                int diff = line[lIndex] - line[rIndex];
                if (diff == 0 || abs(diff) > 3 || ((diff < 0) != upwardDir))
                {
                    validLine = false;
                    break;
                }
                lIndex++;
                if (lIndex == skipIdx)
                    lIndex++;
                rIndex++;
                if (rIndex == skipIdx)
                    rIndex++;
            }

            if (validLine)
                return true;
        }
        return false;
    }
    int Part2(std::istream& in)
    {
        std::string line;
        int sum = 0;
        std::function<void(const std::vector<int>&)> cb = [&sum](const std::vector<int>& numbers)
        {
            if (isLineFixable(numbers))
            {
                sum += 1;
            }
        };
        parseInput(in, cb);

        return sum;
    }
}  // namespace aoc24::day02
