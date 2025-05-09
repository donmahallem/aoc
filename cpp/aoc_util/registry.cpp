#include "registry.h"
#include <functional>
#include <map>
#include <utility>
#include <istream>
#include <iostream>
#include "aoc_input_error.h"
namespace AocUtil
{

    /**
     *
     */
    std::function<void(int, Registry::PartFunc, Registry::PartFunc)>
    Registry::CreateYearRegistry(int year)
    {
        /**
         *
         */
        return [this, year](int day, PartFunc part1, PartFunc part2)
        {
            partRegistry[{year, day}] = {part1, part2};
        };
    }

    Registry::PartFunc Registry::Run(int year, int day, int part)
    {
        Key key = {year, day};
        if (partRegistry.count(key))
        {
            switch (part)
            {
            case 1:
                return partRegistry[key].first;
            case 2:
                return partRegistry[key].second;
            default:
                throw AocInputError::Year(part);
            }
        }
        else
        {
            throw AocInputError::YearDay(year, day);
        }
    }

}