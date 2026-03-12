#include "aoc23.h"

namespace aoc23
{
    void RegisterParts(aoc_util::Registry* registry)
    {
        using R = aoc_util::Registry;
        auto regFunc = registry->CreateYearRegistry(23);
        regFunc(5, R::Wrap(aoc23::day05::Part1), R::Wrap(aoc23::day05::Part2));
    }
}  // namespace aoc23