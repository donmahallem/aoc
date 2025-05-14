#include "day02.h"
namespace Aoc24Day02
{
    void parseInput(std::istream &in, lineCallback callback)
    {
        std::string line;
        std::size_t lastPos, pos;
        std::vector<int> numbers{};
        while (std::getline(in, line))
        {
            numbers.clear();
            lastPos = 0;
            while ((pos = line.find(" ", lastPos)) != std::string::npos)
            {
                if (pos > lastPos)
                {
                    numbers.push_back(std::stoi(line.substr(lastPos, pos - lastPos)));
                }
                lastPos = pos + 1;
            }
            if (lastPos < line.length())
            {
                numbers.push_back(std::stoi(line.substr(lastPos)));
            }
            callback(numbers);
        }
    }
}
