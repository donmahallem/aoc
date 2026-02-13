#include <map>

#include "day01.h"
namespace aoc24::day01
{
    int Part2(std::istream& in)
    {
        std::vector<int> left;
        std::vector<int> right;
        std::string line;
        parseInput(in, left, right);

        std::map<int, int> rightListMap;

        for (auto& element : right)
        {
            std::map<int, int>::iterator iter = rightListMap.find(element);
            if (iter != rightListMap.end())
            {
                rightListMap[element] = iter->second + 1;
            }
            else
            {
                rightListMap[element] = 1;
            }
        }
        int sum = 0;

        for (auto& element : left)
        {
            std::map<int, int>::iterator iter = rightListMap.find(element);
            if (iter != rightListMap.end())
            {
                sum += element * iter->second;
            }
        }

        return sum;
    }
}  // namespace aoc24::day01
