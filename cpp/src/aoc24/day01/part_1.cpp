#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include <algorithm>
#include <cmath>

namespace Aoc24Day01
{
    int Part1(std::istream &in)
    {
        std::vector<int> left;
        std::vector<int> right;
        std::string line;

        while (std::getline(in, line))
        {
            std::size_t pos = line.find("   ");
            if (pos == std::string::npos)
            {
                continue;
            }
            std::string leftToken = line.substr(0, pos);
            std::string rightToken = line.substr(pos + 3);

            int numLeft = std::stoi(leftToken);
            int numRight = std::stoi(rightToken);

            left.push_back(numLeft);
            right.push_back(numRight);
        }
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
