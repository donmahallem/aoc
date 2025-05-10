#include "parse_input.h"

namespace Aoc24Day01
{
    void parseInput(std::istream &in, std::vector<int> &left, std::vector<int> &right)
    {
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
    }
}
