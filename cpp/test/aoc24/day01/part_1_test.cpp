
#include <gtest/gtest.h>
#include "./../../../src/aoc24/day01/day01.h"
namespace AocTest
{
    namespace Aoc24Day01Test
    {
        namespace
        {
            TEST(Day01Part1Test, Part1)
            {
                std::istringstream dummyStream("3   4\n"
                                               "4   3\n"
                                               "2   5\n"
                                               "1   3\n"
                                               "3   9\n"
                                               "3   3");
                int result = Aoc24Day01::Part1(dummyStream);
                EXPECT_EQ(result, 11);
            }

        }
    }
}