// Registry.h
#ifndef AOC_UTIL_REGISTRY_H
#define AOC_UTIL_REGISTRY_H
#include <functional>
#include <map>
#include <utility>
#include <istream>
#include <iostream>
#include "aoc_input_error.h"

namespace AocUtil
{

    class Registry
    {
    public:
        using PartFunc = std::function<int(std::istream &)>;
        using DayPair = std::pair<PartFunc, PartFunc>;
        using Key = std::pair<int, int>;
        using RegistryMap = std::map<Key, DayPair>;

        // Returns a function which registers the day and it's part function to the year
        std::function<void(int, PartFunc, PartFunc)> CreateYearRegistry(int year);

        PartFunc Run(int year, int day, int part);

    private:
        RegistryMap partRegistry;
    };

}
#endif // AOC_UTIL_REGISTRY_H