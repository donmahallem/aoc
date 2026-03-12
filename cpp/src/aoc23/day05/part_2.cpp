#include <algorithm>
#include <iostream>
#include <vector>

#include "day05.h"
namespace aoc23::day05
{

    long long Part2(std::istream& in)
    {
        InputData data;
        parseInput(in, data);

        std::vector<std::pair<long long, long long>> ranges;
        for (size_t i = 0; i < data.seeds.size(); i += 2)
        {
            ranges.push_back({data.seeds[i], data.seeds[i] + data.seeds[i + 1] - 1});
        }

        for (const auto& map : data.mapData)
        {
            std::vector<std::pair<long long, long long>> next_ranges;

            while (!ranges.empty())
            {
                auto [r_start, r_end] = ranges.back();
                ranges.pop_back();

                bool matched = false;
                for (const auto& entry : map)
                {
                    long long targetStart = entry[0];
                    long long sourceStart = entry[1];
                    long long length = entry[2];

                    long long s_start = sourceStart;
                    long long s_end = sourceStart + length - 1;

                    long long o_start = std::max(r_start, s_start);
                    long long o_end = std::min(r_end, s_end);

                    if (o_start <= o_end)
                    {
                        next_ranges.push_back(
                            {targetStart + (o_start - s_start), targetStart + (o_end - s_start)});

                        if (r_start < o_start)
                        {
                            ranges.push_back({r_start, o_start - 1});
                        }
                        if (o_end < r_end)
                        {
                            ranges.push_back({o_end + 1, r_end});
                        }
                        matched = true;
                        break;
                    }
                }

                if (!matched)
                {
                    next_ranges.push_back({r_start, r_end});
                }
            }
            // Merge intervals
            if (!next_ranges.empty())
            {
                std::sort(next_ranges.begin(), next_ranges.end());
                std::vector<std::pair<long long, long long>> merged;
                merged.push_back(next_ranges[0]);
                for (size_t i = 1; i < next_ranges.size(); ++i)
                {
                    auto& last_merged = merged.back();
                    if (next_ranges[i].first <= last_merged.second)
                    {
                        last_merged.second = std::max(last_merged.second, next_ranges[i].second);
                    }
                    else
                    {
                        merged.push_back(next_ranges[i]);
                    }
                }
                ranges = std::move(merged);
            }
            else
            {
                ranges = std::move(next_ranges);
            }
        }

        long long lowestLocation = std::numeric_limits<long long>::max();
        for (const auto& r : ranges)
        {
            lowestLocation = std::min(lowestLocation, r.first);
        }
        return lowestLocation;
    }
}  // namespace aoc23::day05