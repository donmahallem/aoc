#include "aoc24.h"

namespace aoc24
{
    void RegisterParts(aoc_util::Registry* registry)
    {
        using R = aoc_util::Registry;
        auto regFunc = registry->CreateYearRegistry(24);
        regFunc(1, R::Wrap(aoc24::day01::Part1), R::Wrap(aoc24::day01::Part2));
        regFunc(2, R::Wrap(aoc24::day02::Part1), R::Wrap(aoc24::day02::Part2));
        regFunc(3, R::Wrap(aoc24::day03::Part1), R::Wrap(aoc24::day03::Part2));
    }
}  // namespace aoc24