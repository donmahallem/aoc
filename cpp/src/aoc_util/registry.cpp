#include "registry.h"

#include <functional>
#include <iostream>
#include <istream>
#include <map>
#include <utility>

#include "aoc_input_error.h"
namespace aoc_util
{

    /**
     *
     */
    std::function<void(int, Registry::PartFunc, Registry::PartFunc)> Registry::CreateYearRegistry(
        int year)
    {
        /**
         *
         */
        return [this, year](int day, PartFunc part1, PartFunc part2)
        { partRegistry[{year, day}] = {part1, part2}; };
    }

    Registry::PartFunc Registry::Run(int year, int day, int part)
    {
        Key key = {year, day};
        auto it = partRegistry.find(key);
        if (it == partRegistry.end())
        {
            throw AocInputError::YearDay(year, day);
        }

        switch (part)
        {
            case 1:
                return it->second.first;
            case 2:
                return it->second.second;
            default:
                throw AocInputError::Part(part);
        }
    }

}  // namespace aoc_util