#include "aoc24.h"

namespace Aoc24
{
    void RegisterParts(AocUtil::Registry *registry)
    {
        auto regFunc = registry->CreateYearRegistry(24);
        regFunc(1, Aoc24Day01::Part1, Aoc24Day01::Part2);
        regFunc(1, Aoc24Day02::Part1, Aoc24Day02::Part1);
    }
}