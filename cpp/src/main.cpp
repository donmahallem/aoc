#include <iostream>
#include <string>
#include <format>
#include <chrono>

#include "./aoc24/aoc24.h"

int main(int argc, char **argv)
{
    if (argc != 4)
    {
        std::cerr << "Expected args <shortYear> <day> <part>\n";
        return 1;
    }
    int year, day, part;
    try
    {
        year = std ::atoi(argv[1]);
        day = std::atoi(argv[2]);
        part = std::atoi(argv[3]);
    }
    catch (const std::exception &e)
    {
        std::cerr << "Invalid integer provided.\n";
        return 1;
    }
    AocUtil::Registry registry;
    Aoc24::RegisterParts(&registry);

    AocUtil::Registry::PartFunc partFunc = registry.Run(year, day, part);
    try
    {
        auto timeStart = std::chrono::high_resolution_clock::now();
        auto result = partFunc(std::cin);
        auto timeEnd = std::chrono::high_resolution_clock::now();
        double timeDelta = std::chrono::duration<double, std::milli>(timeEnd - timeStart).count();
        std::cout << std::format("Result: {}\nTook: {}ms\n", result, timeDelta);
    }
    catch (const std::exception &e)
    {
        std::cerr << "Error occured: " << e.what();
        return 1;
    }
    return 0;
}