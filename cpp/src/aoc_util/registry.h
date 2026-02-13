#ifndef AOC_UTIL_REGISTRY_H
#define AOC_UTIL_REGISTRY_H
#include <any>
#include <functional>
#include <iostream>
#include <istream>
#include <map>
#include <type_traits>
#include <utility>

#include "aoc_input_error.h"

namespace aoc_util
{

    /**
     * Registry class that allows registering part functions for each year and day
     */
    class Registry
    {
       public:
        struct PartFunc
        {
            std::function<std::any(std::istream&)> run;
        };
        using DayPair = std::pair<PartFunc, PartFunc>;
        using Key = std::pair<int, int>;
        using RegistryMap = std::map<Key, DayPair>;

        // Wrap a callable returning any type into a typed part function.
        template <typename F>
        static PartFunc Wrap(F&& f)
        {
            return {
                [func = std::forward<F>(f)](std::istream& in) -> std::any
                { return std::any(func(in)); },
            };
        }

        // Returns a function which registers the day and it's part function to the year
        std::function<void(int, PartFunc, PartFunc)> CreateYearRegistry(int year);

        // Runs the part function for the given year, day and part
        PartFunc Run(int year, int day, int part);

        template <typename T>
        T RunTyped(int year, int day, int part, std::istream& in)
        {
            PartFunc partFunc = Run(year, day, part);
            return std::any_cast<T>(partFunc.run(in));
        }

       private:
        RegistryMap partRegistry;
    };

}  // namespace aoc_util
#endif  // AOC_UTIL_REGISTRY_H