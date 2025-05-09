#include "../util/registry.h"

#include "./day01/part_1.cpp"

namespace Aoc24
{
    void RegisterParts(AocUtil::Registry *registry)
    {
        auto regFunc = registry->CreateYearRegistry(24);
        regFunc(1, Aoc24Day1::Part1, Aoc24Day1::Part1);
    }
}