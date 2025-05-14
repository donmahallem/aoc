

#include <gtest/gtest.h>
#include <vector>
#include "./../../../src/aoc24/day02/day02.h"
namespace AocTest
{
    namespace Aoc24Day02Test
    {
        namespace
        {

            TEST(Day02Part1Test, Part1)
            {
                std::istringstream dummyStream("7 6 4 2 1\n"
                                               "1 2 7 8 9\n"
                                               "9 7 6 2 1\n"
                                               "1 3 2 4 5\n"
                                               "8 6 4 4 1\n"
                                               "1 3 6 7 9");
                int result = Aoc24Day02::Part1(dummyStream);
                EXPECT_EQ(result, 2);
            }
            TEST(checklineTest, checkPasses)
            {
                std::vector<int> testData{7, 6, 4, 2, 1};
                EXPECT_TRUE(Aoc24Day02::checkLine(testData));
                testData = {1, 3, 6, 7, 9};
                EXPECT_TRUE(Aoc24Day02::checkLine(testData));
            }
            TEST(checklineTest, checkFails)
            {
                std::vector<int> testData{1, 2, 7, 8, 9};
                EXPECT_FALSE(Aoc24Day02::checkLine(testData));
                testData = {9, 7, 6, 2, 1};
                EXPECT_FALSE(Aoc24Day02::checkLine(testData));
                testData = {1, 3, 2, 4, 5};
                EXPECT_FALSE(Aoc24Day02::checkLine(testData));
                testData = {8, 6, 4, 4, 1};
                EXPECT_FALSE(Aoc24Day02::checkLine(testData));
            }
        }
    }
}