#include <iostream>

#include "day05.h"
namespace aoc23::day05
{

    long long Part1(std::istream& in)
    {
        InputData data;
        parseInput(in, data);
        long long lowestLocation = std::numeric_limits<long long>::max();
        for (const auto& seed : data.seeds)
        {
            long long location = seed;
            for (const auto& map : data.mapData)
            {
                for (const auto& entry : map)
                {
                    if (entry[1] <= location && location < entry[1] + entry[2])
                    {
                        location = location - entry[1] + entry[0];
                        break;
                    }
                }
            }
            lowestLocation = std::min(lowestLocation, location);
        }
        return lowestLocation;
    }
}  // namespace aoc23::day05