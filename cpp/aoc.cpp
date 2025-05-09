#include <iostream>
#include <string>

#include "./aoc24/aoc24.cpp"

int main(int argc, char **argv)
{
    if (argc != 4)
    {
        std::cout << "Expected args <shortYear> <day> <part>\n";
        return 1;
    }
    const int year = std ::atoi(argv[1]);
    const int day = std::atoi(argv[2]);
    const int part = std::atoi(argv[3]);

    AocUtil::Registry registry;
    Aoc24::RegisterParts(&registry);

    AocUtil::Registry::PartFunc partFunc = registry.Run(year, day, part);
    return 0;
}