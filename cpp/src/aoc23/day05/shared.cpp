#include "day05.h"

namespace aoc23::day05
{
    int parseInput(std::istream& in, aoc23::day05::InputData& data)
    {
        std::string line;
        std::size_t lastPos, pos;
        std::size_t currentMapIdx = 0;
        while (std::getline(in, line))
        {
            if (line.empty())
            {
                continue;
            }
            size_t rowIdx = line.find(":");
            if (rowIdx != std::string::npos)
            {
                if (line.starts_with("seed-to-soil"))
                {
                    currentMapIdx = 0;
                    continue;
                }
                if (line.starts_with("soil-to-fertilizer"))
                {
                    currentMapIdx = 1;
                    continue;
                }
                if (line.starts_with("fertilizer-to-water"))
                {
                    currentMapIdx = 2;
                    continue;
                }
                if (line.starts_with("water-to-light"))
                {
                    currentMapIdx = 3;
                    continue;
                }
                if (line.starts_with("light-to-temperature"))
                {
                    currentMapIdx = 4;
                    continue;
                }
                if (line.starts_with("temperature-to-humidity"))
                {
                    currentMapIdx = 5;
                    continue;
                }
                if (line.starts_with("humidity-to-location"))
                {
                    currentMapIdx = 6;
                    continue;
                }
                if (line.starts_with("seeds"))
                {
                    currentMapIdx = 7;
                    std::size_t cPos = rowIdx + 1;
                    while (cPos < line.length())
                    {
                        while (cPos < line.length() && line[cPos] == ' ')
                            cPos++;
                        if (cPos >= line.length())
                            break;
                        std::size_t endPos = line.find(' ', cPos);
                        if (endPos == std::string::npos)
                            endPos = line.length();
                        long long number = std::stoll(line.substr(cPos, endPos - cPos));
                        data.seeds.push_back(number);
                        cPos = endPos;
                    }
                    continue;
                }
            }
            // potentially line
            if (currentMapIdx == 7)
            {
                continue;
            }

            size_t cPos = 0;
            long long parsedNumbers[3] = {0, 0, 0};
            int currentNumberIdx = 0;
            while (cPos < line.length() && currentNumberIdx < 3)
            {
                while (cPos < line.length() && line[cPos] == ' ')
                    cPos++;
                if (cPos >= line.length())
                    break;
                std::size_t endPos = line.find(' ', cPos);
                if (endPos == std::string::npos)
                    endPos = line.length();
                parsedNumbers[currentNumberIdx] = std::stoll(line.substr(cPos, endPos - cPos));
                cPos = endPos;
                currentNumberIdx++;
            }
            if (currentNumberIdx == 3)
            {
                std::array<long long, 3> arr = {parsedNumbers[0], parsedNumbers[1],
                                                parsedNumbers[2]};
                data.mapData[currentMapIdx].push_back(arr);
            }
        }
        return 0;
    }
}  // namespace aoc23::day05