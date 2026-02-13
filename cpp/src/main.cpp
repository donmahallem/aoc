#include <any>
#include <chrono>
#include <format>
#include <iostream>
#include <string>

#include "aoc24/aoc24.h"

/**
 * Helper method that converts the result of a part function to a string
 */
static std::string resultToString(const std::any& result)
{
    if (result.type() == typeid(std::string))
    {
        return std::any_cast<std::string>(result);
    }
    if (result.type() == typeid(const char*))
    {
        return std::string(std::any_cast<const char*>(result));
    }
    if (result.type() == typeid(char*))
    {
        return std::string(std::any_cast<char*>(result));
    }
    if (result.type() == typeid(int))
    {
        return std::to_string(std::any_cast<int>(result));
    }
    if (result.type() == typeid(long))
    {
        return std::to_string(std::any_cast<long>(result));
    }
    if (result.type() == typeid(long long))
    {
        return std::to_string(std::any_cast<long long>(result));
    }
    if (result.type() == typeid(unsigned int))
    {
        return std::to_string(std::any_cast<unsigned int>(result));
    }
    if (result.type() == typeid(unsigned long))
    {
        return std::to_string(std::any_cast<unsigned long>(result));
    }
    if (result.type() == typeid(unsigned long long))
    {
        return std::to_string(std::any_cast<unsigned long long>(result));
    }
    throw std::runtime_error("Unsupported result type for CLI output");
}

int main(int argc, char** argv)
{
    if (argc != 4)
    {
        std::cerr << "Expected args <shortYear> <day> <part>\n";
        return 1;
    }
    int year, day, part;
    try
    {
        size_t pos = 0;
        year = std::stoi(argv[1], &pos);
        if (argv[1][pos] != '\0')
        {
            throw std::invalid_argument("year");
        }
        day = std::stoi(argv[2], &pos);
        if (argv[2][pos] != '\0')
        {
            throw std::invalid_argument("day");
        }
        part = std::stoi(argv[3], &pos);
        if (argv[3][pos] != '\0')
        {
            throw std::invalid_argument("part");
        }
    }
    catch (const std::exception& e)
    {
        std::cerr << "Invalid integer provided.\n";
        return 1;
    }
    aoc_util::Registry registry;
    aoc24::RegisterParts(&registry);

    try
    {
        auto timeStart = std::chrono::high_resolution_clock::now();
        aoc_util::Registry::PartFunc partFunc = registry.Run(year, day, part);
        std::string result = resultToString(partFunc.run(std::cin));
        auto timeEnd = std::chrono::high_resolution_clock::now();
        double timeDelta = std::chrono::duration<double, std::milli>(timeEnd - timeStart).count();
        std::cout << std::format("Result: {}\nTook: {}ms\n", result, timeDelta);
    }
    catch (const std::exception& e)
    {
        std::cerr << "Error occured: " << e.what();
        return 1;
    }
    return 0;
}