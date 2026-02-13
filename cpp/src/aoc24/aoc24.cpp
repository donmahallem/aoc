#include "aoc24.h"

namespace Aoc24
{
    void RegisterParts(AocUtil::Registry *registry)
    {
        auto regFunc = registry->CreateYearRegistry(24);
        regFunc(1, Aoc24Day01::Part1, Aoc24Day01::Part2);
        regFunc(2, Aoc24Day02::Part1, Aoc24Day02::Part2);
        regFunc(3, Aoc24Day03::Part1, Aoc24Day03::Part2);
        regFunc(4, Aoc24Day04::Part1, Aoc24Day04::Part2);
        regFunc(5, Aoc24Day05::Part1, Aoc24Day05::Part2);
    }
}